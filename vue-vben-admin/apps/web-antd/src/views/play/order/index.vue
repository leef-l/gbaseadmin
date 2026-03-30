<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getOrderList, deleteOrder } from '#/api/play/order';
import type { OrderItem } from '#/api/play/order/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** æ”¯ä»˜æ–¹å¼选项 */
const payTypeOptions = [
  { label: 'æœªæ”¯ä»˜', value: 0 },
  { label: 'å¾®ä¿¡æ”¯ä»˜', value: 1 },
  { label: 'æ”¯ä»˜å®æ”¯ä»˜', value: 2 },
  { label: 'ä½™é¢æ”¯ä»˜', value: 3 },
];

/** æ”¯ä»˜æ–¹å¼映射 */
const payTypeMap: Record<number, string> = {
  0: 'æœªæ”¯ä»˜',
  1: 'å¾®ä¿¡æ”¯ä»˜',
  2: 'æ”¯ä»˜å®æ”¯ä»˜',
  3: 'ä½™é¢æ”¯ä»˜',
};

/** æ”¯ä»˜æ–¹å¼颜色 */
function getPayTypeColor(val: number): string {
  const keys = [0, 1, 2, 3];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** è®¢å•çŠ¶æ€选项 */
const orderStatusOptions = [
  { label: 'å¾…æ”¯ä»˜', value: 0 },
  { label: 'å·²æ”¯ä»˜', value: 1 },
  { label: 'è¿›è¡Œä¸­', value: 2 },
  { label: 'å·²å®Œæˆ', value: 3 },
  { label: 'å·²å–æ¶ˆ', value: 4 },
  { label: 'é€€æ¬¾ä¸­', value: 5 },
  { label: 'å·²é€€æ¬¾', value: 6 },
];

/** è®¢å•çŠ¶æ€映射 */
const orderStatusMap: Record<number, string> = {
  0: 'å¾…æ”¯ä»˜',
  1: 'å·²æ”¯ä»˜',
  2: 'è¿›è¡Œä¸­',
  3: 'å·²å®Œæˆ',
  4: 'å·²å–æ¶ˆ',
  5: 'é€€æ¬¾ä¸­',
  6: 'å·²é€€æ¬¾',
};

/** è®¢å•çŠ¶æ€颜色 */
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
        options: orderStatusOptions,
        placeholder: '请选择è®¢å•çŠ¶æ€',
        class: 'w-full',
      },
      fieldName: 'orderStatus',
      label: 'è®¢å•çŠ¶æ€',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<OrderItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'orderNo', title: 'è®¢å•ç¼–å·' },
    { field: 'memberID', title: 'ä¸‹å•ä¼šå‘˜ID' },
    { field: 'coachID', title: 'é™ªçŽ©å¸ˆID' },
    { field: 'shopTitle', title: 'åº—é“ºID' },
    { field: 'goodsTitle', title: 'å•†å“ID' },
    { field: 'goodsTitle', title: 'å•†å“åç§°ï¼ˆå†—ä½™ï¼‰' },
    { field: 'goodsPrice', title: 'å•†å“å•ä»·ï¼ˆåˆ†ï¼‰' },
    { field: 'quantity', title: 'æ•°é‡' },
    { field: 'totalAmount', title: 'è®¢å•æ€»é¢ï¼ˆåˆ†ï¼‰' },
    { field: 'discountAmount', title: 'ä¼šå‘˜æŠ˜æ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    { field: 'couponAmount', title: 'ä¼˜æƒ åˆ¸æŠµæ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    { field: 'payAmount', title: 'å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    { field: 'couponMemberID', title: 'ä½¿ç”¨çš„ä¼˜æƒ åˆ¸é¢†å–è®°å½•ID' },
    { field: 'payType', title: 'æ”¯ä»˜æ–¹å¼', width: 120, slots: { default: 'payType_cell' } },
    { field: 'orderStatus', title: 'è®¢å•çŠ¶æ€', width: 120, slots: { default: 'orderStatus_cell' } },
    { field: 'cancelReason', title: 'å–æ¶ˆåŽŸå›' },
    { field: 'remark', title: 'è®¢å•å¤‡æ³¨' },
    { field: 'payAt', title: 'æ”¯ä»˜æ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'startAt', title: 'æœåŠ¡å¼€å§‹æ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'finishAt', title: 'æœåŠ¡å®Œæˆæ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'cancelAt', title: 'å–æ¶ˆæ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
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
    content: '确定要删除该è®¢å•è¡¨吗？',
    okType: 'danger',
    async onOk() {
      await deleteOrder(row.id);
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
      <template #orderStatus_cell="{ row }">
        <Tag :color="getOrderStatusColor(row.orderStatus)">
          {{ orderStatusMap[row.orderStatus] || row.orderStatus }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
