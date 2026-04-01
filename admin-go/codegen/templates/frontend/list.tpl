<script setup lang="ts">
{{- if .HasTooltip}}
import { h } from 'vue';
{{- end}}
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
{{- if .HasTooltip}}
import { Button, message, Modal, Tag, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
{{- else}}
import { Button, message, Modal, Tag } from 'ant-design-vue';
{{- end}}

import { useVbenVxeGrid } from '#/adapter/vxe-table';
{{- if .HasParentID}}
import { get{{.ModelName}}Tree, delete{{.ModelName}} } from '#/api/{{.AppName}}/{{.ModuleName}}';
{{- else}}
import { get{{.ModelName}}List, delete{{.ModelName}} } from '#/api/{{.AppName}}/{{.ModuleName}}';
{{- end}}
import type { {{.ModelName}}Item } from '#/api/{{.AppName}}/{{.ModuleName}}/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];
{{range .Fields}}
{{- if and (not .IsHidden) (.IsEnum)}}
/** {{.Label}}选项 */
const {{.NameLower}}Options = [
{{- range .EnumValues}}
  { label: '{{.Label}}', value: {{.Value}} },
{{- end}}
];

/** {{.Label}}映射 */
const {{.NameLower}}Map: Record<number, string> = {
{{- range .EnumValues}}
  {{.Value}}: '{{.Label}}',
{{- end}}
};

/** {{.Label}}颜色 */
function get{{.NameCamel}}Color(val: number): string {
  const keys = [{{range $i, $v := .EnumValues}}{{if $i}}, {{end}}{{$v.Value}}{{end}}];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}
{{end}}
{{- end}}
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
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (.IsEnum)}}
    {
      component: 'Select',
      componentProps: {
        allowClear: true,
        options: {{.NameLower}}Options,
        placeholder: '请选择{{.Label}}',
        class: 'w-full',
      },
      fieldName: '{{.NameLower}}',
      label: '{{.ShortLabel}}',
    },
{{- end}}
{{- end}}
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<{{.ModelName}}Item> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
{{- $isTree := .HasParentID}}
{{- $firstDataCol := true}}
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (not .IsParentID) (not .IsTimeField) (not .IsMultiFK) (not .IsPassword)}}
{{- if .RefFieldJSON}}
    { field: '{{.RefFieldJSON}}', title: '{{.ShortLabel}}'{{if .TooltipText}}, slots: { header: () => h('span', {}, ['{{.ShortLabel}} ', h(Tooltip, { title: '{{.TooltipText}}' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) }{{end}}{{if and $isTree $firstDataCol}}, treeNode: true{{end}} },
{{- else if .IsEnum}}
    { field: '{{.NameLower}}', title: '{{.ShortLabel}}', width: 120, slots: { default: '{{.NameLower}}_cell' }{{if and $isTree $firstDataCol}}, treeNode: true{{end}} },
{{- else if eq .Component "ImageUpload"}}
    { field: '{{.NameLower}}', title: '{{.ShortLabel}}', width: 100, slots: { default: '{{.NameLower}}_cell' }{{if and $isTree $firstDataCol}}, treeNode: true{{end}} },
{{- else if or (eq .Component "RichText") (eq .Component "JsonEditor")}}
{{- /* 富文本和JSON字段不在列表中显示，不消耗 firstDataCol */}}
{{- else}}
    { field: '{{.NameLower}}', title: '{{.ShortLabel}}'{{if .TooltipText}}, slots: { header: () => h('span', {}, ['{{.ShortLabel}} ', h(Tooltip, { title: '{{.TooltipText}}' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) }{{end}}{{if and $isTree $firstDataCol}}, treeNode: true{{end}} },
{{- end}}
{{- if not (or (eq .Component "RichText") (eq .Component "JsonEditor"))}}
{{- $firstDataCol = false}}
{{- end}}
{{- end}}
{{- end}}
{{- range .Fields}}
{{- if and (not .IsHidden) (.IsTimeField)}}
    { field: '{{.NameLower}}', title: '{{.ShortLabel}}'{{if .TooltipText}}, slots: { header: () => h('span', {}, ['{{.ShortLabel}} ', h(Tooltip, { title: '{{.TooltipText}}' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]) }{{end}}, width: 180, formatter: 'formatDateTime' },
{{- end}}
{{- end}}
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 200, fixed: 'right', slots: { default: 'action' } },
  ],
{{- if .HasParentID}}
  pagerConfig: { enabled: false },
  treeConfig: {
    childrenField: 'children',
    expandAll: false,
  },
  proxyConfig: {
    ajax: {
      query: async (_params, formValues) => {
        return await get{{.ModelName}}Tree(formValues) ?? [];
      },
    },
  },
{{- else}}
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await get{{.ModelName}}List({
          pageNum: page.currentPage,
          pageSize: page.pageSize,
          ...formValues,
        });
        return { items: res?.list ?? [], total: res?.total ?? 0 };
      },
    },
  },
{{- end}}
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
function handleEdit(row: {{.ModelName}}Item) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: {{.ModelName}}Item) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该{{.Comment}}吗？',
    okType: 'danger',
    async onOk() {
      await delete{{.ModelName}}(row.id);
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
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (not .IsParentID) (not .IsTimeField) (not .IsMultiFK) (.IsEnum)}}
      <template #{{.NameLower}}_cell="{ row }">
        <Tag :color="get{{.NameCamel}}Color(row.{{.NameLower}})">
          {{"{{"}} {{.NameLower}}Map[row.{{.NameLower}}] || row.{{.NameLower}} {{"}}"}}
        </Tag>
      </template>
{{- else if eq .Component "ImageUpload"}}
      <template #{{.NameLower}}_cell="{ row }">
        <img v-if="row.{{.NameLower}}" :src="row.{{.NameLower}}" style="width: 48px; height: 48px; object-fit: cover; border-radius: 4px;" />
        <span v-else>-</span>
      </template>
{{- end}}
{{- end}}
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
