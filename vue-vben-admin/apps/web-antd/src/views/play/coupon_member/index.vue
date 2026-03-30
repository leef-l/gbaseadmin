<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getCouponMemberList, deleteCouponMember } from '#/api/play/coupon_member';
import type { CouponMemberItem } from '#/api/play/coupon_member/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** ä½¿ç”¨çŠ¶æ€选项 */
const useStatusOptions = [
  { label: 'æœªä½¿ç”¨', value: 0 },
  { label: 'å·²ä½¿ç”¨', value: 1 },
  { label: 'å·²è¿‡æœŸ', value: 2 },
];

/** ä½¿ç”¨çŠ¶æ€映射 */
const useStatusMap: Record<number, string> = {
  0: 'æœªä½¿ç”¨',
  1: 'å·²ä½¿ç”¨',
  2: 'å·²è¿‡æœŸ',
};

/** ä½¿ç”¨çŠ¶æ€颜色 */
function getUseStatusColor(val: number): string {
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
        options: useStatusOptions,
        placeholder: '请选择ä½¿ç”¨çŠ¶æ€',
        class: 'w-full',
      },
      fieldName: 'useStatus',
      label: 'ä½¿ç”¨çŠ¶æ€',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<CouponMemberItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'couponTitle', title: 'ä¼˜æƒ åˆ¸æ¨¡æ¿ID' },
    { field: 'memberID', title: 'ä¼šå‘˜ID' },
    { field: 'orderID', title: 'ä½¿ç”¨çš„è®¢å•ID' },
    { field: 'useStatus', title: 'ä½¿ç”¨çŠ¶æ€', width: 120, slots: { default: 'useStatus_cell' } },
    { field: 'claimAt', title: 'é¢†å–æ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'useAt', title: 'ä½¿ç”¨æ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'expireAt', title: 'è¿‡æœŸæ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getCouponMemberList({
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
function handleEdit(row: CouponMemberItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: CouponMemberItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨吗？',
    okType: 'danger',
    async onOk() {
      await deleteCouponMember(row.id);
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
      <template #useStatus_cell="{ row }">
        <Tag :color="getUseStatusColor(row.useStatus)">
          {{ useStatusMap[row.useStatus] || row.useStatus }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
