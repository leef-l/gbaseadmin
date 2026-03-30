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

/** ç¬¬ä¸‰æ–¹å¹³å°选项 */
const providerOptions = [
  { label: 'å¾®ä¿¡', value: 1 },
  { label: 'æ”¯ä»˜å®', value: 2 },
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
      label: 'ä¼šå‘˜ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择ä¼šå‘˜ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'provider',
      label: 'ç¬¬ä¸‰æ–¹å¹³å°',
      rules: 'selectRequired',
      componentProps: { options: providerOptions, placeholder: '请选择ç¬¬ä¸‰æ–¹å¹³å°', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'openID',
      label: 'ç¬¬ä¸‰æ–¹OpenID',
      rules: 'selectRequired',
      componentProps: { options: openIDOptions, placeholder: '请选择ç¬¬ä¸‰æ–¹OpenID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'unionID',
      label: 'ç¬¬ä¸‰æ–¹UnionID',
      componentProps: { options: unionIDOptions, placeholder: '请选择ç¬¬ä¸‰æ–¹UnionID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'nickname',
      label: 'ç¬¬ä¸‰æ–¹æ˜µç§°',
      componentProps: { placeholder: '请输入ç¬¬ä¸‰æ–¹æ˜µç§°', maxlength: 50 },
    },
    {
      component: 'Input',
      fieldName: 'avatar',
      label: 'ç¬¬ä¸‰æ–¹å¤´åƒ',
      componentProps: { placeholder: '请输入ç¬¬ä¸‰æ–¹å¤´åƒ', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'accessToken',
      label: 'è®¿é—®ä»¤ç‰Œ',
      componentProps: { placeholder: '请输入è®¿é—®ä»¤ç‰Œ', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'refreshToken',
      label: 'åˆ·æ–°ä»¤ç‰Œ',
      componentProps: { placeholder: '请输入åˆ·æ–°ä»¤ç‰Œ', maxlength: 500 },
    },
    {
      component: 'DatePicker',
      fieldName: 'expireAt',
      label: 'ä»¤ç‰Œè¿‡æœŸæ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择ä»¤ç‰Œè¿‡æœŸæ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
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
        modalApi.setState({ title: '编辑ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨' });
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
        modalApi.setState({ title: '新建ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨' });
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
