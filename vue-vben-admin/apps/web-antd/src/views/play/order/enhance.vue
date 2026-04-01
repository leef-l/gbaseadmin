<script setup lang="ts">
import { h } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, Input, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getOrderList, deleteOrder } from '#/api/play/order';
import { changeOrderStatus } from '#/api/play/order/enhance';
import type { OrderItem } from '#/api/play/order/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 支付方式选项 */
const payTypeOptions = [
  { label: '未支付', value: 0 },
  { label: '微信支付', value: 1 },
  { label: '支付宝支付', value: 2 },
  { label: '余额支付', value: 3 },
];

/** 支付方式映射 */
const payTypeMap: Record<number, string> = {
  0: '未支付',
  1: '微信支付',
  2: '支付宝支付',
  3: '余额支付',
};

/** 支付方式颜色 */
function getPayTypeColor(val: number): string {
  const keys = [0, 1, 2, 3];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 订单状态选项 */
const orderStatusOptions = [
  { label: '待支付', value: 0 },
  { label: '已支付', value: 1 },
  { label: '进行中', value: 2 },
  { label: '已完成', value: 3 },
  { label: '已取消', value: 4 },
  { label: '退款中', value: 5 },
  { label: '已退款', value: 6 },
];

/** 订单状态映射 */
const orderStatusMap: Record<number, string> = {
  0: '待支付',
  1: '已支付',
  2: '进行中',
  3: '已完成',
  4: '已取消',
  5: '退款中',
  6: '已退款',
};

/** 订单状态颜色 */
function getOrderStatusColor(val: number): string {
  const keys = [0, 1, 2, 3, 4, 5, 6];
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
        options: orderStatusOptions,
        placeholder: '请选择订单状态',
        class: 'w-full',
      },
      fieldName: 'orderStatus',
      label: '订单状态',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<OrderItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'orderNo', title: '订单编号' },
    { field: 'memberNickname', title: '下单会员' },
    { field: 'coachRealName', title: '陪玩师' },
    { field: 'shopTitle', title: '店铺ID（0表示无店铺）' },
    { field: 'goodsID', title: '商品ID' },
    { field: 'goodsTitle', title: '商品名称（冗余）' },
    { field: 'goodsPrice', title: '商品单价（元）', formatter: ({ cellValue }: { cellValue: number | null }) => cellValue != null ? (cellValue / 100).toFixed(2) : '-' },
    { field: 'quantity', title: '数量' },
    { field: 'totalAmount', title: '订单总额（元）', formatter: ({ cellValue }: { cellValue: number | null }) => cellValue != null ? (cellValue / 100).toFixed(2) : '-' },
    { field: 'discountAmount', title: '会员折扣金额（元）', formatter: ({ cellValue }: { cellValue: number | null }) => cellValue != null ? (cellValue / 100).toFixed(2) : '-' },
    { field: 'couponAmount', title: '优惠券抵扣金额（元）', formatter: ({ cellValue }: { cellValue: number | null }) => cellValue != null ? (cellValue / 100).toFixed(2) : '-' },
    { field: 'payAmount', title: '实付金额（元）', formatter: ({ cellValue }: { cellValue: number | null }) => cellValue != null ? (cellValue / 100).toFixed(2) : '-' },
    { field: 'couponMemberID', title: '使用的优惠券领取记录ID' },
    { field: 'payType', title: '支付方式', width: 120, slots: { default: 'payType_cell' } },
    { field: 'orderStatus', title: '订单状态', width: 120, slots: { default: 'orderStatus_cell' } },
    { field: 'cancelReason', title: '取消原因' },
    { field: 'remark', title: '订单备注' },
    { field: 'payAt', title: '支付时间', width: 180, formatter: 'formatDateTime' },
    { field: 'startAt', title: '服务开始时间', width: 180, formatter: 'formatDateTime' },
    { field: 'finishAt', title: '服务完成时间', width: 180, formatter: 'formatDateTime' },
    { field: 'cancelAt', title: '取消时间', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 300, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getOrderList({
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
function handleEdit(row: OrderItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: OrderItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该订单表吗？',
    okType: 'danger',
    async onOk() {
      await deleteOrder(row.id);
      message.success('删除成功');
      gridApi.reload();
    },
  });
}

/** 状态流转映射：当前状态 -> 可执行的下一步操作 */
const statusActions: Record<number, { label: string; status: number; danger?: boolean; needReason?: boolean }[]> = {
  0: [{ label: '取消订单', status: 4, danger: true, needReason: true }],
  1: [{ label: '开始服务', status: 2 }],
  2: [{ label: '完成服务', status: 3 }],
  5: [{ label: '同意退款', status: 6, danger: true }, { label: '拒绝退款', status: 1 }],
};

/** 变更订单状态（取消订单需填写原因） */
function handleChangeStatus(row: OrderItem, targetStatus: number, label: string, needReason?: boolean) {
  if (needReason) {
    let cancelReason = '';
    Modal.confirm({
      title: `${label}`,
      content: () => h('div', [
        h('p', `确定要将订单 ${row.orderNo} 取消吗？`),
        h(Input.TextArea, {
          placeholder: '请输入取消原因',
          rows: 3,
          onChange: (e: any) => { cancelReason = e.target.value; },
        }),
      ]),
      okType: 'danger',
      async onOk() {
        if (!cancelReason.trim()) {
          message.warning('请输入取消原因');
          throw new Error('请输入取消原因');
        }
        await changeOrderStatus({ id: row.id, orderStatus: targetStatus, cancelReason });
        message.success(`${label}成功`);
        gridApi.reload();
      },
    });
  } else {
    Modal.confirm({
      title: `${label}确认`,
      content: `确定要将订单 ${row.orderNo} ${label}吗？`,
      okType: targetStatus === 6 ? 'danger' : 'primary',
      async onOk() {
        await changeOrderStatus({ id: row.id, orderStatus: targetStatus });
        message.success(`${label}成功`);
        gridApi.reload();
      },
    });
  }
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
      <template #orderStatus_cell="{ row }">
        <Tag :color="getOrderStatusColor(row.orderStatus)">
          {{ orderStatusMap[row.orderStatus] || row.orderStatus }}
        </Tag>
      </template>
      <template #action="{ row }">
        <template v-for="act in (statusActions[row.orderStatus] || [])" :key="act.status">
          <Button type="link" :danger="act.danger" size="small" @click="handleChangeStatus(row, act.status, act.label, act.needReason)">{{ act.label }}</Button>
        </template>
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>