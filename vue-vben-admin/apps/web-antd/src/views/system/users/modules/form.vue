<script setup lang="ts">
import { ref, reactive, nextTick } from 'vue';
import { message } from 'ant-design-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  getUsersDetail,
  createUsers,
  updateUsers,
} from '#/api/system/users';
import { getDeptTree } from '#/api/system/dept';
import { getRoleTree } from '#/api/system/role';
import type { DeptItem } from '#/api/system/dept/types';
import type { RoleItem } from '#/api/system/role/types';

const emit = defineEmits<{ success: [] }>();

const visible = ref(false);
const confirmLoading = ref(false);
const isEdit = ref(false);
const editId = ref('');
const formRef = ref<FormInstance>();
const deptTreeData = ref<DeptItem[]>([]);
const roleList = ref<RoleItem[]>([]);

/** 表单初始值 */
const initialForm = {
  username: '',
  password: '',
  nickname: '',
  email: '',
  avatar: '',
  status: 1,
  deptId: undefined as string | undefined,
  roleIds: [] as string[],
};

const formData = reactive({ ...initialForm });

/** 校验规则 */
const rules = {
  username: [{ required: true, message: '请输入登录用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
};

/** 递归展平角色树为列表 */
function flattenRoles(nodes: RoleItem[], result: { id: string; title: string }[] = [], prefix = '') {
  for (const node of nodes) {
    result.push({ id: node.id, title: prefix + node.title });
    if (node.children?.length) {
      flattenRoles(node.children, result, prefix + '  ');
    }
  }
  return result;
}

/** 打开弹窗 */
async function open(id?: string) {
  visible.value = true;
  isEdit.value = !!id;
  Object.assign(formData, { ...initialForm, roleIds: [] });
  await nextTick();
  formRef.value?.clearValidate();

  // 加载部门树
  try {
    const res = await getDeptTree();
    deptTreeData.value = [
      { id: '0', title: '顶级部门', children: res ?? [] } as any,
    ];
  } catch { /* ignore */ }

  // 加载角色列表
  try {
    const res = await getRoleTree();
    roleList.value = res ?? [];
  } catch { /* ignore */ }

  if (id) {
    editId.value = id;
    try {
      const detail = await getUsersDetail(id);
      if (detail) {
        formData.username = detail.username ?? initialForm.username;
        formData.password = '';
        formData.nickname = detail.nickname ?? initialForm.nickname;
        formData.email = detail.email ?? initialForm.email;
        formData.avatar = detail.avatar ?? initialForm.avatar;
        formData.status = detail.status ?? initialForm.status;
        formData.deptId = detail.deptId ?? initialForm.deptId;
        formData.roleIds = detail.roleIds ?? [];
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
      await updateUsers({
        id: editId.value,
        username: formData.username,
        password: formData.password,
        nickname: formData.nickname,
        email: formData.email,
        avatar: formData.avatar,
        status: formData.status,
        deptId: formData.deptId,
        roleIds: formData.roleIds,
      });
      message.success('更新成功');
    } else {
      await createUsers({
        username: formData.username,
        password: formData.password,
        nickname: formData.nickname,
        email: formData.email,
        avatar: formData.avatar,
        status: formData.status,
        deptId: formData.deptId,
        roleIds: formData.roleIds,
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
    :title="isEdit ? '编辑用户表' : '新建用户表'"
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
      <a-form-item label="登录用户名" name="username">
        <a-input v-model:value="formData.username" placeholder="请输入登录用户名" :maxlength="50" />
      </a-form-item>
      <a-form-item label="密码" name="password">
        <a-input-password v-model:value="formData.password" :placeholder="isEdit ? '留空则不修改密码' : '请输入密码'" />
      </a-form-item>
      <a-form-item label="昵称" name="nickname">
        <a-input v-model:value="formData.nickname" placeholder="请输入昵称" :maxlength="50" />
      </a-form-item>
      <a-form-item label="邮箱" name="email">
        <a-input v-model:value="formData.email" placeholder="请输入邮箱" :maxlength="100" />
      </a-form-item>
      <a-form-item label="所属部门" name="deptId">
        <a-tree-select
          v-model:value="formData.deptId"
          :tree-data="deptTreeData"
          :field-names="{ label: 'title', value: 'id', children: 'children' }"
          placeholder="请选择所属部门"
          allow-clear
          tree-default-expand-all
        />
      </a-form-item>
      <a-form-item label="角色" name="roleIds">
        <a-select
          v-model:value="formData.roleIds"
          mode="multiple"
          placeholder="请选择角色"
          allow-clear
        >
          <a-select-option
            v-for="role in flattenRoles(roleList)"
            :key="role.id"
            :value="role.id"
          >
            {{ role.title }}
          </a-select-option>
        </a-select>
      </a-form-item>
      <a-form-item label="状态" name="status">
        <a-switch v-model:checked="formData.status" :checked-value="1" :un-checked-value="0" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>
