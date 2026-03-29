<script setup lang="ts">
import { ref, reactive, nextTick } from 'vue';
import { message } from 'ant-design-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  getRoleDetail,
  createRole,
  updateRole,
} from '#/api/system/role';
import { getRoleTree } from '#/api/system/role';
import type { RoleItem } from '#/api/system/role/types';

const emit = defineEmits<{ success: [] }>();

const visible = ref(false);
const confirmLoading = ref(false);
const isEdit = ref(false);
const editId = ref('');
const formRef = ref<FormInstance>();
const treeData = ref<RoleItem[]>([]);

/** 表单初始值 */
const initialForm = {
  parentID: '' as string,
  title: '',
  dataScope: 1 as number | undefined,
  sort: 0,
  status: 1,
};

const formData = reactive({ ...initialForm });

/** 校验规则 */
const rules = {
  title: [
    {
      required: true,
      message: '请输入角色名称',
      trigger: 'blur',
    },
  ],
};

/** 数据范围选项 */
const dataScopeOptions = [
  { label: '全部', value: 1 },
  { label: '本部门及以下', value: 2 },
  { label: '本部门', value: 3 },
  { label: '仅本人', value: 4 },
  { label: '自定义', value: 5 },
];

/** 打开弹窗 */
async function open(id?: string) {
  visible.value = true;
  isEdit.value = !!id;
  // 重置表单
  Object.assign(formData, { ...initialForm });
  await nextTick();
  formRef.value?.clearValidate();

  // 加载树形数据
  try {
    const res = await getRoleTree();
    treeData.value = [
      { id: '0', name: '顶级节点', children: res ?? [] } as any,
    ];
  } catch {
    // ignore
  }

  if (id) {
    editId.value = id;
    try {
      const detail = await getRoleDetail(id);
      if (detail) {
        formData.parentID = detail.parentID ?? initialForm.parentID;
        formData.title = detail.title ?? initialForm.title;
        formData.dataScope = detail.dataScope ?? initialForm.dataScope;
        formData.sort = detail.sort ?? initialForm.sort;
        formData.status = detail.status ?? initialForm.status;
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
      await updateRole({
        id: editId.value,
        parentID: formData.parentID,
        title: formData.title,
        dataScope: formData.dataScope,
        sort: formData.sort,
        status: formData.status,
      });
      message.success('更新成功');
    } else {
      await createRole({
        parentID: formData.parentID,
        title: formData.title,
        dataScope: formData.dataScope,
        sort: formData.sort,
        status: formData.status,
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
    :title="isEdit ? '编辑角色表' : '新建角色表'"
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
      <a-form-item label="上级角色ID，0 表示顶级角色" name="parentID">
        <a-tree-select
          v-model:value="formData.parentID"
          :tree-data="treeData"
          :field-names="{ label: 'name', value: 'id', children: 'children' }"
          placeholder="请选择上级角色ID，0 表示顶级角色"
          allow-clear
          tree-default-expand-all
        />
      </a-form-item>
      <a-form-item label="角色名称" name="title">
        <a-input v-model:value="formData.title" placeholder="请输入角色名称" :maxlength="50" />
      </a-form-item>
      <a-form-item label="数据范围" name="dataScope">
        <a-select v-model:value="formData.dataScope" :options="dataScopeOptions" placeholder="请选择数据范围" allow-clear />
      </a-form-item>
      <a-form-item label="排序（升序）" name="sort">
        <a-input-number v-model:value="formData.sort" placeholder="请输入排序（升序）" style="width: 100%" />
      </a-form-item>
      <a-form-item label="状态" name="status">
        <a-switch v-model:checked="formData.status" :checked-value="1" :un-checked-value="0" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>
