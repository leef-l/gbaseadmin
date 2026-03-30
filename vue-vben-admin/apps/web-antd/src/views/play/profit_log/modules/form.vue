<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getProfitLogDetail,
  createProfitLog,
  updateProfitLog,
} from '#/api/play/profit_log';

/** ç»“ç®—çŠ¶æ€选项 */
const settleStatusOptions = [
  { label: 'å¾…ç»“ç®—', value: 0 },
  { label: 'å·²ç»“ç®—', value: 1 },
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
      fieldName: 'orderID',
      label: 'è®¢å•ID',
      rules: 'selectRequired',
      componentProps: { options: orderIDOptions, placeholder: '请选择è®¢å•ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'orderNo',
      label: 'è®¢å•ç¼–å·',
      rules: 'required',
      componentProps: { placeholder: '请输入è®¢å•ç¼–å·', maxlength: 32 },
    },
    {
      component: 'Input',
      fieldName: 'payAmount',
      label: 'å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Select',
      fieldName: 'coachID',
      label: 'é™ªçŽ©å¸ˆID',
      rules: 'selectRequired',
      componentProps: { options: coachIDOptions, placeholder: '请选择é™ªçŽ©å¸ˆID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'shopID',
      label: 'åº—é“ºID',
      componentProps: { options: shopIDOptions, placeholder: '请选择åº—é“ºID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'platformRate',
      label: 'å¹³å°æŠ½æˆæ¯”ä¾‹',
      componentProps: { placeholder: '请输入å¹³å°æŠ½æˆæ¯”ä¾‹' },
    },
    {
      component: 'Input',
      fieldName: 'platformAmount',
      label: 'å¹³å°æŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入å¹³å°æŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'shopRate',
      label: 'åº—é“ºæŠ½æˆæ¯”ä¾‹',
      componentProps: { placeholder: '请输入åº—é“ºæŠ½æˆæ¯”ä¾‹' },
    },
    {
      component: 'Input',
      fieldName: 'shopAmount',
      label: 'åº—é“ºæŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入åº—é“ºæŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'coachAmount',
      label: 'é™ªçŽ©å¸ˆæ”¶å…¥ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入é™ªçŽ©å¸ˆæ”¶å…¥ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Select',
      fieldName: 'settleStatus',
      label: 'ç»“ç®—çŠ¶æ€',
      componentProps: { options: settleStatusOptions, placeholder: '请选择ç»“ç®—çŠ¶æ€', allowClear: true, class: 'w-full' },
    },
    {
      component: 'DatePicker',
      fieldName: 'settleAt',
      label: 'ç»“ç®—æ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择ç»“ç®—æ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
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
        await updateProfitLog({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createProfitLog(values);
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
        modalApi.setState({ title: '编辑åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨' });
        try {
          const detail = await getProfitLogDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨' });
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
