<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getGoodsList, deleteGoods } from '#/api/play/goods';
import type { GoodsItem } from '#/api/play/goods/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** çŠ¶æ€选项 */
const statusOptions = [
  { label: 'ä¸‹æž¶', value: 0 },
  { label: 'ä¸Šæž¶', value: 1 },
];

/** çŠ¶æ€映射 */
const statusMap: Record<number, string> = {
  0: 'ä¸‹æž¶',
  1: 'ä¸Šæž¶',
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
const gridOptions: VxeGridProps<GoodsItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'categoryTitle', title: 'åˆ†ç±»ID' },
    { field: 'coachID', title: 'é™ªçŽ©å¸ˆID' },
    { field: 'title', title: 'å•†å“åç§°' },
    { field: 'coverImage', title: 'å•†å“å°é¢å›¾' },
    { field: 'descContent', title: 'å•†å“è¯¦æƒ…æè¿°' },
    { field: 'price', title: 'å•ä»·ï¼ˆåˆ†ï¼‰' },
    { field: 'unit', title: 'è®¡é‡å•ä½' },
    { field: 'salesNum', title: 'é”€é‡' },
    { field: 'sort', title: 'æŽ’åºï¼ˆå‡åºï¼‰' },
    { field: 'status', title: 'çŠ¶æ€', width: 120, slots: { default: 'status_cell' } },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getGoodsList({
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
function handleEdit(row: GoodsItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: GoodsItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该å•†å“è¡¨吗？',
    okType: 'danger',
    async onOk() {
      await deleteGoods(row.id);
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
