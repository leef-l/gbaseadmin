<script setup lang="ts">
import { ref, onMounted } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, Card, Col, message, Modal, Row, Statistic, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getProfitLogList, deleteProfitLog } from '#/api/play/profit_log';
import type { ProfitLogItem } from '#/api/play/profit_log/types';
import FormModal from './modules/form.vue';

/** 统计数据 */
const stats = ref({ total: 0, platform: 0, shop: 0, coach: 0 });

async function loadStats() {
  try {
    const res = await getProfitLogList({ pageNum: 1, pageSize: 9999 });
    const list = res?.list ?? [];
    let total = 0, platform = 0, shop = 0, coach = 0;
    for (const item of list) {
      const pay = Number(item.payAmount) || 0;
      total += pay;
      platform += Number(item.platformAmount) || 0;
      shop += Number(item.shopAmount) || 0;
      coach += Number(item.coachAmount) || 0;
    }
    stats.value = { total, platform, shop, coach };
  } catch { /* ignore */ }
}

onMounted(loadStats);

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 结算状态选项 */
const settleStatusOptions = [
  { label: '待结算', value: 0 },
  { label: '已结算', value: 1 },
];

/** 结算状态映射 */
const settleStatusMap: Record<number, string> = {
  0: '待结算',
  1: '已结算',
};

/** 结算状态颜色 */
function getSettleStatusColor(val: number): string {
  const keys = [0, 1];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 表单弹窗 */
const [FormModalComp, formModalApi] = useVbenModal({
  connectedComponent: FormModal,
  destroyOnClose: true,
});

/** 搜索表单配置 */
const formOptions: VbenFormProps = {
  collapsed: false,
  showCollapseButton: true,
  submitOnChange: false,
  submitOnEnter: true,
  schema: [
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: settleStatusOptions,
        placeholder: '请选择结算状态',
        class: 'w-full',
      },
      fieldName: 'settleStatus',
      label: '结算状态',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<ProfitLogItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'orderID', title: '订单ID' },
    { field: 'orderNo', title: '订单编号' },
    { field: 'payAmount', title: '实付金额（元）', formatter: ({ cellValue }: { cellValue: number | null }) => cellValue != null ? (cellValue / 100).toFixed(2) : '-' },
    { field: 'coachRealName', title: '陪玩师' },
    { field: 'shopTitle', title: '店铺ID（0表示无店铺）' },
    { field: 'platformRate', title: '平台抽成比例（百分比）' },
    { field: 'platformAmount', title: '平台抽成金额（元）', formatter: ({ cellValue }: { cellValue: number | null }) => cellValue != null ? (cellValue / 100).toFixed(2) : '-' },
    { field: 'shopRate', title: '店铺抽成比例（百分比）' },
    { field: 'shopAmount', title: '店铺抽成金额（元）', formatter: ({ cellValue }: { cellValue: number | null }) => cellValue != null ? (cellValue / 100).toFixed(2) : '-' },
    { field: 'coachAmount', title: '陪玩师收入（元）', formatter: ({ cellValue }: { cellValue: number | null }) => cellValue != null ? (cellValue / 100).toFixed(2) : '-' },
    { field: 'settleStatus', title: '结算状态', width: 120, slots: { default: 'settleStatus_cell' } },
    { field: 'settleAt', title: '结算时间', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getProfitLogList({
          pageNum: page.currentPage,
          pageSize: page.pageSize,
          ...formValues,
        });
        return { items: res?.list ?? [], total: res?.total ?? 0 };
      },
    },
  },
  toolbarConfig: {
    custom: true,
    refresh: true,
    search: true,
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions,
  gridOptions,
});

/** 新建 */
function handleCreate() {
  formModalApi.setData(null).open();
}

/** 编辑 */
function handleEdit(row: ProfitLogItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: ProfitLogItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该利润分成流水表吗？',
    okType: 'danger',
    async onOk() {
      await deleteProfitLog(row.id);
      message.success('删除成功');
      gridApi.reload();
    },
  });
}
</script>
<template>
  <Page auto-content-height>
    <Row :gutter="16" style="margin-bottom: 16px">
      <Col :span="6">
        <Card>
          <Statistic title="总利润（元）" :value="(stats.total / 100).toFixed(2)" :value-style="{ color: '#1890ff' }" />
        </Card>
      </Col>
      <Col :span="6">
        <Card>
          <Statistic title="平台收入（元）" :value="(stats.platform / 100).toFixed(2)" :value-style="{ color: '#7c3aed' }" />
        </Card>
      </Col>
      <Col :span="6">
        <Card>
          <Statistic title="店铺收入（元）" :value="(stats.shop / 100).toFixed(2)" :value-style="{ color: '#f59e0b' }" />
        </Card>
      </Col>
      <Col :span="6">
        <Card>
          <Statistic title="陪玩师收入（元）" :value="(stats.coach / 100).toFixed(2)" :value-style="{ color: '#10b981' }" />
        </Card>
      </Col>
    </Row>
    <FormModalComp @success="() => { gridApi.reload(); loadStats(); }" />
    <Grid>
      <template #toolbar-actions>
        <Button type="primary" @click="handleCreate">新建</Button>
      </template>
      <template #settleStatus_cell="{ row }">
        <Tag :color="getSettleStatusColor(row.settleStatus)">
          {{ settleStatusMap[row.settleStatus] || row.settleStatus }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
