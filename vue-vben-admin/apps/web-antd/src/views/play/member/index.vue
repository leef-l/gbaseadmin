<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getMemberList, deleteMember } from '#/api/play/member';
import type { MemberItem } from '#/api/play/member/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 性别选项 */
const genderOptions = [
  { label: '未知', value: 0 },
  { label: '男', value: 1 },
  { label: '女', value: 2 },
];

/** 性别映射 */
const genderMap: Record<number, string> = {
  0: '未知',
  1: '男',
  2: '女',
};

/** 性别颜色 */
function getGenderColor(val: number): string {
  const keys = [0, 1, 2];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 是否陪玩师选项 */
const isCoachOptions = [
  { label: '否', value: 0 },
  { label: '是', value: 1 },
];

/** 是否陪玩师映射 */
const isCoachMap: Record<number, string> = {
  0: '否',
  1: '是',
};

/** 是否陪玩师颜色 */
function getIsCoachColor(val: number): string {
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
        options: genderOptions,
        placeholder: '请选择性别',
        class: 'w-full',
      },
      fieldName: 'gender',
      label: '性别',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: isCoachOptions,
        placeholder: '请选择是否陪玩师',
        class: 'w-full',
      },
      fieldName: 'isCoach',
      label: '是否陪玩师',
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
const gridOptions: VxeGridProps<MemberItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'phone', title: '手机号' },
    { field: 'nickname', title: '昵称' },
    { field: 'avatar', title: '头像', width: 80, slots: { default: 'avatar_cell' } },
    { field: 'gender', title: '性别', width: 120, slots: { default: 'gender_cell' } },
    { field: 'memberLevelTitle', title: '会员等级' },
    { field: 'exp', title: '经验值' },
    { field: 'balance', title: '余额（元）', formatter: ({ cellValue }: { cellValue: number | null }) => cellValue != null ? (cellValue / 100).toFixed(2) : '-' },
    { field: 'isCoach', title: '是否陪玩师', width: 120, slots: { default: 'isCoach_cell' } },
    { field: 'status', title: '状态', width: 120, slots: { default: 'status_cell' } },
    { field: 'lastLoginAt', title: '最后登录时间', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getMemberList({
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
function handleEdit(row: MemberItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: MemberItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该会员表吗？',
    okType: 'danger',
    async onOk() {
      await deleteMember(row.id);
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
      <template #avatar_cell="{ row }">
        <img v-if="row.avatar" :src="row.avatar" style="width:40px;height:40px;object-fit:cover;border-radius:4px;" />
        <span v-else>-</span>
      </template>
      <template #gender_cell="{ row }">
        <Tag :color="getGenderColor(row.gender)">
          {{ genderMap[row.gender] || row.gender }}
        </Tag>
      </template>
      <template #isCoach_cell="{ row }">
        <Tag :color="getIsCoachColor(row.isCoach)">
          {{ isCoachMap[row.isCoach] || row.isCoach }}
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
