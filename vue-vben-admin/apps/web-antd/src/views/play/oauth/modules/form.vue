<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getOauthDetail,
  createOauth,
  updateOauth,
} from '#/api/play/oauth';

/** 第三方平台选项 */
const providerOptions = [
  { label: '微信', value: 1 },
  { label: '支付宝', value: 2 },
];

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Select',
      fieldName: 'memberID',
      label: '会员ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'provider',
      label: '第三方平台',
      rules: 'selectRequired',
      componentProps: { options: providerOptions, placeholder: '请选择第三方平台', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'openID',
      label: '第三方OpenID',
      rules: 'selectRequired',
      componentProps: { options: openIDOptions, placeholder: '请选择第三方OpenID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'unionID',
      label: '第三方UnionID',
      componentProps: { options: unionIDOptions, placeholder: '请选择第三方UnionID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'nickname',
      label: '第三方昵称',
      componentProps: { placeholder: '请输入第三方昵称', maxlength: 50 },
    },
    {
      component: 'Input',
      fieldName: 'avatar',
      label: '第三方头像',
      componentProps: { placeholder: '请输入第三方头像', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'accessToken',
      label: '访问令牌',
      componentProps: { placeholder: '请输入访问令牌', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'refreshToken',
      label: '刷新令牌',
      componentProps: { placeholder: '请输入刷新令牌', maxlength: 500 },
    },
    {
      component: 'DatePicker',
      fieldName: 'expireAt',
      label: '令牌过期时间',
      componentProps: { showTime: true, placeholder: '请选择令牌过期时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
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
        await updateOauth({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createOauth(values);
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
        modalApi.setState({ title: '编辑第三方登录绑定表' });
        try {
          const detail = await getOauthDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建第三方登录绑定表' });
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
