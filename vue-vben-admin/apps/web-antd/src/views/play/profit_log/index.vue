<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getProfitLogList, deleteProfitLog } from '#/api/play/profit_log';
import type { ProfitLogItem } from '#/api/play/profit_log/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** ç»“ç®—çŠ¶æ€选项 */
const settleStatusOptions = [
  { label: 'å¾…ç»“ç®—', value: 0 },
  { label: 'å·²ç»“ç®—', value: 1 },
];

/** ç»“ç®—çŠ¶æ€映射 */
const settleStatusMap: Record<number, string> = {
  0: 'å¾…ç»“ç®—',
  1: 'å·²ç»“ç®—',
};

/** ç»“ç®—çŠ¶æ€颜色 */
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
        placeholder: '请选择ç»“ç®—çŠ¶æ€',
        class: 'w-full',
      },
      fieldName: 'settleStatus',
      label: 'ç»“ç®—çŠ¶æ€',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<ProfitLogItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'orderID', title: 'è®¢å•ID' },
    { field: 'orderNo', title: 'è®¢å•ç¼–å·' },
    { field: 'payAmount', title: 'å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    { field: 'coachID', title: 'é™ªçŽ©å¸ˆID' },
    { field: 'shopTitle', title: 'åº—é“ºID' },
    { field: 'platformRate', title: 'å¹³å°æŠ½æˆæ¯”ä¾‹' },
    { field: 'platformAmount', title: 'å¹³å°æŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰' },
    { field: 'shopRate', title: 'åº—é“ºæŠ½æˆæ¯”ä¾‹' },
    { field: 'shopAmount', title: 'åº—é“ºæŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰' },
    { field: 'coachAmount', title: 'é™ªçŽ©å¸ˆæ”¶å…¥ï¼ˆåˆ†ï¼‰' },
    { field: 'settleStatus', title: 'ç»“ç®—çŠ¶æ€', width: 120, slots: { default: 'settleStatus_cell' } },
    { field: 'settleAt', title: 'ç»“ç®—æ—¶é—´', width: 180, formatter: 'formatDateTime' },
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
    content: '确定要删除该åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨吗？',
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
    <FormModalComp @success="() => gridApi.reload()" />
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
