<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getUsersDetail,
  createUsers,
  updateUsers,
} from '#/api/system/users';

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'username',
      label: '登录用户名',
      rules: 'required',
      componentProps: { placeholder: '请输入登录用户名', maxlength: 50 },
    },
    {
      component: 'InputPassword',
      fieldName: 'password',
      label: '密码（bcrypt 加密）',
      rules: 'required',
      componentProps: { placeholder: '请输入密码（bcrypt 加密）' },
    },
    {
      component: 'Input',
      fieldName: 'nickname',
      label: '昵称/显示名',
      componentProps: { placeholder: '请输入昵称/显示名', maxlength: 50 },
    },
    {
      component: 'Input',
      fieldName: 'email',
      label: '邮箱地址',
      componentProps: { placeholder: '请输入邮箱地址', maxlength: 100 },
    },
    {
      component: 'Input',
      fieldName: 'avatar',
      label: '头像图片 URL',
      componentProps: { placeholder: '请输入头像图片 URL', maxlength: 500 },
    },
    {
      component: 'Switch',
      fieldName: 'status',
      label: '状态',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 1,
    },
  ],
});

/** Modal 配置 */
const [Modal, modalApi] = useVbenModal({
  fullscreenButton: false,
  onCancel() {
    modalApi.close();
  },
  onConfirm: async () => {
    const { valid, values } = await formApi.validateAndSubmitForm();
    if (!valid) return;
    modalApi.lock();
    try {
      if (isEdit.value) {
        await updateUsers({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createUsers(values);
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
      if (data?.id) {
        isEdit.value = true;
        editId.value = data.id;
        modalApi.setState({ title: '编辑用户表' });
        try {
          const detail = await getUsersDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建用户表' });
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
