<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getOauthList, deleteOauth } from '#/api/play/oauth';
import type { OauthItem } from '#/api/play/oauth/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 第三方平台选项 */
const providerOptions = [
  { label: '微信', value: 1 },
  { label: '支付宝', value: 2 },
];

/** 第三方平台映射 */
const providerMap: Record<number, string> = {
  1: '微信',
  2: '支付宝',
};

/** 第三方平台颜色 */
function getProviderColor(val: number): string {
  const keys = [1, 2];
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
        options: providerOptions,
        placeholder: '请选择第三方平台',
        class: 'w-full',
      },
      fieldName: 'provider',
      label: '第三方平台',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<OauthItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'memberID', title: '会员ID' },
    { field: 'provider', title: '第三方平台', width: 120, slots: { default: 'provider_cell' } },
    { field: 'openID', title: '第三方OpenID' },
    { field: 'unionID', title: '第三方UnionID' },
    { field: 'nickname', title: '第三方昵称' },
    { field: 'avatar', title: '第三方头像' },
    { field: 'accessToken', title: '访问令牌' },
    { field: 'refreshToken', title: '刷新令牌' },
    { field: 'expireAt', title: '令牌过期时间', width: 180, formatter: 'formatDateTime' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getOauthList({
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
function handleEdit(row: OauthItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: OauthItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该第三方登录绑定表吗？',
    okType: 'danger',
    async onOk() {
      await deleteOauth(row.id);
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
      <template #provider_cell="{ row }">
        <Tag :color="getProviderColor(row.provider)">
          {{ providerMap[row.provider] || row.provider }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
