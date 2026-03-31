<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getRechargeOrderDetail,
  createRechargeOrder,
  updateRechargeOrder,
} from '#/api/play/recharge_order';

/** 支付方式选项 */
const payTypeOptions = [
  { label: '微信支付', value: 1 },
  { label: '支付宝支付', value: 2 },
];

/** 支付状态选项 */
const payStatusOptions = [
  { label: '待支付', value: 0 },
  { label: '支付成功', value: 1 },
  { label: '支付失败', value: 2 },
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
      label: '充值订单号',
      rules: 'required',
      componentProps: { placeholder: '请输入充值订单号', maxlength: 32 },
    },
    {
      component: 'Select',
      fieldName: 'memberID',
      label: '会员ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'rechargePlanID',
      label: '充值方案ID',
      rules: 'selectRequired',
      componentProps: { options: rechargePlanIDOptions, placeholder: '请选择充值方案ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'amount',
      label: () => h('span', {}, ['充值金额 ', h(Tooltip, { title: '分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      rules: 'required',
      componentProps: { placeholder: '请输入充值金额（分）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'giftAmount',
      label: () => h('span', {}, ['赠送金额 ', h(Tooltip, { title: '分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入赠送金额（分）', class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'payType',
      label: '支付方式',
      componentProps: { options: payTypeOptions, placeholder: '请选择支付方式', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'tradeNo',
      label: '第三方交易号',
      componentProps: { placeholder: '请输入第三方交易号', maxlength: 64 },
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
      label: '支付时间',
      componentProps: { showTime: true, placeholder: '请选择支付时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
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
        modalApi.setState({ title: '编辑充值订单表' });
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
        modalApi.setState({ title: '新建充值订单表' });
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
