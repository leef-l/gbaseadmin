<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getReviewList, deleteReview } from '#/api/play/review';
import type { ReviewItem } from '#/api/play/review/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** æ˜¯å¦åŒ¿å选项 */
const isAnonymousOptions = [
  { label: 'å¦', value: 0 },
  { label: 'æ˜¯', value: 1 },
];

/** æ˜¯å¦åŒ¿å映射 */
const isAnonymousMap: Record<number, string> = {
  0: 'å¦',
  1: 'æ˜¯',
};

/** æ˜¯å¦åŒ¿å颜色 */
function getIsAnonymousColor(val: number): string {
  const keys = [0, 1];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** çŠ¶æ€选项 */
const statusOptions = [
  { label: 'éšè—', value: 0 },
  { label: 'æ˜¾ç¤º', value: 1 },
];

/** çŠ¶æ€映射 */
const statusMap: Record<number, string> = {
  0: 'éšè—',
  1: 'æ˜¾ç¤º',
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
        options: isAnonymousOptions,
        placeholder: '请选择æ˜¯å¦åŒ¿å',
        class: 'w-full',
      },
      fieldName: 'isAnonymous',
      label: 'æ˜¯å¦åŒ¿å',
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
const gridOptions: VxeGridProps<ReviewItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'orderID', title: 'è®¢å•ID' },
    { field: 'memberID', title: 'è¯„ä»·ä¼šå‘˜ID' },
    { field: 'coachID', title: 'è¢«è¯„é™ªçŽ©å¸ˆID' },
    { field: 'score', title: 'è¯„åˆ†ï¼ˆä¹˜100ï¼‰' },
    { field: 'reviewContent', title: 'è¯„ä»·å†…å®¹' },
    { field: 'reviewImage', title: 'è¯„ä»·å›¾ç‰‡ï¼ˆå¤šå¼ é€—å·åˆ†éš”ï¼‰' },
    { field: 'replyContent', title: 'é™ªçŽ©å¸ˆå›žå¤å†…å®¹' },
    { field: 'isAnonymous', title: 'æ˜¯å¦åŒ¿å', width: 120, slots: { default: 'isAnonymous_cell' } },
    { field: 'status', title: 'çŠ¶æ€', width: 120, slots: { default: 'status_cell' } },
    { field: 'replyAt', title: 'å›žå¤æ—¶é—´', width: 180, formatter: 'formatDateTime' },
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
    content: '确定要删除该è¯„ä»·è¡¨吗？',
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
