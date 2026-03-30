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

/** ç¬¬ä¸‰æ–¹å¹³å°选项 */
const providerOptions = [
  { label: 'å¾®ä¿¡', value: 1 },
  { label: 'æ”¯ä»˜å®', value: 2 },
];

/** ç¬¬ä¸‰æ–¹å¹³å°映射 */
const providerMap: Record<number, string> = {
  1: 'å¾®ä¿¡',
  2: 'æ”¯ä»˜å®',
};

/** ç¬¬ä¸‰æ–¹å¹³å°颜色 */
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
        placeholder: '请选择ç¬¬ä¸‰æ–¹å¹³å°',
        class: 'w-full',
      },
      fieldName: 'provider',
      label: 'ç¬¬ä¸‰æ–¹å¹³å°',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<OauthItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'memberID', title: 'ä¼šå‘˜ID' },
    { field: 'provider', title: 'ç¬¬ä¸‰æ–¹å¹³å°', width: 120, slots: { default: 'provider_cell' } },
    { field: 'openID', title: 'ç¬¬ä¸‰æ–¹OpenID' },
    { field: 'unionID', title: 'ç¬¬ä¸‰æ–¹UnionID' },
    { field: 'nickname', title: 'ç¬¬ä¸‰æ–¹æ˜µç§°' },
    { field: 'avatar', title: 'ç¬¬ä¸‰æ–¹å¤´åƒ' },
    { field: 'accessToken', title: 'è®¿é—®ä»¤ç‰Œ' },
    { field: 'refreshToken', title: 'åˆ·æ–°ä»¤ç‰Œ' },
    { field: 'expireAt', title: 'ä»¤ç‰Œè¿‡æœŸæ—¶é—´', width: 180, formatter: 'formatDateTime' },
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
    content: '确定要删除该ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨吗？',
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
