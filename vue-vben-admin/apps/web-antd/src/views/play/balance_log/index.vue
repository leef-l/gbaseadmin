<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getBalanceLogList, deleteBalanceLog } from '#/api/play/balance_log';
import type { BalanceLogItem } from '#/api/play/balance_log/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 业务类型选项 */
const bizTypeOptions = [
  { label: '充值', value: 1 },
  { label: '消费', value: 2 },
  { label: '退款', value: 3 },
  { label: '活动赠送', value: 4 },
  { label: '提现', value: 5 },
];

/** 业务类型映射 */
const bizTypeMap: Record<number, string> = {
  1: '充值',
  2: '消费',
  3: '退款',
  4: '活动赠送',
  5: '提现',
};

/** 业务类型颜色 */
function getBizTypeColor(val: number): string {
  const keys = [1, 2, 3, 4, 5];
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
        options: bizTypeOptions,
        placeholder: '请选择业务类型',
        class: 'w-full',
      },
      fieldName: 'bizType',
      label: '业务类型',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<BalanceLogItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'memberID', title: '会员ID' },
    { field: 'bizType', title: '业务类型', width: 120, slots: { default: 'bizType_cell' } },
    { field: 'bizID', title: '关联业务ID（订单ID/充值订单ID/活动ID）' },
    { field: 'changeAmount', title: '变动金额（分，正数增加负数减少）' },
    { field: 'beforeBalance', title: '变动前余额（分）' },
    { field: 'afterBalance', title: '变动后余额（分）' },
    { field: 'remark', title: '备注说明' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getBalanceLogList({
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
function handleEdit(row: BalanceLogItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: BalanceLogItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该余额流水表吗？',
    okType: 'danger',
    async onOk() {
      await deleteBalanceLog(row.id);
      message.success('删除成功');
      gridApi.reload();
    },
  });
}
</script>

<template>
  <Page auto-content-height>
    <FormModalComp @success="() => gridApi.reload()" />
    <Grid>
      <template #toolbar-actions>
        <Button type="primary" @click="handleCreate">新建</Button>
      </template>
      <template #bizType_cell="{ row }">
        <Tag :color="getBizTypeColor(row.bizType)">
          {{ bizTypeMap[row.bizType] || row.bizType }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
