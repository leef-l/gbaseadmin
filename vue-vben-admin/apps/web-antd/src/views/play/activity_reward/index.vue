<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getActivityRewardList, deleteActivityReward } from '#/api/play/activity_reward';
import type { ActivityRewardItem } from '#/api/play/activity_reward/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** å¥–åŠ±ç±»åž‹选项 */
const rewardTypeOptions = [
  { label: 'ä½™é¢', value: 1 },
  { label: 'ä¼˜æƒ åˆ¸', value: 2 },
  { label: 'ç»éªŒå€¼', value: 3 },
  { label: 'ä¼šå‘˜ç­‰çº§å¤©æ•°', value: 4 },
];

/** å¥–åŠ±ç±»åž‹映射 */
const rewardTypeMap: Record<number, string> = {
  1: 'ä½™é¢',
  2: 'ä¼˜æƒ åˆ¸',
  3: 'ç»éªŒå€¼',
  4: 'ä¼šå‘˜ç­‰çº§å¤©æ•°',
};

/** å¥–åŠ±ç±»åž‹颜色 */
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
        placeholder: '请选择å¥–åŠ±ç±»åž‹',
        class: 'w-full',
      },
      fieldName: 'rewardType',
      label: 'å¥–åŠ±ç±»åž‹',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<ActivityRewardItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'activityTitle', title: 'æ´»åŠ¨ID' },
    { field: 'rewardType', title: 'å¥–åŠ±ç±»åž‹', width: 120, slots: { default: 'rewardType_cell' } },
    { field: 'rewardValue', title: 'å¥–åŠ±æ•°å€¼' },
    { field: 'rewardName', title: 'å¥–åŠ±åç§°' },
    { field: 'sort', title: 'æŽ’åº' },
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
    content: '确定要删除该æ´»åŠ¨å¥–åŠ±è¡¨吗？',
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
