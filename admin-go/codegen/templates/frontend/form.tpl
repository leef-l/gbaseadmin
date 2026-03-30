<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  get{{.ModelName}}Detail,
  create{{.ModelName}},
  update{{.ModelName}},{{if .HasParentID}}
  get{{.ModelName}}Tree,{{end}}
} from '#/api/{{.AppName}}/{{.ModuleName}}';
{{- if .HasParentID}}
import type { {{.ModelName}}Item } from '#/api/{{.AppName}}/{{.ModuleName}}/types';

const treeData = ref<{{.ModelName}}Item[]>([]);
{{- end}}
{{range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (.IsEnum) (ne .Component "Switch")}}
/** {{.Label}}选项 */
const {{.NameLower}}Options = [
{{- range .EnumValues}}
  { label: '{{.Label}}', value: {{.Value}} },
{{- end}}
];
{{end}}
{{- end}}
const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID)}}
{{- if eq .Component "Password"}}
    {
      component: 'InputPassword',
      fieldName: '{{.NameLower}}',
      label: '{{.Label}}',
{{- if .IsRequired}}
      rules: 'required',
{{- end}}
      componentProps: { placeholder: '请输入{{.Label}}' },
    },
{{- else if eq .Component "InputNumber"}}
    {
      component: 'InputNumber',
      fieldName: '{{.NameLower}}',
      label: '{{.Label}}',
{{- if .IsRequired}}
      rules: 'required',
{{- end}}
      componentProps: { placeholder: '请输入{{.Label}}', class: 'w-full' },
    },
{{- else if eq .Component "Textarea"}}
    {
      component: 'Textarea',
      fieldName: '{{.NameLower}}',
      label: '{{.Label}}',
{{- if .IsRequired}}
      rules: 'required',
{{- end}}
      componentProps: { placeholder: '请输入{{.Label}}', rows: 4{{if gt .MaxLength 0}}, maxlength: {{.MaxLength}}{{end}} },
    },
{{- else if eq .Component "Switch"}}
    {
      component: 'Switch',
      fieldName: '{{.NameLower}}',
      label: '{{.Label}}',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: {{if .DefaultValue}}{{.DefaultValue}}{{else}}0{{end}},
    },
{{- else if eq .Component "Radio"}}
    {
      component: 'RadioGroup',
      fieldName: '{{.NameLower}}',
      label: '{{.Label}}',
{{- if .IsRequired}}
      rules: 'required',
{{- end}}
      componentProps: { options: {{.NameLower}}Options },
    },
{{- else if eq .Component "Select"}}
    {
      component: 'Select',
      fieldName: '{{.NameLower}}',
      label: '{{.Label}}',
{{- if .IsRequired}}
      rules: 'selectRequired',
{{- end}}
      componentProps: { options: {{.NameLower}}Options, placeholder: '请选择{{.Label}}', allowClear: true, class: 'w-full' },
    },
{{- else if eq .Component "SelectMulti"}}
    {
      component: 'Select',
      fieldName: '{{.NameLower}}',
      label: '{{.Label}}',
{{- if .IsRequired}}
      rules: 'selectRequired',
{{- end}}
      componentProps: { options: {{.NameLower}}Options, placeholder: '请选择{{.Label}}', mode: 'multiple', allowClear: true, class: 'w-full' },
    },
{{- else if eq .Component "TreeSelectSingle"}}
    {
      component: 'TreeSelect',
      fieldName: '{{.NameLower}}',
      label: '{{.Label}}',
{{- if .IsRequired}}
      rules: 'selectRequired',
{{- end}}
      componentProps: {
        treeData: treeData.value,
        fieldNames: { label: '{{if .RefDisplayField}}{{.RefDisplayField}}{{else}}title{{end}}', value: 'id', children: 'children' },
        placeholder: '请选择{{.Label}}',
        allowClear: true,
        treeDefaultExpandAll: true,
        class: 'w-full',
      },
    },
{{- else if eq .Component "TreeSelectMulti"}}
    {
      component: 'TreeSelect',
      fieldName: '{{.NameLower}}',
      label: '{{.Label}}',
{{- if .IsRequired}}
      rules: 'selectRequired',
{{- end}}
      componentProps: {
        treeData: treeData.value,
        fieldNames: { label: '{{if .RefDisplayField}}{{.RefDisplayField}}{{else}}title{{end}}', value: 'id', children: 'children' },
        placeholder: '请选择{{.Label}}',
        allowClear: true,
        treeCheckable: true,
        treeDefaultExpandAll: true,
        class: 'w-full',
      },
    },
{{- else if eq .Component "DateTimePicker"}}
    {
      component: 'DatePicker',
      fieldName: '{{.NameLower}}',
      label: '{{.Label}}',
{{- if .IsRequired}}
      rules: 'required',
{{- end}}
      componentProps: { showTime: true, placeholder: '请选择{{.Label}}', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
{{- else}}
    {
      component: 'Input',
      fieldName: '{{.NameLower}}',
      label: '{{.Label}}',
{{- if .IsRequired}}
      rules: 'required',
{{- end}}
      componentProps: { placeholder: '请输入{{.Label}}'{{if gt .MaxLength 0}}, maxlength: {{.MaxLength}}{{end}} },
    },
{{- end}}
{{- end}}
{{- end}}
  ],
});

/** Modal 配置 */
const [Modal, modalApi] = useVbenModal({
  fullscreenButton: false,
  onCancel() {
    modalApi.close();
  },
  onConfirm: async () => {
    const values = await formApi.validateAndSubmitForm();
    if (!values) return;
    modalApi.lock();
    try {
      if (isEdit.value) {
        await update{{.ModelName}}({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await create{{.ModelName}}(values);
        message.success('创建成功');
      }
      emit('success');
      modalApi.close();
    } finally {
      modalApi.lock(false);
    }
  },
  async onOpenChange(isOpen: boolean) {
    if (isOpen) {
      const data = modalApi.getData<{ id?: string } | null>();
{{- if .HasParentID}}
      // 加载树形数据
      try {
        const res = await get{{.ModelName}}Tree();
        treeData.value = [
          { id: '0', title: '顶级节点', children: res ?? [] } as {{.ModelName}}Item,
        ];
        formApi.updateSchema([
          {
            fieldName: '{{range .Fields}}{{if .IsParentID}}{{.NameLower}}{{end}}{{end}}',
            componentProps: { treeData: treeData.value },
          },
        ]);
      } catch {
        // ignore
      }
{{- end}}
      if (data?.id) {
        isEdit.value = true;
        editId.value = data.id;
        modalApi.setState({ title: '编辑{{.Comment}}' });
        try {
          const detail = await get{{.ModelName}}Detail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建{{.Comment}}' });
        formApi.resetForm();
      }
    }
  },
});
</script>

<template>
  <Modal class="w-[600px]">
    <Form />
  </Modal>
</template>
