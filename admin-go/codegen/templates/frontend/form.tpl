<script setup lang="ts">
import { ref, reactive, nextTick } from 'vue';
import { message } from 'ant-design-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  get{{.ModelName}}Detail,
  create{{.ModelName}},
  update{{.ModelName}},
} from '#/api/system/{{.ModuleName}}';
{{- if .HasParentID}}
import { get{{.ModelName}}Tree } from '#/api/system/{{.ModuleName}}';
import type { {{.ModelName}}Item } from '#/api/system/{{.ModuleName}}/types';
{{- end}}

const emit = defineEmits<{ success: [] }>();

const visible = ref(false);
const confirmLoading = ref(false);
const isEdit = ref(false);
const editId = ref('');
const formRef = ref<FormInstance>();
{{- if .HasParentID}}
const treeData = ref<{{.ModelName}}Item[]>([]);
{{- end}}

/** 表单初始值 */
const initialForm = {
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID)}}
{{- if eq .Component "Switch"}}
  {{.NameLower}}: {{if .DefaultValue}}{{.DefaultValue}}{{else}}0{{end}},
{{- else if eq .Component "InputNumber"}}
  {{.NameLower}}: {{if .DefaultValue}}{{.DefaultValue}}{{else}}0{{end}},
{{- else if or (eq .Component "SelectMulti") (eq .Component "TreeSelectMulti") .IsMultiFK}}
  {{.NameLower}}: [] as string[],
{{- else if .IsEnum}}
  {{.NameLower}}: {{if .DefaultValue}}{{.DefaultValue}}{{else}}undefined{{end}} as number | undefined,
{{- else if or .IsForeignKey .IsParentID}}
  {{.NameLower}}: '' as string,
{{- else if eq .TSType "number"}}
  {{.NameLower}}: {{if .DefaultValue}}{{.DefaultValue}}{{else}}0{{end}},
{{- else if eq .TSType "boolean"}}
  {{.NameLower}}: {{if .DefaultValue}}{{.DefaultValue}}{{else}}false{{end}},
{{- else}}
  {{.NameLower}}: '{{.DefaultValue}}',
{{- end}}
{{- end}}
{{- end}}
};

const formData = reactive({ ...initialForm });

/** 校验规则 */
const rules = {
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) .IsRequired}}
  {{.NameLower}}: [
    {
      required: true,
      message: '请{{if or (eq .Component "Select") (eq .Component "Radio") (eq .Component "TreeSelectSingle") (eq .Component "TreeSelectMulti") (eq .Component "SelectMulti") (eq .Component "DateTimePicker")}}选择{{else}}输入{{end}}{{.Label}}',
      trigger: '{{if or (eq .Component "Select") (eq .Component "Radio") (eq .Component "TreeSelectSingle") (eq .Component "TreeSelectMulti") (eq .Component "SelectMulti") (eq .Component "DateTimePicker")}}change{{else}}blur{{end}}',
    },
  ],
{{- end}}
{{- end}}
};
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
/** 打开弹窗 */
async function open(id?: string) {
  visible.value = true;
  isEdit.value = !!id;
  // 重置表单
  Object.assign(formData, { ...initialForm });
  await nextTick();
  formRef.value?.clearValidate();
{{- if .HasParentID}}

  // 加载树形数据
  try {
    const res = await get{{.ModelName}}Tree();
    treeData.value = [
      { id: '0', name: '顶级节点', children: res ?? [] } as any,
    ];
  } catch {
    // ignore
  }
{{- end}}

  if (id) {
    editId.value = id;
    try {
      const detail = await get{{.ModelName}}Detail(id);
      if (detail) {
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID)}}
        formData.{{.NameLower}} = detail.{{.NameLower}} ?? initialForm.{{.NameLower}};
{{- end}}
{{- end}}
      }
    } catch {
      message.error('获取详情失败');
    }
  }
}

/** 提交 */
async function handleOk() {
  try {
    await formRef.value?.validate();
  } catch {
    return;
  }

  confirmLoading.value = true;
  try {
    if (isEdit.value) {
      await update{{.ModelName}}({
        id: editId.value,
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID)}}
        {{.NameLower}}: formData.{{.NameLower}},
{{- end}}
{{- end}}
      });
      message.success('更新成功');
    } else {
      await create{{.ModelName}}({
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID)}}
        {{.NameLower}}: formData.{{.NameLower}},
{{- end}}
{{- end}}
      });
      message.success('创建成功');
    }
    visible.value = false;
    emit('success');
  } finally {
    confirmLoading.value = false;
  }
}

defineExpose({ open });
</script>

<template>
  <a-modal
    v-model:open="visible"
    :title="isEdit ? '编辑{{.Comment}}' : '新建{{.Comment}}'"
    :confirm-loading="confirmLoading"
    :destroy-on-close="false"
    width="600px"
    @ok="handleOk"
  >
    <a-form
      ref="formRef"
      :model="formData"
      :rules="rules"
      :label-col="{ span: 5 }"
      :wrapper-col="{ span: 17 }"
      class="mt-4"
    >
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID)}}
{{- if eq .Component "Input"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-input v-model:value="formData.{{.NameLower}}" placeholder="请输入{{.Label}}"{{if gt .MaxLength 0}} :maxlength="{{.MaxLength}}"{{end}} />
      </a-form-item>
{{- else if eq .Component "InputNumber"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-input-number v-model:value="formData.{{.NameLower}}" placeholder="请输入{{.Label}}" style="width: 100%" />
      </a-form-item>
{{- else if eq .Component "Textarea"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-textarea v-model:value="formData.{{.NameLower}}" placeholder="请输入{{.Label}}" :rows="4"{{if gt .MaxLength 0}} :maxlength="{{.MaxLength}}"{{end}} />
      </a-form-item>
{{- else if eq .Component "Password"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-input-password v-model:value="formData.{{.NameLower}}" placeholder="请输入{{.Label}}" />
      </a-form-item>
{{- else if eq .Component "Switch"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-switch v-model:checked="formData.{{.NameLower}}" :checked-value="1" :un-checked-value="0" />
      </a-form-item>
{{- else if eq .Component "Radio"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-radio-group v-model:value="formData.{{.NameLower}}" :options="{{.NameLower}}Options" />
      </a-form-item>
{{- else if eq .Component "Select"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-select v-model:value="formData.{{.NameLower}}" :options="{{.NameLower}}Options" placeholder="请选择{{.Label}}" allow-clear />
      </a-form-item>
{{- else if eq .Component "SelectMulti"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-select v-model:value="formData.{{.NameLower}}" :options="{{.NameLower}}Options" placeholder="请选择{{.Label}}" mode="multiple" allow-clear />
      </a-form-item>
{{- else if eq .Component "TreeSelectSingle"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-tree-select
          v-model:value="formData.{{.NameLower}}"
          :tree-data="treeData"
          :field-names="{ label: 'name', value: 'id', children: 'children' }"
          placeholder="请选择{{.Label}}"
          allow-clear
          tree-default-expand-all
        />
      </a-form-item>
{{- else if eq .Component "TreeSelectMulti"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-tree-select
          v-model:value="formData.{{.NameLower}}"
          :tree-data="treeData"
          :field-names="{ label: 'name', value: 'id', children: 'children' }"
          placeholder="请选择{{.Label}}"
          allow-clear
          tree-checkable
          tree-default-expand-all
        />
      </a-form-item>
{{- else if eq .Component "DateTimePicker"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-date-picker
          v-model:value="formData.{{.NameLower}}"
          show-time
          placeholder="请选择{{.Label}}"
          style="width: 100%"
          value-format="YYYY-MM-DD HH:mm:ss"
        />
      </a-form-item>
{{- else if eq .Component "InputUrl"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-input v-model:value="formData.{{.NameLower}}" placeholder="请输入{{.Label}}" addon-before="https://" />
      </a-form-item>
{{- else if eq .Component "ImageUpload"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-upload list-type="picture-card" :max-count="1">
          <div>
            <plus-outlined />
            <div class="mt-2">上传图片</div>
          </div>
        </a-upload>
      </a-form-item>
{{- else if eq .Component "FileUpload"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-upload>
          <a-button>上传文件</a-button>
        </a-upload>
      </a-form-item>
{{- else if or (eq .Component "RichText") (eq .Component "JsonEditor")}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-textarea v-model:value="formData.{{.NameLower}}" placeholder="请输入{{.Label}}" :rows="6" />
      </a-form-item>
{{- else if eq .Component "IconPicker"}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-input v-model:value="formData.{{.NameLower}}" placeholder="请输入图标名称" />
      </a-form-item>
{{- else}}
      <a-form-item label="{{.Label}}" name="{{.NameLower}}">
        <a-input v-model:value="formData.{{.NameLower}}" placeholder="请输入{{.Label}}" />
      </a-form-item>
{{- end}}
{{- end}}
{{- end}}
    </a-form>
  </a-modal>
</template>
