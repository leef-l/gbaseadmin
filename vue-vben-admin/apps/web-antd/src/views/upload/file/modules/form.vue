<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getFileDetail,
  createFile,
  updateFile,
} from '#/api/upload/file';
import { getDirTree } from '#/api/upload/dir';
import type { DirItem } from '#/api/upload/dir/types';

/** å­˜å‚¨ç±»åž‹选项 */
const storageOptions = [
  { label: 'æœ¬åœ°', value: 1 },
  { label: 'é˜¿é‡Œäº‘OSS', value: 2 },
  { label: 'è…¾è®¯äº‘COS', value: 3 },
];

/** æ˜¯å¦å›¾ç‰‡选项 */
const isImageOptions = [
  { label: 'å¦', value: 0 },
  { label: 'æ˜¯', value: 1 },
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
      fieldName: 'dirID',
      label: 'æ‰€å±žç›®å½•',
      componentProps: { options: dirIDOptions, placeholder: '请选择æ‰€å±žç›®å½•', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'name',
      label: 'æ–‡ä»¶åç§°',
      rules: 'required',
      componentProps: { placeholder: '请输入æ–‡ä»¶åç§°', maxlength: 255 },
    },
    {
      component: 'Input',
      fieldName: 'url',
      label: 'æ–‡ä»¶åœ°å€',
      rules: 'required',
      componentProps: { placeholder: '请输入æ–‡ä»¶åœ°å€', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'ext',
      label: 'æ–‡ä»¶æ‰©å±•å',
      componentProps: { placeholder: '请输入æ–‡ä»¶æ‰©å±•å', maxlength: 20 },
    },
    {
      component: 'Input',
      fieldName: 'size',
      label: 'æ–‡ä»¶å¤§å°',
      componentProps: { placeholder: '请输入æ–‡ä»¶å¤§å°' },
    },
    {
      component: 'Input',
      fieldName: 'mime',
      label: 'MIMEç±»åž‹',
      componentProps: { placeholder: '请输入MIMEç±»åž‹', maxlength: 100 },
    },
    {
      component: 'Select',
      fieldName: 'storage',
      label: 'å­˜å‚¨ç±»åž‹',
      componentProps: { options: storageOptions, placeholder: '请选择å­˜å‚¨ç±»åž‹', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'isImage',
      label: 'æ˜¯å¦å›¾ç‰‡',
      componentProps: { options: isImageOptions, placeholder: '请选择是否图片', allowClear: true, class: 'w-full' },
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
        await updateFile({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createFile(values);
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
      await loadDirOptions();
      const data = modalApi.getData<{ id?: string } | null>();
      if (data?.id) {
        isEdit.value = true;
        editId.value = data.id;
        modalApi.setState({ title: '编辑æ–‡ä»¶è®°å½•' });
        try {
          const detail = await getFileDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建æ–‡ä»¶è®°å½•' });
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
