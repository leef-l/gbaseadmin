<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';
import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getActivityJoinList, deleteActivityJoin } from '#/api/play/activity_join';
import type { ActivityJoinItem } from '#/api/play/activity_join/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 参与状态选项 */
const joinStatusOptions = [
  { label: '已报名', value: 0 },
  { label: '进行中', value: 1 },
  { label: '已完成', value: 2 },
  { label: '已领奖', value: 3 },
];

/** 参与状态映射 */
const joinStatusMap: Record<number, string> = {
  0: '已报名',
  1: '进行中',
  2: '已完成',
  3: '已领奖',
};

/** 参与状态颜色 */
function getJoinStatusColor(val: number): string {
  const keys = [0, 1, 2, 3];
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
        options: joinStatusOptions,
        placeholder: '请选择参与状态',
        class: 'w-full',
      },
      fieldName: 'joinStatus',
      label: '参与状态',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<ActivityJoinItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'activityTitle', title: '活动ID' },
    { field: 'memberID', title: '会员ID' },
    { field: 'joinStatus', title: '参与状态', width: 120, slots: { default: 'joinStatus_cell' } },
    { field: 'currentStep', title: '当前步骤', slots: { header: () => h('span', {}, ['当前步骤 ', h(Tooltip, { title: '步骤活动用' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'remark', title: '备注' },
    { field: 'finishAt', title: '完成时间', width: 180, formatter: 'formatDateTime' },
    { field: 'rewardAt', title: '领奖时间', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getActivityJoinList({
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
function handleEdit(row: ActivityJoinItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: ActivityJoinItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该活动参与记录表吗？',
    okType: 'danger',
    async onOk() {
      await deleteActivityJoin(row.id);
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
      <template #joinStatus_cell="{ row }">
        <Tag :color="getJoinStatusColor(row.joinStatus)">
          {{ joinStatusMap[row.joinStatus] || row.joinStatus }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
