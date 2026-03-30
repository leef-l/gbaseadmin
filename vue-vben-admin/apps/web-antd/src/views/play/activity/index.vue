<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getActivityList, deleteActivity } from '#/api/play/activity';
import type { ActivityItem } from '#/api/play/activity/types';
import FormModal from './modules/form.vue';
import DetailDrawer from './modules/detail-drawer.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 活动类型选项 */
const typeOptions = [
  { label: '充值活动', value: 1 },
  { label: '下单活动', value: 2 },
  { label: '注册活动', value: 3 },
  { label: '图文步骤活动', value: 4 },
  { label: '自定义活动', value: 5 },
];

/** 活动类型映射 */
const typeMap: Record<number, string> = {
  1: '充值活动',
  2: '下单活动',
  3: '注册活动',
  4: '图文步骤活动',
  5: '自定义活动',
};

/** 活动类型颜色 */
function getTypeColor(val: number): string {
  const keys = [1, 2, 3, 4, 5];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 参与条件选项 */
const conditionTypeOptions = [
  { label: '无条件', value: 0 },
  { label: '需报名', value: 1 },
  { label: '充值满额', value: 2 },
  { label: '下单满额', value: 3 },
  { label: '完成步骤', value: 4 },
];

/** 参与条件映射 */
const conditionTypeMap: Record<number, string> = {
  0: '无条件',
  1: '需报名',
  2: '充值满额',
  3: '下单满额',
  4: '完成步骤',
};

/** 参与条件颜色 */
function getConditionTypeColor(val: number): string {
  const keys = [0, 1, 2, 3, 4];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 是否自动发奖选项 */
const isAutoRewardOptions = [
  { label: '否（需审核）', value: 0 },
  { label: '是（用户完成即发）', value: 1 },
];

/** 是否自动发奖映射 */
const isAutoRewardMap: Record<number, string> = {
  0: '否（需审核）',
  1: '是（用户完成即发）',
};

/** 是否自动发奖颜色 */
function getIsAutoRewardColor(val: number): string {
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

/** 详情弹窗 */
const [DetailDrawerComp, detailDrawerApi] = useVbenModal({
  connectedComponent: DetailDrawer,
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
        placeholder: '请选择活动类型',
        class: 'w-full',
      },
      fieldName: 'type',
      label: '活动类型',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: conditionTypeOptions,
        placeholder: '请选择参与条件',
        class: 'w-full',
      },
      fieldName: 'conditionType',
      label: '参与条件',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: isAutoRewardOptions,
        placeholder: '请选择是否自动发奖',
        class: 'w-full',
      },
      fieldName: 'isAutoReward',
      label: '是否自动发奖',
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
const gridOptions: VxeGridProps<ActivityItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'title', title: '活动名称' },
    { field: 'coverImage', title: '活动封面图' },
    { field: 'descContent', title: '活动详情描述（富文本，支持图文混排）' },
    { field: 'type', title: '活动类型', width: 120, slots: { default: 'type_cell' } },
    { field: 'conditionType', title: '参与条件', width: 120, slots: { default: 'conditionType_cell' } },
    { field: 'conditionValue', title: '条件值（分/次，如充值满5000分、下单满3次）' },
    { field: 'isAutoReward', title: '是否自动发奖', width: 120, slots: { default: 'isAutoReward_cell' } },
    { field: 'maxNum', title: '参与人数上限（0表示不限）' },
    { field: 'joinNum', title: '已参与人数' },
    { field: 'sort', title: '排序（升序）' },
    { field: 'status', title: '状态', width: 120, slots: { default: 'status_cell' } },
    { field: 'startAt', title: '活动开始时间', width: 180, formatter: 'formatDateTime' },
    { field: 'endAt', title: '活动结束时间', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 280, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getActivityList({
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
function handleEdit(row: ActivityItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: ActivityItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该活动表吗？',
    okType: 'danger',
    async onOk() {
      await deleteActivity(row.id);
      message.success('删除成功');
      gridApi.reload();
    },
  });
}

/** 查看详情（奖励+步骤） */
function handleDetail(row: ActivityItem) {
  detailDrawerApi.setData({ id: row.id, title: row.title, type: row.type, tab: 'rewards' }).open();
}

/** 管理奖励 */
function handleRewards(row: ActivityItem) {
  detailDrawerApi.setData({ id: row.id, title: row.title, type: row.type, tab: 'rewards' }).open();
}

/** 管理步骤 */
function handleSteps(row: ActivityItem) {
  detailDrawerApi.setData({ id: row.id, title: row.title, type: row.type, tab: 'steps' }).open();
}
</script>

<template>
  <Page auto-content-height>
    <FormModalComp @success="() => gridApi.reload()" />
    <DetailDrawerComp />
    <Grid>
      <template #toolbar-actions>
        <Button type="primary" @click="handleCreate">新建</Button>
      </template>
      <template #type_cell="{ row }">
        <Tag :color="getTypeColor(row.type)">
          {{ typeMap[row.type] || row.type }}
        </Tag>
      </template>
      <template #conditionType_cell="{ row }">
        <Tag :color="getConditionTypeColor(row.conditionType)">
          {{ conditionTypeMap[row.conditionType] || row.conditionType }}
        </Tag>
      </template>
      <template #isAutoReward_cell="{ row }">
        <Tag :color="getIsAutoRewardColor(row.isAutoReward)">
          {{ isAutoRewardMap[row.isAutoReward] || row.isAutoReward }}
        </Tag>
      </template>
      <template #status_cell="{ row }">
        <Tag :color="getStatusColor(row.status)">
          {{ statusMap[row.status] || row.status }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleDetail(row)">详情</Button>
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
