<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getReviewDetail,
  createReview,
  updateReview,
} from '#/api/play/review';

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
      label: 'è¯„ä»·ä¼šå‘˜ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择è¯„ä»·ä¼šå‘˜ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'coachID',
      label: 'è¢«è¯„é™ªçŽ©å¸ˆID',
      rules: 'selectRequired',
      componentProps: { options: coachIDOptions, placeholder: '请选择è¢«è¯„é™ªçŽ©å¸ˆID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'score',
      label: 'è¯„åˆ†ï¼ˆä¹˜100ï¼‰',
      componentProps: { placeholder: '请输入è¯„åˆ†ï¼ˆä¹˜100ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'reviewContent',
      label: 'è¯„ä»·å†…å®¹',
      componentProps: { placeholder: '请输入è¯„ä»·å†…å®¹', maxlength: 65535 },
    },
    {
      component: 'Input',
      fieldName: 'reviewImage',
      label: 'è¯„ä»·å›¾ç‰‡ï¼ˆå¤šå¼ é€—å·åˆ†éš”ï¼‰',
      componentProps: { placeholder: '请输入è¯„ä»·å›¾ç‰‡ï¼ˆå¤šå¼ é€—å·åˆ†éš”ï¼‰', maxlength: 2000 },
    },
    {
      component: 'Input',
      fieldName: 'replyContent',
      label: 'é™ªçŽ©å¸ˆå›žå¤å†…å®¹',
      componentProps: { placeholder: '请输入é™ªçŽ©å¸ˆå›žå¤å†…å®¹', maxlength: 65535 },
    },
    {
      component: 'DatePicker',
      fieldName: 'replyAt',
      label: 'å›žå¤æ—¶é—´',
      componentProps: { showTime: true, placeholder: '请选择å›žå¤æ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'Switch',
      fieldName: 'isAnonymous',
      label: 'æ˜¯å¦åŒ¿å',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 0,
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
        await updateReview({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createReview(values);
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
        modalApi.setState({ title: '编辑è¯„ä»·è¡¨' });
        try {
          const detail = await getReviewDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建è¯„ä»·è¡¨' });
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
