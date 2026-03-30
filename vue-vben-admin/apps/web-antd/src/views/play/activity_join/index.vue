<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getActivityJoinList, deleteActivityJoin } from '#/api/play/activity_join';
import type { ActivityJoinItem } from '#/api/play/activity_join/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** å‚ä¸ŽçŠ¶æ€选项 */
const joinStatusOptions = [
  { label: 'å·²æŠ¥å', value: 0 },
  { label: 'è¿›è¡Œä¸­', value: 1 },
  { label: 'å·²å®Œæˆ', value: 2 },
  { label: 'å·²é¢†å¥–', value: 3 },
];

/** å‚ä¸ŽçŠ¶æ€映射 */
const joinStatusMap: Record<number, string> = {
  0: 'å·²æŠ¥å',
  1: 'è¿›è¡Œä¸­',
  2: 'å·²å®Œæˆ',
  3: 'å·²é¢†å¥–',
};

/** å‚ä¸ŽçŠ¶æ€颜色 */
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
        placeholder: '请选择å‚ä¸ŽçŠ¶æ€',
        class: 'w-full',
      },
      fieldName: 'joinStatus',
      label: 'å‚ä¸ŽçŠ¶æ€',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<ActivityJoinItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'activityTitle', title: 'æ´»åŠ¨ID' },
    { field: 'memberID', title: 'ä¼šå‘˜ID' },
    { field: 'joinStatus', title: 'å‚ä¸ŽçŠ¶æ€', width: 120, slots: { default: 'joinStatus_cell' } },
    { field: 'currentStep', title: 'å½“å‰å®Œæˆåˆ°ç¬¬å‡ æ­¥' },
    { field: 'remark', title: 'å¤‡æ³¨' },
    { field: 'finishAt', title: 'å®Œæˆæ—¶é—´', width: 180, formatter: 'formatDateTime' },
    { field: 'rewardAt', title: 'é¢†å¥–æ—¶é—´', width: 180, formatter: 'formatDateTime' },
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
    content: '确定要删除该æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨吗？',
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
