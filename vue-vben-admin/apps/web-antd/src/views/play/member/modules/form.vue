<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getMemberDetail,
  createMember,
  updateMember,
} from '#/api/play/member';

/** 性别选项 */
const genderOptions = [
  { label: '未知', value: 0 },
  { label: '男', value: 1 },
  { label: '女', value: 2 },
];

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'phone',
      label: '手机号',
      rules: 'required',
      componentProps: { placeholder: '请输入手机号', maxlength: 20 },
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
      label: '昵称',
      componentProps: { placeholder: '请输入昵称', maxlength: 50 },
    },
    {
      component: 'Input',
      fieldName: 'avatar',
      label: '头像',
      componentProps: { placeholder: '请输入头像', maxlength: 500 },
    },
    {
      component: 'Select',
      fieldName: 'gender',
      label: '性别',
      componentProps: { options: genderOptions, placeholder: '请选择性别', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'memberLevelID',
      label: '会员等级ID',
      componentProps: { options: memberLevelIDOptions, placeholder: '请选择会员等级ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'exp',
      label: '经验值',
      componentProps: { placeholder: '请输入经验值' },
    },
    {
      component: 'InputNumber',
      fieldName: 'balance',
      label: '账户余额（分）',
      componentProps: { placeholder: '请输入账户余额（分）', class: 'w-full' },
    },
    {
      component: 'Switch',
      fieldName: 'isCoach',
      label: '是否陪玩师',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 0,
    },
    {
      component: 'Switch',
      fieldName: 'status',
      label: '状态',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 1,
    },
    {
      component: 'DatePicker',
      fieldName: 'lastLoginAt',
      label: '最后登录时间',
      componentProps: { showTime: true, placeholder: '请选择最后登录时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
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
    const values = await formApi.validateAndSubmitForm();
    if (!values) return;
    modalApi.lock();
    try {
      if (isEdit.value) {
        await updateMember({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createMember(values);
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
        modalApi.setState({ title: '编辑会员表' });
        try {
          const detail = await getMemberDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建会员表' });
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
