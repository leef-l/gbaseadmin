<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getConfigList, deleteConfig } from '#/api/upload/config';
import type { ConfigItem } from '#/api/upload/config/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** å­˜å‚¨ç±»åž‹选项 */
const storageOptions = [
  { label: 'æœ¬åœ°', value: 1 },
  { label: 'é˜¿é‡Œäº‘OSS', value: 2 },
  { label: 'è…¾è®¯äº‘COS', value: 3 },
];

/** å­˜å‚¨ç±»åž‹映射 */
const storageMap: Record<number, string> = {
  1: 'æœ¬åœ°',
  2: 'é˜¿é‡Œäº‘OSS',
  3: 'è…¾è®¯äº‘COS',
};

/** å­˜å‚¨ç±»åž‹颜色 */
function getStorageColor(val: number): string {
  const keys = [1, 2, 3];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** æ˜¯å¦é»˜è®¤选项 */
const isDefaultOptions = [
  { label: 'å¦', value: 0 },
  { label: 'æ˜¯', value: 1 },
];

/** æ˜¯å¦é»˜è®¤映射 */
const isDefaultMap: Record<number, string> = {
  0: 'å¦',
  1: 'æ˜¯',
};

/** æ˜¯å¦é»˜è®¤颜色 */
function getIsDefaultColor(val: number): string {
  const keys = [0, 1];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** çŠ¶æ€选项 */
const statusOptions = [
  { label: 'ç¦ç”¨', value: 0 },
  { label: 'å¯ç”¨', value: 1 },
];

/** çŠ¶æ€映射 */
const statusMap: Record<number, string> = {
  0: 'ç¦ç”¨',
  1: 'å¯ç”¨',
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
        options: storageOptions,
        placeholder: '请选择å­˜å‚¨ç±»åž‹',
        class: 'w-full',
      },
      fieldName: 'storage',
      label: 'å­˜å‚¨ç±»åž‹',
    },
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: isDefaultOptions,
        placeholder: '请选择æ˜¯å¦é»˜è®¤',
        class: 'w-full',
      },
      fieldName: 'isDefault',
      label: 'æ˜¯å¦é»˜è®¤',
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
const gridOptions: VxeGridProps<ConfigItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'name', title: 'é…ç½®åç§°' },
    { field: 'storage', title: 'å­˜å‚¨ç±»åž‹', width: 120, slots: { default: 'storage_cell' } },
    { field: 'isDefault', title: 'æ˜¯å¦é»˜è®¤', width: 120, slots: { default: 'isDefault_cell' } },
    { field: 'localPath', title: 'æœ¬åœ°å­˜å‚¨è·¯å¾„' },
    { field: 'ossEndpoint', title: 'OSS Endpoint' },
    { field: 'ossBucket', title: 'OSS Bucket' },
    { field: 'ossAccessKey', title: 'OSS AccessKey' },
    { field: 'ossSecretKey', title: 'OSS SecretKey' },
    { field: 'cosRegion', title: 'COS Region' },
    { field: 'cosBucket', title: 'COS Bucket' },
    { field: 'cosSecretID', title: 'COS SecretId' },
    { field: 'cosSecretKey', title: 'COS SecretKey' },
    { field: 'maxSize', title: 'æœ€å¤§æ–‡ä»¶å¤§å°(MB)' },
    { field: 'status', title: 'çŠ¶æ€', width: 120, slots: { default: 'status_cell' } },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getConfigList({
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
function handleEdit(row: ConfigItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: ConfigItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该ä¸Šä¼ é…ç½®吗？',
    okType: 'danger',
    async onOk() {
      await deleteConfig(row.id);
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
      <template #storage_cell="{ row }">
        <Tag :color="getStorageColor(row.storage)">
          {{ storageMap[row.storage] || row.storage }}
        </Tag>
      </template>
      <template #isDefault_cell="{ row }">
        <Tag :color="getIsDefaultColor(row.isDefault)">
          {{ isDefaultMap[row.isDefault] || row.isDefault }}
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
