<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getCategoryTree, deleteCategory } from '#/api/play/category';
import type { CategoryItem } from '#/api/play/category/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

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
const gridOptions: VxeGridProps<CategoryItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'title', title: '分类名称', treeNode: true },
    { field: 'icon', title: '分类图标' },
    { field: 'coverImage', title: '分类封面图' },
    { field: 'sort', title: '排序', slots: { header: () => h('span', {}, ['排序 ', h(Tooltip, { title: '升序' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'status', title: '状态', width: 120, slots: { default: 'status_cell' } },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  pagerConfig: { enabled: false },
  treeConfig: {
    childrenField: 'children',
    expandAll: false,
  },
  proxyConfig: {
    ajax: {
      query: async (_params, formValues) => {
        return await getCategoryTree(formValues) ?? [];
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
function handleEdit(row: CategoryItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: CategoryItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该商品分类表吗？',
    okType: 'danger',
    async onOk() {
      await deleteCategory(row.id);
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
