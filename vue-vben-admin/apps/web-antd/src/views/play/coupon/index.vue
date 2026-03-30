<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getCouponList, deleteCoupon } from '#/api/play/coupon';
import type { CouponItem } from '#/api/play/coupon/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** ä¼˜æƒ åˆ¸ç±»åž‹选项 */
const typeOptions = [
  { label: 'æ»¡å‡åˆ¸', value: 1 },
  { label: 'æŠ˜æ‰£åˆ¸', value: 2 },
  { label: 'æ— é—¨æ§›åˆ¸', value: 3 },
];

/** ä¼˜æƒ åˆ¸ç±»åž‹映射 */
const typeMap: Record<number, string> = {
  1: 'æ»¡å‡åˆ¸',
  2: 'æŠ˜æ‰£åˆ¸',
  3: 'æ— é—¨æ§›åˆ¸',
};

/** ä¼˜æƒ åˆ¸ç±»åž‹颜色 */
function getTypeColor(val: number): string {
  const keys = [1, 2, 3];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** æ˜¯å¦æ–°äººä¸“äº«选项 */
const isNewMemberOptions = [
  { label: 'å¦', value: 0 },
  { label: 'æ˜¯', value: 1 },
];

/** æ˜¯å¦æ–°äººä¸“äº«映射 */
const isNewMemberMap: Record<number, string> = {
  0: 'å¦',
  1: 'æ˜¯',
};

/** æ˜¯å¦æ–°äººä¸“äº«颜色 */
function getIsNewMemberColor(val: number): string {
  const keys = [0, 1];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** çŠ¶æ€选项 */
const statusOptions = [
  { label: 'å…³é—­', value: 0 },
  { label: 'å¼€å¯', value: 1 },
];

/** çŠ¶æ€映射 */
const statusMap: Record<number, string> = {
  0: 'å…³é—­',
  1: 'å¼€å¯',
};

/** çŠ¶æ€颜色 */
function getStatusColor(val: number): string {
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
        options: typeOptions,
        placeholder: '请选择ä¼˜æƒ åˆ¸ç±»åž‹',
        class: 'w-full',
      },
      fieldName: 'type',
      label: 'ä¼˜æƒ åˆ¸ç±»åž‹',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: isNewMemberOptions,
        placeholder: '请选择æ˜¯å¦æ–°äººä¸“äº«',
        class: 'w-full',
      },
      fieldName: 'isNewMember',
      label: 'æ˜¯å¦æ–°äººä¸“äº«',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: statusOptions,
        placeholder: '请选择çŠ¶æ€',
        class: 'w-full',
      },
      fieldName: 'status',
      label: 'çŠ¶æ€',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<CouponItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'title', title: 'ä¼˜æƒ åˆ¸åç§°' },
    { field: 'type', title: 'ä¼˜æƒ åˆ¸ç±»åž‹', width: 120, slots: { default: 'type_cell' } },
    { field: 'isNewMember', title: 'æ˜¯å¦æ–°äººä¸“äº«', width: 120, slots: { default: 'isNewMember_cell' } },
    { field: 'faceValue', title: 'é¢å€¼ï¼ˆåˆ†ï¼‰' },
    { field: 'minAmount', title: 'æœ€ä½Žæ¶ˆè´¹é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    { field: 'totalNum', title: 'å‘æ”¾æ€»é‡ï¼ˆ0ä¸é™ï¼‰' },
    { field: 'usedNum', title: 'å·²ä½¿ç”¨æ•°é‡' },
    { field: 'claimNum', title: 'å·²é¢†å–æ•°é‡' },
    { field: 'perLimit', title: 'æ¯äººé™é¢†å¼ æ•°' },
    { field: 'sort', title: 'æŽ’åº' },
    { field: 'status', title: 'çŠ¶æ€', width: 120, slots: { default: 'status_cell' } },
    { field: 'validStartAt', title: 'æœ‰æ•ˆæœŸå¼€å§‹æ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'validEndAt', title: 'æœ‰æ•ˆæœŸç»“æŸæ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getCouponList({
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
function handleEdit(row: CouponItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: CouponItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨吗？',
    okType: 'danger',
    async onOk() {
      await deleteCoupon(row.id);
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
      <template #type_cell="{ row }">
        <Tag :color="getTypeColor(row.type)">
          {{ typeMap[row.type] || row.type }}
        </Tag>
      </template>
      <template #isNewMember_cell="{ row }">
        <Tag :color="getIsNewMemberColor(row.isNewMember)">
          {{ isNewMemberMap[row.isNewMember] || row.isNewMember }}
        </Tag>
      </template>
      <template #status_cell="{ row }">
        <Tag :color="getStatusColor(row.status)">
          {{ statusMap[row.status] || row.status }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
