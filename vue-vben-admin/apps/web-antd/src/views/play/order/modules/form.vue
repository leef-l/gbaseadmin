<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getOrderDetail,
  createOrder,
  updateOrder,
} from '#/api/play/order';

/** æ”¯ä»˜æ–¹å¼选项 */
const payTypeOptions = [
  { label: 'æœªæ”¯ä»˜', value: 0 },
  { label: 'å¾®ä¿¡æ”¯ä»˜', value: 1 },
  { label: 'æ”¯ä»˜å®æ”¯ä»˜', value: 2 },
  { label: 'ä½™é¢æ”¯ä»˜', value: 3 },
];

/** è®¢å•çŠ¶æ€选项 */
const orderStatusOptions = [
  { label: 'å¾…æ”¯ä»˜', value: 0 },
  { label: 'å·²æ”¯ä»˜', value: 1 },
  { label: 'è¿›è¡Œä¸­', value: 2 },
  { label: 'å·²å®Œæˆ', value: 3 },
  { label: 'å·²å–æ¶ˆ', value: 4 },
  { label: 'é€€æ¬¾ä¸­', value: 5 },
  { label: 'å·²é€€æ¬¾', value: 6 },
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
      fieldName: 'orderNo',
      label: 'è®¢å•ç¼–å·',
      rules: 'required',
      componentProps: { placeholder: '请输入è®¢å•ç¼–å·', maxlength: 32 },
    },
    {
      component: 'Select',
      fieldName: 'memberID',
      label: 'ä¸‹å•ä¼šå‘˜ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择ä¸‹å•ä¼šå‘˜ID', allowClear: true, class: 'w-full' },
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
      component: 'Select',
      fieldName: 'goodsID',
      label: 'å•†å“ID',
      rules: 'selectRequired',
      componentProps: { options: goodsIDOptions, placeholder: '请选择å•†å“ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'goodsTitle',
      label: 'å•†å“åç§°ï¼ˆå†—ä½™ï¼‰',
      rules: 'required',
      componentProps: { placeholder: '请输入å•†å“åç§°ï¼ˆå†—ä½™ï¼‰', maxlength: 100 },
    },
    {
      component: 'Input',
      fieldName: 'goodsPrice',
      label: 'å•†å“å•ä»·ï¼ˆåˆ†ï¼‰',
      rules: 'required',
      componentProps: { placeholder: '请输入å•†å“å•ä»·ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'quantity',
      label: 'æ•°é‡',
      componentProps: { placeholder: '请输入æ•°é‡' },
    },
    {
      component: 'Input',
      fieldName: 'totalAmount',
      label: 'è®¢å•æ€»é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入è®¢å•æ€»é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'discountAmount',
      label: 'ä¼šå‘˜æŠ˜æ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入ä¼šå‘˜æŠ˜æ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'couponAmount',
      label: 'ä¼˜æƒ åˆ¸æŠµæ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入ä¼˜æƒ åˆ¸æŠµæ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'payAmount',
      label: 'å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Select',
      fieldName: 'couponMemberID',
      label: 'ä½¿ç”¨çš„ä¼˜æƒ åˆ¸é¢†å–è®°å½•ID',
      componentProps: { options: couponMemberIDOptions, placeholder: '请选择ä½¿ç”¨çš„ä¼˜æƒ åˆ¸é¢†å–è®°å½•ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'payType',
      label: 'æ”¯ä»˜æ–¹å¼',
      componentProps: { options: payTypeOptions, placeholder: '请选择æ”¯ä»˜æ–¹å¼', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'orderStatus',
      label: 'è®¢å•çŠ¶æ€',
      componentProps: { options: orderStatusOptions, placeholder: '请选择è®¢å•çŠ¶æ€', allowClear: true, class: 'w-full' },
    },
    {
      component: 'DatePicker',
      fieldName: 'payAt',
      label: 'æ”¯ä»˜æ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择æ”¯ä»˜æ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'startAt',
      label: 'æœåŠ¡å¼€å§‹æ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择æœåŠ¡å¼€å§‹æ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'finishAt',
      label: 'æœåŠ¡å®Œæˆæ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择æœåŠ¡å®Œæˆæ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'cancelAt',
      label: 'å–æ¶ˆæ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择å–æ¶ˆæ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'Input',
      fieldName: 'cancelReason',
      label: 'å–æ¶ˆåŽŸå›',
      componentProps: { placeholder: '请输入å–æ¶ˆåŽŸå›', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'remark',
      label: 'è®¢å•å¤‡æ³¨',
      componentProps: { placeholder: '请输入è®¢å•å¤‡æ³¨', maxlength: 500 },
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
        await updateOrder({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createOrder(values);
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
        modalApi.setState({ title: '编辑è®¢å•è¡¨' });
        try {
          const detail = await getOrderDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建è®¢å•è¡¨' });
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
