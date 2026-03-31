<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getActivityRewardList, deleteActivityReward } from '#/api/play/activity_reward';
import type { ActivityRewardItem } from '#/api/play/activity_reward/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 奖励类型选项 */
const rewardTypeOptions = [
  { label: '余额', value: 1 },
  { label: '优惠券', value: 2 },
  { label: '经验值', value: 3 },
  { label: '会员等级天数', value: 4 },
];

/** 奖励类型映射 */
const rewardTypeMap: Record<number, string> = {
  1: '余额',
  2: '优惠券',
  3: '经验值',
  4: '会员等级天数',
};

/** 奖励类型颜色 */
function getRewardTypeColor(val: number): string {
  const keys = [1, 2, 3, 4];
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
        options: rewardTypeOptions,
        placeholder: '请选择奖励类型',
        class: 'w-full',
      },
      fieldName: 'rewardType',
      label: '奖励类型',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<ActivityRewardItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'activityTitle', title: '活动ID' },
    { field: 'rewardType', title: '奖励类型', width: 120, slots: { default: 'rewardType_cell' } },
    { field: 'rewardValue', title: '奖励数值', slots: { header: () => h('span', {}, ['奖励数值 ', h(Tooltip, { title: '余额=分，优惠券=coupon_id，经验=值，等级天数=天' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'rewardName', title: '奖励名称', slots: { header: () => h('span', {}, ['奖励名称 ', h(Tooltip, { title: '展示用，如"送50元余额"' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'sort', title: '排序', slots: { header: () => h('span', {}, ['排序 ', h(Tooltip, { title: '升序' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) } },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getActivityRewardList({
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
function handleEdit(row: ActivityRewardItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: ActivityRewardItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该活动奖励表吗？',
    okType: 'danger',
    async onOk() {
      await deleteActivityReward(row.id);
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
      <template #rewardType_cell="{ row }">
        <Tag :color="getRewardTypeColor(row.rewardType)">
          {{ rewardTypeMap[row.rewardType] || row.rewardType }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
