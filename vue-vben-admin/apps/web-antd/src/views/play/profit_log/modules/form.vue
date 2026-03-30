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

/** 结算状态选项 */
const settleStatusOptions = [
  { label: '待结算', value: 0 },
  { label: '已结算', value: 1 },
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
      component: 'Input',
      fieldName: 'orderNo',
      label: '订单编号',
      rules: 'required',
      componentProps: { placeholder: '请输入订单编号', maxlength: 32 },
    },
    {
      component: 'InputNumber',
      fieldName: 'payAmount',
      label: '实付金额（分）',
      componentProps: { placeholder: '请输入实付金额（分）', class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'coachID',
      label: '陪玩师ID',
      rules: 'selectRequired',
      componentProps: { options: coachIDOptions, placeholder: '请选择陪玩师ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'shopID',
      label: '店铺ID（0表示无店铺）',
      componentProps: { options: shopIDOptions, placeholder: '请选择店铺ID（0表示无店铺）', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'platformRate',
      label: '平台抽成比例（百分比）',
      componentProps: { placeholder: '请输入平台抽成比例（百分比）' },
    },
    {
      component: 'InputNumber',
      fieldName: 'platformAmount',
      label: '平台抽成金额（分）',
      componentProps: { placeholder: '请输入平台抽成金额（分）', class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'shopRate',
      label: '店铺抽成比例（百分比）',
      componentProps: { placeholder: '请输入店铺抽成比例（百分比）' },
    },
    {
      component: 'InputNumber',
      fieldName: 'shopAmount',
      label: '店铺抽成金额（分）',
      componentProps: { placeholder: '请输入店铺抽成金额（分）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'coachAmount',
      label: '陪玩师收入（分）',
      componentProps: { placeholder: '请输入陪玩师收入（分）', class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'settleStatus',
      label: '结算状态',
      componentProps: { options: settleStatusOptions, placeholder: '请选择结算状态', allowClear: true, class: 'w-full' },
    },
    {
      component: 'DatePicker',
      fieldName: 'settleAt',
      label: '结算时间',
      componentProps: { showTime: true, placeholder: '请选择结算时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
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
        modalApi.setState({ title: '编辑利润分成流水表' });
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
        modalApi.setState({ title: '新建利润分成流水表' });
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
