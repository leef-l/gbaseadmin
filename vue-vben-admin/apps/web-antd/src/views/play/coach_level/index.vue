<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';
import { Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getCoachLevelList, deleteCoachLevel } from '#/api/play/coach_level';
import type { CoachLevelItem } from '#/api/play/coach_level/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 等级选项 */
const levelOptions = [
  { label: '青铜', value: 1 },
  { label: '白银', value: 2 },
  { label: '黄金', value: 3 },
  { label: '铂金', value: 4 },
  { label: '钻石', value: 5 },
];

/** 等级映射 */
const levelMap: Record<number, string> = {
  1: '青铜',
  2: '白银',
  3: '黄金',
  4: '铂金',
  5: '钻石',
};

/** 等级颜色 */
function getLevelColor(val: number): string {
  const keys = [1, 2, 3, 4, 5];
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
        options: levelOptions,
        placeholder: '请选择等级',
        class: 'w-full',
      },
      fieldName: 'level',
      label: '等级',
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
const gridOptions: VxeGridProps<CoachLevelItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'title', title: '等级名称' },
    { field: 'level', title: '等级', width: 120, slots: { default: 'level_cell' } },
    { field: 'icon', title: '等级图标' },
    { field: 'minOrders', title: '所需最低接单数' },
    { field: 'minScore', title: '最低评分', slots: { header: () => h('span', {}, ['最低评分 ', h(Tooltip, { title: '乘100存储，如 450=4.50分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'commissionRate', title: '抽成比例', slots: { header: () => h('span', {}, ['抽成比例 ', h(Tooltip, { title: '百分比，如 20 表示 20%' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'sort', title: '排序', slots: { header: () => h('span', {}, ['排序 ', h(Tooltip, { title: '升序' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'status', title: '状态', width: 120, slots: { default: 'status_cell' } },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getCoachLevelList({
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
function handleEdit(row: CoachLevelItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: CoachLevelItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该陪玩师等级表吗？',
    okType: 'danger',
    async onOk() {
      await deleteCoachLevel(row.id);
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
      <template #level_cell="{ row }">
        <Tag :color="getLevelColor(row.level)">
          {{ levelMap[row.level] || row.level }}
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
