<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getReviewList, deleteReview } from '#/api/play/review';
import type { ReviewItem } from '#/api/play/review/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 是否匿名选项 */
const isAnonymousOptions = [
  { label: '否', value: 0 },
  { label: '是', value: 1 },
];

/** 是否匿名映射 */
const isAnonymousMap: Record<number, string> = {
  0: '否',
  1: '是',
};

/** 是否匿名颜色 */
function getIsAnonymousColor(val: number): string {
  const keys = [0, 1];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 状态选项 */
const statusOptions = [
  { label: '隐藏', value: 0 },
  { label: '显示', value: 1 },
];

/** 状态映射 */
const statusMap: Record<number, string> = {
  0: '隐藏',
  1: '显示',
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
        options: isAnonymousOptions,
        placeholder: '请选择是否匿名',
        class: 'w-full',
      },
      fieldName: 'isAnonymous',
      label: '是否匿名',
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
const gridOptions: VxeGridProps<ReviewItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'orderID', title: '订单ID' },
    { field: 'memberID', title: '评价会员ID' },
    { field: 'coachID', title: '被评陪玩师ID' },
    { field: 'score', title: '评分', slots: { header: () => h('span', {}, ['评分 ', h(Tooltip, { title: '乘100，如 500=5.00分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'reviewContent', title: '评价内容' },
    { field: 'reviewImage', title: '评价图片', width: 120, slots: { default: 'reviewImage_cell', header: () => h('span', {}, ['评价图片 ', h(Tooltip, { title: '多张逗号分隔' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'replyContent', title: '陪玩师回复内容' },
    { field: 'isAnonymous', title: '是否匿名', width: 120, slots: { default: 'isAnonymous_cell' } },
    { field: 'status', title: '状态', width: 120, slots: { default: 'status_cell' } },
    { field: 'replyAt', title: '回复时间', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getReviewList({
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
function handleEdit(row: ReviewItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: ReviewItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该评价表吗？',
    okType: 'danger',
    async onOk() {
      await deleteReview(row.id);
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
      <template #reviewImage_cell="{ row }">
        <div v-if="row.reviewImage" style="display: flex; gap: 4px; flex-wrap: wrap;">
          <img v-for="(img, idx) in row.reviewImage.split(',')" :key="idx" :src="img.trim()" style="width:40px;height:40px;object-fit:cover;border-radius:4px;" />
        </div>
        <span v-else>-</span>
      </template>
      <template #isAnonymous_cell="{ row }">
        <Tag :color="getIsAnonymousColor(row.isAnonymous)">
          {{ isAnonymousMap[row.isAnonymous] || row.isAnonymous }}
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
