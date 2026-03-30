<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getActivityStepDetail,
  createActivityStep,
  updateActivityStep,
} from '#/api/play/activity_step';

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Select',
      fieldName: 'activityID',
      label: 'æ´»åŠ¨ID',
      rules: 'selectRequired',
      componentProps: { options: activityIDOptions, placeholder: '请选择æ´»åŠ¨ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'stepNum',
      label: 'æ­¥éª¤åºå·',
      componentProps: { placeholder: '请输入æ­¥éª¤åºå·', class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'title',
      label: 'æ­¥éª¤æ ‡é¢˜',
      rules: 'required',
      componentProps: { placeholder: '请输入æ­¥éª¤æ ‡é¢˜', maxlength: 100 },
    },
    {
      component: 'Input',
      fieldName: 'descContent',
      label: 'æ­¥éª¤è¯´æ˜Ž',
      componentProps: { placeholder: '请输入æ­¥éª¤è¯´æ˜Ž', maxlength: 65535 },
    },
    {
      component: 'Input',
      fieldName: 'stepImage',
      label: 'æ­¥éª¤ç¤ºä¾‹å›¾ç‰‡',
      componentProps: { placeholder: '请输入æ­¥éª¤ç¤ºä¾‹å›¾ç‰‡', maxlength: 500 },
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: 'æŽ’åº',
      componentProps: { placeholder: '请输入æŽ’åº', class: 'w-full' },
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
        await updateActivityStep({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createActivityStep(values);
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
        modalApi.setState({ title: '编辑æ´»åŠ¨æ­¥éª¤è¡¨' });
        try {
          const detail = await getActivityStepDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建æ´»åŠ¨æ­¥éª¤è¡¨' });
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
