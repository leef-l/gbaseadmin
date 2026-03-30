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

/** ä½¿ç”¨çŠ¶æ€选项 */
const useStatusOptions = [
  { label: 'æœªä½¿ç”¨', value: 0 },
  { label: 'å·²ä½¿ç”¨', value: 1 },
  { label: 'å·²è¿‡æœŸ', value: 2 },
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
      label: 'ä¼˜æƒ åˆ¸æ¨¡æ¿ID',
      rules: 'selectRequired',
      componentProps: { options: couponIDOptions, placeholder: '请选择ä¼˜æƒ åˆ¸æ¨¡æ¿ID', allowClear: true, class: 'w-full' },
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
      fieldName: 'orderID',
      label: 'ä½¿ç”¨çš„è®¢å•ID',
      componentProps: { options: orderIDOptions, placeholder: '请选择ä½¿ç”¨çš„è®¢å•ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'useStatus',
      label: 'ä½¿ç”¨çŠ¶æ€',
      componentProps: { options: useStatusOptions, placeholder: '请选择ä½¿ç”¨çŠ¶æ€', allowClear: true, class: 'w-full' },
    },
    {
      component: 'DatePicker',
      fieldName: 'claimAt',
      label: 'é¢†å–æ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择é¢†å–æ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'useAt',
      label: 'ä½¿ç”¨æ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择ä½¿ç”¨æ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'expireAt',
      label: 'è¿‡æœŸæ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择è¿‡æœŸæ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
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
        modalApi.setState({ title: '编辑ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨' });
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
        modalApi.setState({ title: '新建ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨' });
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
