<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getCouponList, deleteCoupon } from '#/api/play/coupon';
import type { CouponItem } from '#/api/play/coupon/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 优惠券类型选项 */
const typeOptions = [
  { label: '满减券', value: 1 },
  { label: '折扣券', value: 2 },
  { label: '无门槛券', value: 3 },
];

/** 优惠券类型映射 */
const typeMap: Record<number, string> = {
  1: '满减券',
  2: '折扣券',
  3: '无门槛券',
};

/** 优惠券类型颜色 */
function getTypeColor(val: number): string {
  const keys = [1, 2, 3];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 是否新人专享选项 */
const isNewMemberOptions = [
  { label: '否', value: 0 },
  { label: '是', value: 1 },
];

/** 是否新人专享映射 */
const isNewMemberMap: Record<number, string> = {
  0: '否',
  1: '是',
};

/** 是否新人专享颜色 */
function getIsNewMemberColor(val: number): string {
  const keys = [0, 1];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 状态选项 */
const statusOptions = [
  { label: '关闭', value: 0 },
  { label: '开启', value: 1 },
];

/** 状态映射 */
const statusMap: Record<number, string> = {
  0: '关闭',
  1: '开启',
};

/** 状态颜色 */
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
        placeholder: '请选择优惠券类型',
        class: 'w-full',
      },
      fieldName: 'type',
      label: '优惠券类型',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: isNewMemberOptions,
        placeholder: '请选择是否新人专享',
        class: 'w-full',
      },
      fieldName: 'isNewMember',
      label: '是否新人专享',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: statusOptions,
        placeholder: '请选择状态',
        class: 'w-full',
      },
      fieldName: 'status',
      label: '状态',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<CouponItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'title', title: '优惠券名称' },
    { field: 'type', title: '优惠券类型', width: 120, slots: { default: 'type_cell' } },
    { field: 'isNewMember', title: '是否新人专享', width: 120, slots: { default: 'isNewMember_cell' } },
    { field: 'faceValue', title: '面值', slots: { header: () => h('span', {}, ['面值 ', h(Tooltip, { title: '分，满减/无门槛时为抵扣额，折扣时为折扣值如 85=8.5折' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'minAmount', title: '最低消费', slots: { header: () => h('span', {}, ['最低消费 ', h(Tooltip, { title: '分，0表示无门槛' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'totalNum', title: '发放总量', slots: { header: () => h('span', {}, ['发放总量 ', h(Tooltip, { title: '0表示不限' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'usedNum', title: '已使用数量' },
    { field: 'claimNum', title: '已领取数量' },
    { field: 'perLimit', title: '每人限领张数' },
    { field: 'sort', title: '排序', slots: { header: () => h('span', {}, ['排序 ', h(Tooltip, { title: '升序' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'status', title: '状态', width: 120, slots: { default: 'status_cell' } },
    { field: 'validStartAt', title: '有效期开始时间', width: 180, formatter: 'formatDateTime' },
    { field: 'validEndAt', title: '有效期结束时间', width: 180, formatter: 'formatDateTime' },
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
    content: '确定要删除该优惠券模板表吗？',
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
