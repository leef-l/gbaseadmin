<script setup lang="ts">
import { ref, reactive, nextTick } from 'vue';
import { message } from 'ant-design-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  getUsersDetail,
  createUsers,
  updateUsers,
} from '#/api/system/users';

const emit = defineEmits<{ success: [] }>();

const visible = ref(false);
const confirmLoading = ref(false);
const isEdit = ref(false);
const editId = ref('');
const formRef = ref<FormInstance>();

/** 表单初始值 */
const initialForm = {
  username: '',
  password: '',
  nickname: '',
  email: '',
  avatar: '',
  status: 1,
};

const formData = reactive({ ...initialForm });

/** 校验规则 */
const rules = {
  username: [
    {
      required: true,
      message: '请输入登录用户名',
      trigger: 'blur',
    },
  ],
  password: [
    {
      required: true,
      message: '请输入密码（bcrypt 加密）',
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

  if (id) {
    editId.value = id;
    try {
      const detail = await getUsersDetail(id);
      if (detail) {
        formData.username = detail.username ?? initialForm.username;
        formData.password = detail.password ?? initialForm.password;
        formData.nickname = detail.nickname ?? initialForm.nickname;
        formData.email = detail.email ?? initialForm.email;
        formData.avatar = detail.avatar ?? initialForm.avatar;
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
      await updateUsers({
        id: editId.value,
        username: formData.username,
        password: formData.password,
        nickname: formData.nickname,
        email: formData.email,
        avatar: formData.avatar,
        status: formData.status,
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
      <a-form-item label="密码（bcrypt 加密）" name="password">
        <a-input-password v-model:value="formData.password" placeholder="请输入密码（bcrypt 加密）" />
      </a-form-item>
      <a-form-item label="昵称/显示名" name="nickname">
        <a-input v-model:value="formData.nickname" placeholder="请输入昵称/显示名" :maxlength="50" />
      </a-form-item>
      <a-form-item label="邮箱地址" name="email">
        <a-input v-model:value="formData.email" placeholder="请输入邮箱地址" :maxlength="100" />
      </a-form-item>
      <a-form-item label="头像图片 URL" name="avatar">
        <a-upload list-type="picture-card" :max-count="1">
          <div>
            <plus-outlined />
            <div class="mt-2">上传图片</div>
          </div>
        </a-upload>
      </a-form-item>
      <a-form-item label="状态" name="status">
        <a-switch v-model:checked="formData.status" :checked-value="1" :un-checked-value="0" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>
