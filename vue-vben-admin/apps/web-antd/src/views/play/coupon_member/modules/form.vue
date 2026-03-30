<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getCouponMemberDetail,
  createCouponMember,
  updateCouponMember,
} from '#/api/play/coupon_member';

/** 使用状态选项 */
const useStatusOptions = [
  { label: '未使用', value: 0 },
  { label: '已使用', value: 1 },
  { label: '已过期', value: 2 },
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
      fieldName: 'couponID',
      label: '优惠券模板ID',
      rules: 'selectRequired',
      componentProps: { options: couponIDOptions, placeholder: '请选择优惠券模板ID', allowClear: true, class: 'w-full' },
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
      fieldName: 'orderID',
      label: '使用的订单ID（0表示未使用）',
      componentProps: { options: orderIDOptions, placeholder: '请选择使用的订单ID（0表示未使用）', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'useStatus',
      label: '使用状态',
      componentProps: { options: useStatusOptions, placeholder: '请选择使用状态', allowClear: true, class: 'w-full' },
    },
    {
      component: 'DatePicker',
      fieldName: 'claimAt',
      label: '领取时间',
      componentProps: { showTime: true, placeholder: '请选择领取时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'useAt',
      label: '使用时间',
      componentProps: { showTime: true, placeholder: '请选择使用时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'expireAt',
      label: '过期时间',
      componentProps: { showTime: true, placeholder: '请选择过期时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
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
        await updateCouponMember({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createCouponMember(values);
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
        modalApi.setState({ title: '编辑会员优惠券表' });
        try {
          const detail = await getCouponMemberDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建会员优惠券表' });
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
