<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysTestTree, deleteSysTest } from '#/api/system/sys_test';
import type { SysTestItem } from '#/api/system/sys_test/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** ç±»åž‹选项 */
const typeOptions = [
  { label: 'æ™®é€š', value: 1 },
  { label: 'ç‰¹æ®Š', value: 2 },
  { label: 'é«˜çº§', value: 3 },
];

/** ç±»åž‹映射 */
const typeMap: Record<number, string> = {
  1: 'æ™®é€š',
  2: 'ç‰¹æ®Š',
  3: 'é«˜çº§',
};

/** ç±»åž‹颜色 */
function getTypeColor(val: number): string {
  const keys = [1, 2, 3];
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
        placeholder: '请选择ç±»åž‹',
        class: 'w-full',
      },
      fieldName: 'type',
      label: 'ç±»åž‹',
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
const gridOptions: VxeGridProps<SysTestItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'title', title: 'åç§°', treeNode: true },
    { field: 'code', title: 'ç¼–ç ' },
    { field: 'type', title: 'ç±»åž‹', width: 120, slots: { default: 'type_cell' } },
    { field: 'status', title: 'çŠ¶æ€', width: 120, slots: { default: 'status_cell' } },
    { field: 'sort', title: 'æŽ’åº' },
    { field: 'remark', title: 'å¤‡æ³¨' },
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
        return await getSysTestTree(formValues) ?? [];
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
function handleEdit(row: SysTestItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: SysTestItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该æµ‹è¯•è¡¨吗？',
    okType: 'danger',
    async onOk() {
      await deleteSysTest(row.id);
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
