<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getPaymentList, deletePayment } from '#/api/play/payment';
import type { PaymentItem } from '#/api/play/payment/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 支付方式选项 */
const payTypeOptions = [
  { label: '微信支付', value: 1 },
  { label: '支付宝支付', value: 2 },
  { label: '余额支付', value: 3 },
];

/** 支付方式映射 */
const payTypeMap: Record<number, string> = {
  1: '微信支付',
  2: '支付宝支付',
  3: '余额支付',
};

/** 支付方式颜色 */
function getPayTypeColor(val: number): string {
  const keys = [1, 2, 3];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 支付状态选项 */
const payStatusOptions = [
  { label: '待支付', value: 0 },
  { label: '支付成功', value: 1 },
  { label: '支付失败', value: 2 },
  { label: '已退款', value: 3 },
];

/** 支付状态映射 */
const payStatusMap: Record<number, string> = {
  0: '待支付',
  1: '支付成功',
  2: '支付失败',
  3: '已退款',
};

/** 支付状态颜色 */
function getPayStatusColor(val: number): string {
  const keys = [0, 1, 2, 3];
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
        placeholder: '请选择支付方式',
        class: 'w-full',
      },
      fieldName: 'payType',
      label: '支付方式',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: payStatusOptions,
        placeholder: '请选择支付状态',
        class: 'w-full',
      },
      fieldName: 'payStatus',
      label: '支付状态',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<PaymentItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'orderID', title: '订单ID' },
    { field: 'memberID', title: '会员ID' },
    { field: 'paymentNo', title: '支付流水号（平台内部）' },
    { field: 'tradeNo', title: '第三方交易号' },
    { field: 'payType', title: '支付方式', width: 120, slots: { default: 'payType_cell' } },
    { field: 'payAmount', title: '支付金额（分）' },
    { field: 'payStatus', title: '支付状态', width: 120, slots: { default: 'payStatus_cell' } },
    { field: 'refundAmount', title: '退款金额（分）' },
    { field: 'callbackContent', title: '回调报文' },
    { field: 'payAt', title: '支付成功时间', width: 180, formatter: 'formatDateTime' },
    { field: 'refundAt', title: '退款时间', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getPaymentList({
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
function handleEdit(row: PaymentItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: PaymentItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该支付记录表吗？',
    okType: 'danger',
    async onOk() {
      await deletePayment(row.id);
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
