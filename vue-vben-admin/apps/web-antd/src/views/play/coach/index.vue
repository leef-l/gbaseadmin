<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getCoachList, deleteCoach } from '#/api/play/coach';
import type { CoachItem } from '#/api/play/coach/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 是否在线选项 */
const isOnlineOptions = [
  { label: '离线', value: 0 },
  { label: '在线', value: 1 },
];

/** 是否在线映射 */
const isOnlineMap: Record<number, string> = {
  0: '离线',
  1: '在线',
};

/** 是否在线颜色 */
function getIsOnlineColor(val: number): string {
  const keys = [0, 1];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 状态选项 */
const statusOptions = [
  { label: '禁用', value: 0 },
  { label: '正常', value: 1 },
];

/** 状态映射 */
const statusMap: Record<number, string> = {
  0: '禁用',
  1: '正常',
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
        options: isOnlineOptions,
        placeholder: '请选择是否在线',
        class: 'w-full',
      },
      fieldName: 'isOnline',
      label: '是否在线',
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
const gridOptions: VxeGridProps<CoachItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'memberID', title: '关联会员ID' },
    { field: 'coachLevelTitle', title: '陪玩师等级' },
    { field: 'shopTitle', title: '所属店铺', slots: { header: () => h('span', {}, ['所属店铺 ', h(Tooltip, { title: '0表示无店铺' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'realName', title: '真实姓名' },
    { field: 'intro', title: '个人简介' },
    { field: 'coverImage', title: '封面图', width: 100, slots: { default: 'coverImage_cell' } },
    { field: 'totalOrders', title: '总接单数' },
    { field: 'totalScore', title: '总评分', slots: { header: () => h('span', {}, ['总评分 ', h(Tooltip, { title: '乘100，如 500=5.00' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'scoreNum', title: '评分人数' },
    { field: 'incomeTotal', title: '累计收入', slots: { header: () => h('span', {}, ['累计收入 ', h(Tooltip, { title: '单位：分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'incomeBalance', title: '可提现余额', slots: { header: () => h('span', {}, ['可提现余额 ', h(Tooltip, { title: '单位：分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'isOnline', title: '是否在线', width: 120, slots: { default: 'isOnline_cell' } },
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
        const res = await getCoachList({
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
function handleEdit(row: CoachItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: CoachItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该陪玩师表吗？',
    okType: 'danger',
    async onOk() {
      await deleteCoach(row.id);
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
      <template #coverImage_cell="{ row }">
        <img v-if="row.coverImage" :src="row.coverImage" style="width:48px;height:48px;object-fit:cover;border-radius:4px;" />
        <span v-else>-</span>
      </template>
      <template #isOnline_cell="{ row }">
        <Tag :color="getIsOnlineColor(row.isOnline)">
          {{ isOnlineMap[row.isOnline] || row.isOnline }}
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
