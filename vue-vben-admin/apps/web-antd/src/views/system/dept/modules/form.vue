<script setup lang="ts">
import { ref, reactive, nextTick } from 'vue';
import { message } from 'ant-design-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  getDeptDetail,
  createDept,
  updateDept,
} from '#/api/system/dept';
import { getDeptTree } from '#/api/system/dept';
import type { DeptItem } from '#/api/system/dept/types';

const emit = defineEmits<{ success: [] }>();

const visible = ref(false);
const confirmLoading = ref(false);
const isEdit = ref(false);
const editId = ref('');
const formRef = ref<FormInstance>();
const treeData = ref<DeptItem[]>([]);

/** 表单初始值 */
const initialForm = {
  parentID: '' as string,
  title: '',
  username: '',
  email: '',
  sort: 0,
  status: 1,
};

const formData = reactive({ ...initialForm });

/** 校验规则 */
const rules = {
  title: [
    {
      required: true,
      message: '请输入部门名称',
      trigger: 'blur',
    },
  ],
};

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
    const res = await getDeptTree();
    treeData.value = [
      { id: '0', name: '顶级节点', children: res ?? [] } as any,
    ];
  } catch {
    // ignore
  }

  if (id) {
    editId.value = id;
    try {
      const detail = await getDeptDetail(id);
      if (detail) {
        formData.parentID = detail.parentID ?? initialForm.parentID;
        formData.title = detail.title ?? initialForm.title;
        formData.username = detail.username ?? initialForm.username;
        formData.email = detail.email ?? initialForm.email;
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
      await updateDept({
        id: editId.value,
        parentID: formData.parentID,
        title: formData.title,
        username: formData.username,
        email: formData.email,
        sort: formData.sort,
        status: formData.status,
      });
      message.success('更新成功');
    } else {
      await createDept({
        parentID: formData.parentID,
        title: formData.title,
        username: formData.username,
        email: formData.email,
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
    :title="isEdit ? '编辑部门表' : '新建部门表'"
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
      <a-form-item label="上级部门ID，0 表示顶级部门" name="parentID">
        <a-tree-select
          v-model:value="formData.parentID"
          :tree-data="treeData"
          :field-names="{ label: 'name', value: 'id', children: 'children' }"
          placeholder="请选择上级部门ID，0 表示顶级部门"
          allow-clear
          tree-default-expand-all
        />
      </a-form-item>
      <a-form-item label="部门名称" name="title">
        <a-input v-model:value="formData.title" placeholder="请输入部门名称" :maxlength="50" />
      </a-form-item>
      <a-form-item label="部门负责人姓名" name="username">
        <a-input v-model:value="formData.username" placeholder="请输入部门负责人姓名" :maxlength="50" />
      </a-form-item>
      <a-form-item label="负责人邮箱" name="email">
        <a-input v-model:value="formData.email" placeholder="请输入负责人邮箱" :maxlength="100" />
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
