<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getActivityStepLogList, deleteActivityStepLog } from '#/api/play/activity_step_log';
import type { ActivityStepLogItem } from '#/api/play/activity_step_log/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 步骤类型选项 */
const stepTypeOptions = [
  { label: '文字', value: 1 },
  { label: '链接', value: 2 },
  { label: '图片', value: 3 },
];

/** 步骤类型映射 */
const stepTypeMap: Record<number, string> = {
  1: '文字',
  2: '链接',
  3: '图片',
};

/** 步骤类型颜色 */
function getStepTypeColor(val?: number): string {
  if (val == null) {
    return 'default';
  }
  const keys = [1, 2, 3];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 审核状态选项 */
const auditStatusOptions = [
  { label: '待审核', value: 0 },
  { label: '通过', value: 1 },
  { label: '驳回', value: 2 },
];

/** 审核状态映射 */
const auditStatusMap: Record<number, string> = {
  0: '待审核',
  1: '通过',
  2: '驳回',
};

/** 审核状态颜色 */
function getAuditStatusColor(val?: number): string {
  if (val == null) {
    return 'default';
  }
  const keys = [0, 1, 2];
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
        options: stepTypeOptions,
        placeholder: '请选择步骤类型',
        class: 'w-full',
      },
      fieldName: 'stepType',
      label: '步骤类型',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: auditStatusOptions,
        placeholder: '请选择审核状态',
        class: 'w-full',
      },
      fieldName: 'auditStatus',
      label: '审核状态',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<ActivityStepLogItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'activityTitle', title: '活动ID' },
    { field: 'stepID', title: '步骤ID' },
    { field: 'joinID', title: '参与记录ID' },
    { field: 'memberID', title: '会员ID' },
    { field: 'stepType', title: '步骤类型', width: 120, slots: { default: 'stepType_cell' } },
    { field: 'submitText', title: '用户提交的文字或链接' },
    { field: 'submitImage', title: '用户提交的图片URL', width: 100, slots: { default: 'submitImage_cell' } },
    { field: 'auditStatus', title: '审核状态', width: 120, slots: { default: 'auditStatus_cell' } },
    { field: 'auditRemark', title: '审核备注' },
    { field: 'auditBy', title: '审核人ID' },
    { field: 'auditAt', title: '审核时间', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getActivityStepLogList({
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
function handleEdit(row: ActivityStepLogItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: ActivityStepLogItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该活动步骤提交记录吗？',
    okType: 'danger',
    async onOk() {
      await deleteActivityStepLog(row.id);
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
      <template #stepType_cell="{ row }">
        <Tag :color="getStepTypeColor(row.stepType)">
          {{ row.stepType == null ? '-' : (stepTypeMap[row.stepType] ?? row.stepType) }}
        </Tag>
      </template>
      <template #submitImage_cell="{ row }">
        <img v-if="row.submitImage" :src="row.submitImage" style="width: 48px; height: 48px; object-fit: cover; border-radius: 4px;" />
        <span v-else>-</span>
      </template>
      <template #auditStatus_cell="{ row }">
        <Tag :color="getAuditStatusColor(row.auditStatus)">
          {{ row.auditStatus == null ? '-' : (auditStatusMap[row.auditStatus] ?? row.auditStatus) }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
