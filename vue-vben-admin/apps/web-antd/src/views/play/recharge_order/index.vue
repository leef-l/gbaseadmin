<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';
import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getRechargeOrderList, deleteRechargeOrder } from '#/api/play/recharge_order';
import type { RechargeOrderItem } from '#/api/play/recharge_order/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 支付方式选项 */
const payTypeOptions = [
  { label: '微信支付', value: 1 },
  { label: '支付宝支付', value: 2 },
];

/** 支付方式映射 */
const payTypeMap: Record<number, string> = {
  1: '微信支付',
  2: '支付宝支付',
};

/** 支付方式颜色 */
function getPayTypeColor(val: number): string {
  const keys = [1, 2];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 支付状态选项 */
const payStatusOptions = [
  { label: '待支付', value: 0 },
  { label: '支付成功', value: 1 },
  { label: '支付失败', value: 2 },
];

/** 支付状态映射 */
const payStatusMap: Record<number, string> = {
  0: '待支付',
  1: '支付成功',
  2: '支付失败',
};

/** 支付状态颜色 */
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
const gridOptions: VxeGridProps<RechargeOrderItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'orderNo', title: '充值订单号' },
    { field: 'memberID', title: '会员ID' },
    { field: 'rechargePlanTitle', title: '充值方案ID' },
    { field: 'amount', title: '充值金额', slots: { header: () => h('span', {}, ['充值金额 ', h(Tooltip, { title: '单位：分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'giftAmount', title: '赠送金额', slots: { header: () => h('span', {}, ['赠送金额 ', h(Tooltip, { title: '单位：分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'payType', title: '支付方式', width: 120, slots: { default: 'payType_cell' } },
    { field: 'tradeNo', title: '第三方交易号' },
    { field: 'payStatus', title: '支付状态', width: 120, slots: { default: 'payStatus_cell' } },
    { field: 'payAt', title: '支付时间', width: 180, formatter: 'formatDateTime' },
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
    content: '确定要删除该充值订单表吗？',
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
