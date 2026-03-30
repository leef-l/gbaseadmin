<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getCouponDetail,
  createCoupon,
  updateCoupon,
} from '#/api/play/coupon';

/** 优惠券类型选项 */
const typeOptions = [
  { label: '满减券', value: 1 },
  { label: '折扣券', value: 2 },
  { label: '无门槛券', value: 3 },
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
      fieldName: 'title',
      label: '优惠券名称',
      rules: 'required',
      componentProps: { placeholder: '请输入优惠券名称', maxlength: 100 },
    },
    {
      component: 'Select',
      fieldName: 'type',
      label: '优惠券类型',
      componentProps: { options: typeOptions, placeholder: '请选择优惠券类型', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Switch',
      fieldName: 'isNewMember',
      label: '是否新人专享',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 0,
    },
    {
      component: 'Input',
      fieldName: 'faceValue',
      label: '面值（分，满减/无门槛时为抵扣额，折扣时为折扣值如 85=8.5折）',
      componentProps: { placeholder: '请输入面值（分，满减/无门槛时为抵扣额，折扣时为折扣值如 85=8.5折）' },
    },
    {
      component: 'InputNumber',
      fieldName: 'minAmount',
      label: '最低消费金额（分，0表示无门槛）',
      componentProps: { placeholder: '请输入最低消费金额（分，0表示无门槛）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'totalNum',
      label: '发放总量（0表示不限）',
      componentProps: { placeholder: '请输入发放总量（0表示不限）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'usedNum',
      label: '已使用数量',
      componentProps: { placeholder: '请输入已使用数量', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'claimNum',
      label: '已领取数量',
      componentProps: { placeholder: '请输入已领取数量', class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'perLimit',
      label: '每人限领张数',
      componentProps: { placeholder: '请输入每人限领张数' },
    },
    {
      component: 'DatePicker',
      fieldName: 'validStartAt',
      label: '有效期开始时间',
      rules: 'required',
      componentProps: { showTime: true, placeholder: '请选择有效期开始时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'validEndAt',
      label: '有效期结束时间',
      rules: 'required',
      componentProps: { showTime: true, placeholder: '请选择有效期结束时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: '排序（升序）',
      componentProps: { placeholder: '请输入排序（升序）', class: 'w-full' },
    },
    {
      component: 'Switch',
      fieldName: 'status',
      label: '状态',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 1,
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
        await updateCoupon({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createCoupon(values);
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
        modalApi.setState({ title: '编辑优惠券模板表' });
        try {
          const detail = await getCouponDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建优惠券模板表' });
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
