<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getActivityJoinDetail,
  createActivityJoin,
  updateActivityJoin,
} from '#/api/play/activity_join';

/** å‚ä¸ŽçŠ¶æ€选项 */
const joinStatusOptions = [
  { label: 'å·²æŠ¥å', value: 0 },
  { label: 'è¿›è¡Œä¸­', value: 1 },
  { label: 'å·²å®Œæˆ', value: 2 },
  { label: 'å·²é¢†å¥–', value: 3 },
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
      fieldName: 'activityID',
      label: 'æ´»åŠ¨ID',
      rules: 'selectRequired',
      componentProps: { options: activityIDOptions, placeholder: '请选择æ´»åŠ¨ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'memberID',
      label: 'ä¼šå‘˜ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择ä¼šå‘˜ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'joinStatus',
      label: 'å‚ä¸ŽçŠ¶æ€',
      componentProps: { options: joinStatusOptions, placeholder: '请选择å‚ä¸ŽçŠ¶æ€', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'currentStep',
      label: 'å½“å‰å®Œæˆåˆ°ç¬¬å‡ æ­¥',
      componentProps: { placeholder: '请输入å½“å‰å®Œæˆåˆ°ç¬¬å‡ æ­¥' },
    },
    {
      component: 'DatePicker',
      fieldName: 'finishAt',
      label: 'å®Œæˆæ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择å®Œæˆæ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'rewardAt',
      label: 'é¢†å¥–æ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择é¢†å¥–æ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'Input',
      fieldName: 'remark',
      label: 'å¤‡æ³¨',
      componentProps: { placeholder: '请输入å¤‡æ³¨', maxlength: 500 },
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
        await updateActivityJoin({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createActivityJoin(values);
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
        modalApi.setState({ title: '编辑æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨' });
        try {
          const detail = await getActivityJoinDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨' });
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
