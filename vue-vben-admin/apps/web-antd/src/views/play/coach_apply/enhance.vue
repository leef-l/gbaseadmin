<script setup lang="ts">
import { h } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, Input, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getCoachApplyList, deleteCoachApply } from '#/api/play/coach_apply';
import { auditCoachApply } from '#/api/play/coach_apply/enhance';
import type { CoachApplyItem } from '#/api/play/coach_apply/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 审核状态选项 */
const auditStatusOptions = [
  { label: '待审核', value: 0 },
  { label: '通过', value: 1 },
  { label: '拒绝', value: 2 },
];

/** 审核状态映射 */
const auditStatusMap: Record<number, string> = {
  0: '待审核',
  1: '通过',
  2: '拒绝',
};

/** 审核状态颜色 */
function getAuditStatusColor(val: number): string {
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
const gridOptions: VxeGridProps<CoachApplyItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'memberNickname', title: '申请会员' },
    { field: 'realName', title: '真实姓名' },
    { field: 'idCard', title: '身份证号' },
    { field: 'idCardFrontImage', title: '身份证正面照' },
    { field: 'idCardBackImage', title: '身份证反面照' },
    { field: 'skillDesc', title: '技能描述' },
    { field: 'auditStatus', title: '审核状态', width: 120, slots: { default: 'auditStatus_cell' } },
    { field: 'auditRemark', title: '审核备注' },
    { field: 'auditAt', title: '审核时间', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getCoachApplyList({
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
function handleEdit(row: CoachApplyItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: CoachApplyItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该陪玩师申请表吗？',
    okType: 'danger',
    async onOk() {
      await deleteCoachApply(row.id);
      message.success('删除成功');
      gridApi.reload();
    },
  });
}

/** 审核通过 */
function handleApprove(row: CoachApplyItem) {
  Modal.confirm({
    title: '审核通过',
    content: `确定要通过 ${row.realName} 的陪玩师申请吗？`,
    async onOk() {
      await auditCoachApply({ id: row.id, auditStatus: 1 });
      message.success('审核通过');
      gridApi.reload();
    },
  });
}

/** 审核拒绝（需填写原因） */
function handleReject(row: CoachApplyItem) {
  let auditRemark = '';
  Modal.confirm({
    title: '审核拒绝',
    content: () => h('div', [
      h('p', `确定要拒绝 ${row.realName} 的陪玩师申请吗？`),
      h(Input.TextArea, {
        placeholder: '请输入拒绝原因',
        rows: 3,
        onChange: (e: any) => { auditRemark = e.target.value; },
      }),
    ]),
    okType: 'danger',
    async onOk() {
      if (!auditRemark.trim()) {
        message.warning('请输入拒绝原因');
        throw new Error('请输入拒绝原因');
      }
      await auditCoachApply({ id: row.id, auditStatus: 2, auditRemark });
      message.success('已拒绝');
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
      <template #auditStatus_cell="{ row }">
        <Tag :color="getAuditStatusColor(row.auditStatus)">
          {{ auditStatusMap[row.auditStatus] || row.auditStatus }}
        </Tag>
      </template>
      <template #action="{ row }">
        <template v-if="row.auditStatus === 0">
          <Button type="link" size="small" @click="handleApprove(row)">通过</Button>
          <Button type="link" danger size="small" @click="handleReject(row)">拒绝</Button>
        </template>
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
