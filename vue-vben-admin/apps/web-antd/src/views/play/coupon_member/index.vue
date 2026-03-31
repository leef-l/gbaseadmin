<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getCouponMemberList, deleteCouponMember } from '#/api/play/coupon_member';
import type { CouponMemberItem } from '#/api/play/coupon_member/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 使用状态选项 */
const useStatusOptions = [
  { label: '未使用', value: 0 },
  { label: '已使用', value: 1 },
  { label: '已过期', value: 2 },
];

/** 使用状态映射 */
const useStatusMap: Record<number, string> = {
  0: '未使用',
  1: '已使用',
  2: '已过期',
};

/** 使用状态颜色 */
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
        placeholder: '请选择使用状态',
        class: 'w-full',
      },
      fieldName: 'useStatus',
      label: '使用状态',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<CouponMemberItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'couponTitle', title: '优惠券模板ID' },
    { field: 'memberID', title: '会员ID' },
    { field: 'orderID', title: '订单ID', slots: { header: () => h('span', {}, ['订单ID ', h(Tooltip, { title: '0表示未使用' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'useStatus', title: '使用状态', width: 120, slots: { default: 'useStatus_cell' } },
    { field: 'claimAt', title: '领取时间', width: 180, formatter: 'formatDateTime' },
    { field: 'useAt', title: '使用时间', width: 180, formatter: 'formatDateTime' },
    { field: 'expireAt', title: '过期时间', width: 180, formatter: 'formatDateTime' },
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
    content: '确定要删除该会员优惠券表吗？',
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
