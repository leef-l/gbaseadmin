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

/** ä¼˜æƒ åˆ¸ç±»åž‹选项 */
const typeOptions = [
  { label: 'æ»¡å‡åˆ¸', value: 1 },
  { label: 'æŠ˜æ‰£åˆ¸', value: 2 },
  { label: 'æ— é—¨æ§›åˆ¸', value: 3 },
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
      label: 'ä¼˜æƒ åˆ¸åç§°',
      rules: 'required',
      componentProps: { placeholder: '请输入ä¼˜æƒ åˆ¸åç§°', maxlength: 100 },
    },
    {
      component: 'Select',
      fieldName: 'type',
      label: 'ä¼˜æƒ åˆ¸ç±»åž‹',
      componentProps: { options: typeOptions, placeholder: '请选择ä¼˜æƒ åˆ¸ç±»åž‹', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Switch',
      fieldName: 'isNewMember',
      label: 'æ˜¯å¦æ–°äººä¸“äº«',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 0,
    },
    {
      component: 'Input',
      fieldName: 'faceValue',
      label: 'é¢å€¼ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入é¢å€¼ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'minAmount',
      label: 'æœ€ä½Žæ¶ˆè´¹é‡‘é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入æœ€ä½Žæ¶ˆè´¹é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'InputNumber',
      fieldName: 'totalNum',
      label: 'å‘æ”¾æ€»é‡ï¼ˆ0ä¸é™ï¼‰',
      componentProps: { placeholder: '请输入å‘æ”¾æ€»é‡ï¼ˆ0ä¸é™ï¼‰', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'usedNum',
      label: 'å·²ä½¿ç”¨æ•°é‡',
      componentProps: { placeholder: '请输入å·²ä½¿ç”¨æ•°é‡', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'claimNum',
      label: 'å·²é¢†å–æ•°é‡',
      componentProps: { placeholder: '请输入å·²é¢†å–æ•°é‡', class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'perLimit',
      label: 'æ¯äººé™é¢†å¼ æ•°',
      componentProps: { placeholder: '请输入æ¯äººé™é¢†å¼ æ•°' },
    },
    {
      component: 'DatePicker',
      fieldName: 'validStartAt',
      label: 'æœ‰æ•ˆæœŸå¼€å§‹æ—¶é—´',
      rules: 'required',
      componentProps: { showTime: true, placeholder: '请选择æœ‰æ•ˆæœŸå¼€å§‹æ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'validEndAt',
      label: 'æœ‰æ•ˆæœŸç»“æŸæ—¶é—´',
      rules: 'required',
      componentProps: { showTime: true, placeholder: '请选择æœ‰æ•ˆæœŸç»“æŸæ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: 'æŽ’åº',
      componentProps: { placeholder: '请输入æŽ’åº', class: 'w-full' },
    },
    {
      component: 'Switch',
      fieldName: 'status',
      label: 'çŠ¶æ€',
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
        modalApi.setState({ title: '编辑ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨' });
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
        modalApi.setState({ title: '新建ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨' });
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
