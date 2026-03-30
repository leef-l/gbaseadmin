<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getActivityList, deleteActivity } from '#/api/play/activity';
import type { ActivityItem } from '#/api/play/activity/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** æ´»åŠ¨ç±»åž‹选项 */
const typeOptions = [
  { label: 'å……å€¼æ´»åŠ¨', value: 1 },
  { label: 'ä¸‹å•æ´»åŠ¨', value: 2 },
  { label: 'æ³¨å†Œæ´»åŠ¨', value: 3 },
  { label: 'å›¾æ–‡æ­¥éª¤æ´»åŠ¨', value: 4 },
  { label: 'è‡ªå®šä¹‰æ´»åŠ¨', value: 5 },
];

/** æ´»åŠ¨ç±»åž‹映射 */
const typeMap: Record<number, string> = {
  1: 'å……å€¼æ´»åŠ¨',
  2: 'ä¸‹å•æ´»åŠ¨',
  3: 'æ³¨å†Œæ´»åŠ¨',
  4: 'å›¾æ–‡æ­¥éª¤æ´»åŠ¨',
  5: 'è‡ªå®šä¹‰æ´»åŠ¨',
};

/** æ´»åŠ¨ç±»åž‹颜色 */
function getTypeColor(val: number): string {
  const keys = [1, 2, 3, 4, 5];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** å‚ä¸Žæ¡ä»¶选项 */
const conditionTypeOptions = [
  { label: 'æ— æ¡ä»¶', value: 0 },
  { label: 'éœ€æŠ¥å', value: 1 },
  { label: 'å……å€¼æ»¡é¢', value: 2 },
  { label: 'ä¸‹å•æ»¡é¢', value: 3 },
  { label: 'å®Œæˆæ­¥éª¤', value: 4 },
];

/** å‚ä¸Žæ¡ä»¶映射 */
const conditionTypeMap: Record<number, string> = {
  0: 'æ— æ¡ä»¶',
  1: 'éœ€æŠ¥å',
  2: 'å……å€¼æ»¡é¢',
  3: 'ä¸‹å•æ»¡é¢',
  4: 'å®Œæˆæ­¥éª¤',
};

/** å‚ä¸Žæ¡ä»¶颜色 */
function getConditionTypeColor(val: number): string {
  const keys = [0, 1, 2, 3, 4];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** æ˜¯å¦è‡ªåŠ¨å‘å¥–选项 */
const isAutoRewardOptions = [
  { label: 'å¦', value: 0 },
  { label: 'æ˜¯', value: 1 },
];

/** æ˜¯å¦è‡ªåŠ¨å‘å¥–映射 */
const isAutoRewardMap: Record<number, string> = {
  0: 'å¦',
  1: 'æ˜¯',
};

/** æ˜¯å¦è‡ªåŠ¨å‘å¥–颜色 */
function getIsAutoRewardColor(val: number): string {
  const keys = [0, 1];
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
        placeholder: '请选择æ´»åŠ¨ç±»åž‹',
        class: 'w-full',
      },
      fieldName: 'type',
      label: 'æ´»åŠ¨ç±»åž‹',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: conditionTypeOptions,
        placeholder: '请选择å‚ä¸Žæ¡ä»¶',
        class: 'w-full',
      },
      fieldName: 'conditionType',
      label: 'å‚ä¸Žæ¡ä»¶',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: isAutoRewardOptions,
        placeholder: '请选择æ˜¯å¦è‡ªåŠ¨å‘å¥–',
        class: 'w-full',
      },
      fieldName: 'isAutoReward',
      label: 'æ˜¯å¦è‡ªåŠ¨å‘å¥–',
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
const gridOptions: VxeGridProps<ActivityItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'title', title: 'æ´»åŠ¨åç§°' },
    { field: 'coverImage', title: 'æ´»åŠ¨å°é¢å›¾' },
    { field: 'descContent', title: 'æ´»åŠ¨è¯¦æƒ…æè¿°' },
    { field: 'type', title: 'æ´»åŠ¨ç±»åž‹', width: 120, slots: { default: 'type_cell' } },
    { field: 'conditionType', title: 'å‚ä¸Žæ¡ä»¶', width: 120, slots: { default: 'conditionType_cell' } },
    { field: 'conditionValue', title: 'æ¡ä»¶å€¼' },
    { field: 'isAutoReward', title: 'æ˜¯å¦è‡ªåŠ¨å‘å¥–', width: 120, slots: { default: 'isAutoReward_cell' } },
    { field: 'maxNum', title: 'å‚ä¸Žäººæ•°ä¸Šé™ï¼ˆ0ä¸é™ï¼‰' },
    { field: 'joinNum', title: 'å·²å‚ä¸Žäººæ•°' },
    { field: 'sort', title: 'æŽ’åº' },
    { field: 'status', title: 'çŠ¶æ€', width: 120, slots: { default: 'status_cell' } },
    { field: 'startAt', title: 'æ´»åŠ¨å¼€å§‹æ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'endAt', title: 'æ´»åŠ¨ç»“æŸæ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
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
    content: '确定要删除该æ´»åŠ¨è¡¨吗？',
    okType: 'danger',
    async onOk() {
      await deleteActivity(row.id);
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
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
