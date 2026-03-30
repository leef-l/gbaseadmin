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

/** 支付方式选项 */
const payTypeOptions = [
  { label: '微信支付', value: 1 },
  { label: '支付宝支付', value: 2 },
  { label: '余额支付', value: 3 },
];

/** 支付状态选项 */
const payStatusOptions = [
  { label: '待支付', value: 0 },
  { label: '支付成功', value: 1 },
  { label: '支付失败', value: 2 },
  { label: '已退款', value: 3 },
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
      label: '订单ID',
      rules: 'selectRequired',
      componentProps: { options: orderIDOptions, placeholder: '请选择订单ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'memberID',
      label: '会员ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'paymentNo',
      label: '支付流水号（平台内部）',
      rules: 'required',
      componentProps: { placeholder: '请输入支付流水号（平台内部）', maxlength: 64 },
    },
    {
      component: 'Input',
      fieldName: 'tradeNo',
      label: '第三方交易号',
      componentProps: { placeholder: '请输入第三方交易号', maxlength: 64 },
    },
    {
      component: 'Select',
      fieldName: 'payType',
      label: '支付方式',
      componentProps: { options: payTypeOptions, placeholder: '请选择支付方式', allowClear: true, class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'payAmount',
      label: '支付金额（分）',
      componentProps: { placeholder: '请输入支付金额（分）', class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'payStatus',
      label: '支付状态',
      componentProps: { options: payStatusOptions, placeholder: '请选择支付状态', allowClear: true, class: 'w-full' },
    },
    {
      component: 'DatePicker',
      fieldName: 'payAt',
      label: '支付成功时间',
      componentProps: { showTime: true, placeholder: '请选择支付成功时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'refundAt',
      label: '退款时间',
      componentProps: { showTime: true, placeholder: '请选择退款时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'InputNumber',
      fieldName: 'refundAmount',
      label: '退款金额（分）',
      componentProps: { placeholder: '请输入退款金额（分）', class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'callbackContent',
      label: '回调报文',
      componentProps: { placeholder: '请输入回调报文', maxlength: 65535 },
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
        modalApi.setState({ title: '编辑支付记录表' });
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
        modalApi.setState({ title: '新建支付记录表' });
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
