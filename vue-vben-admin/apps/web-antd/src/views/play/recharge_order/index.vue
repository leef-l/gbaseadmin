<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getRechargeOrderList, deleteRechargeOrder } from '#/api/play/recharge_order';
import type { RechargeOrderItem } from '#/api/play/recharge_order/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** æ”¯ä»˜æ–¹å¼选项 */
const payTypeOptions = [
  { label: 'å¾®ä¿¡æ”¯ä»˜', value: 1 },
  { label: 'æ”¯ä»˜å®æ”¯ä»˜', value: 2 },
];

/** æ”¯ä»˜æ–¹å¼映射 */
const payTypeMap: Record<number, string> = {
  1: 'å¾®ä¿¡æ”¯ä»˜',
  2: 'æ”¯ä»˜å®æ”¯ä»˜',
};

/** æ”¯ä»˜æ–¹å¼颜色 */
function getPayTypeColor(val: number): string {
  const keys = [1, 2];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** æ”¯ä»˜çŠ¶æ€选项 */
const payStatusOptions = [
  { label: 'å¾…æ”¯ä»˜', value: 0 },
  { label: 'æ”¯ä»˜æˆåŠŸ', value: 1 },
  { label: 'æ”¯ä»˜å¤±è´¥', value: 2 },
];

/** æ”¯ä»˜çŠ¶æ€映射 */
const payStatusMap: Record<number, string> = {
  0: 'å¾…æ”¯ä»˜',
  1: 'æ”¯ä»˜æˆåŠŸ',
  2: 'æ”¯ä»˜å¤±è´¥',
};

/** æ”¯ä»˜çŠ¶æ€颜色 */
function getPayStatusColor(val: number): string {
  const keys = [0, 1, 2];
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
        options: payTypeOptions,
        placeholder: '请选择æ”¯ä»˜æ–¹å¼',
        class: 'w-full',
      },
      fieldName: 'payType',
      label: 'æ”¯ä»˜æ–¹å¼',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: payStatusOptions,
        placeholder: '请选择æ”¯ä»˜çŠ¶æ€',
        class: 'w-full',
      },
      fieldName: 'payStatus',
      label: 'æ”¯ä»˜çŠ¶æ€',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<RechargeOrderItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'orderNo', title: 'å……å€¼è®¢å•å·' },
    { field: 'memberID', title: 'ä¼šå‘˜ID' },
    { field: 'rechargePlanTitle', title: 'å……å€¼æ–¹æ¡ˆID' },
    { field: 'amount', title: 'å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    { field: 'giftAmount', title: 'èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    { field: 'payType', title: 'æ”¯ä»˜æ–¹å¼', width: 120, slots: { default: 'payType_cell' } },
    { field: 'tradeNo', title: 'ç¬¬ä¸‰æ–¹äº¤æ˜“å·' },
    { field: 'payStatus', title: 'æ”¯ä»˜çŠ¶æ€', width: 120, slots: { default: 'payStatus_cell' } },
    { field: 'payAt', title: 'æ”¯ä»˜æ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getRechargeOrderList({
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
function handleEdit(row: RechargeOrderItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: RechargeOrderItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该å……å€¼è®¢å•è¡¨吗？',
    okType: 'danger',
    async onOk() {
      await deleteRechargeOrder(row.id);
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
      <template #payType_cell="{ row }">
        <Tag :color="getPayTypeColor(row.payType)">
          {{ payTypeMap[row.payType] || row.payType }}
        </Tag>
      </template>
      <template #payStatus_cell="{ row }">
        <Tag :color="getPayStatusColor(row.payStatus)">
          {{ payStatusMap[row.payStatus] || row.payStatus }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
