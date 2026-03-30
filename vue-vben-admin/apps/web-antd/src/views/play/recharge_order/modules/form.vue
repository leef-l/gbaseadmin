<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getRechargeOrderDetail,
  createRechargeOrder,
  updateRechargeOrder,
} from '#/api/play/recharge_order';

/** æ”¯ä»˜æ–¹å¼选项 */
const payTypeOptions = [
  { label: 'å¾®ä¿¡æ”¯ä»˜', value: 1 },
  { label: 'æ”¯ä»˜å®æ”¯ä»˜', value: 2 },
];

/** æ”¯ä»˜çŠ¶æ€选项 */
const payStatusOptions = [
  { label: 'å¾…æ”¯ä»˜', value: 0 },
  { label: 'æ”¯ä»˜æˆåŠŸ', value: 1 },
  { label: 'æ”¯ä»˜å¤±è´¥', value: 2 },
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
      label: 'å……å€¼è®¢å•å·',
      rules: 'required',
      componentProps: { placeholder: '请输入å……å€¼è®¢å•å·', maxlength: 32 },
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
      fieldName: 'rechargePlanID',
      label: 'å……å€¼æ–¹æ¡ˆID',
      rules: 'selectRequired',
      componentProps: { options: rechargePlanIDOptions, placeholder: '请选择å……å€¼æ–¹æ¡ˆID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'amount',
      label: 'å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰',
      rules: 'required',
      componentProps: { placeholder: '请输入å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'giftAmount',
      label: 'èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Select',
      fieldName: 'payType',
      label: 'æ”¯ä»˜æ–¹å¼',
      componentProps: { options: payTypeOptions, placeholder: '请选择æ”¯ä»˜æ–¹å¼', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'tradeNo',
      label: 'ç¬¬ä¸‰æ–¹äº¤æ˜“å·',
      componentProps: { placeholder: '请输入ç¬¬ä¸‰æ–¹äº¤æ˜“å·', maxlength: 64 },
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
      label: 'æ”¯ä»˜æ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择æ”¯ä»˜æ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
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
        await updateRechargeOrder({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createRechargeOrder(values);
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
        modalApi.setState({ title: '编辑å……å€¼è®¢å•è¡¨' });
        try {
          const detail = await getRechargeOrderDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建å……å€¼è®¢å•è¡¨' });
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
