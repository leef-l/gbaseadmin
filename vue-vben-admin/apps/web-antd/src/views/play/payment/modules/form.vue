<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getPaymentDetail,
  createPayment,
  updatePayment,
} from '#/api/play/payment';

/** æ”¯ä»˜æ–¹å¼选项 */
const payTypeOptions = [
  { label: 'å¾®ä¿¡æ”¯ä»˜', value: 1 },
  { label: 'æ”¯ä»˜å®æ”¯ä»˜', value: 2 },
  { label: 'ä½™é¢æ”¯ä»˜', value: 3 },
];

/** æ”¯ä»˜çŠ¶æ€选项 */
const payStatusOptions = [
  { label: 'å¾…æ”¯ä»˜', value: 0 },
  { label: 'æ”¯ä»˜æˆåŠŸ', value: 1 },
  { label: 'æ”¯ä»˜å¤±è´¥', value: 2 },
  { label: 'å·²é€€æ¬¾', value: 3 },
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
      component: 'Select',
      fieldName: 'memberID',
      label: 'ä¼šå‘˜ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择ä¼šå‘˜ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'paymentNo',
      label: 'æ”¯ä»˜æµæ°´å·',
      rules: 'required',
      componentProps: { placeholder: '请输入æ”¯ä»˜æµæ°´å·', maxlength: 64 },
    },
    {
      component: 'Input',
      fieldName: 'tradeNo',
      label: 'ç¬¬ä¸‰æ–¹äº¤æ˜“å·',
      componentProps: { placeholder: '请输入ç¬¬ä¸‰æ–¹äº¤æ˜“å·', maxlength: 64 },
    },
    {
      component: 'Select',
      fieldName: 'payType',
      label: 'æ”¯ä»˜æ–¹å¼',
      componentProps: { options: payTypeOptions, placeholder: '请选择æ”¯ä»˜æ–¹å¼', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'payAmount',
      label: 'æ”¯ä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入æ”¯ä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Select',
      fieldName: 'payStatus',
      label: 'æ”¯ä»˜çŠ¶æ€',
      componentProps: { options: payStatusOptions, placeholder: '请选择æ”¯ä»˜çŠ¶æ€', allowClear: true, class: 'w-full' },
    },
    {
      component: 'DatePicker',
      fieldName: 'payAt',
      label: 'æ”¯ä»˜æˆåŠŸæ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择æ”¯ä»˜æˆåŠŸæ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'refundAt',
      label: 'é€€æ¬¾æ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择é€€æ¬¾æ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'Input',
      fieldName: 'refundAmount',
      label: 'é€€æ¬¾é‡‘é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入é€€æ¬¾é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'callbackContent',
      label: 'å›žè°ƒæŠ¥æ–‡',
      componentProps: { placeholder: '请输入å›žè°ƒæŠ¥æ–‡', maxlength: 65535 },
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
        await updatePayment({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createPayment(values);
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
        modalApi.setState({ title: '编辑æ”¯ä»˜è®°å½•è¡¨' });
        try {
          const detail = await getPaymentDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建æ”¯ä»˜è®°å½•è¡¨' });
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
