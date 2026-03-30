<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getBalanceLogDetail,
  createBalanceLog,
  updateBalanceLog,
} from '#/api/play/balance_log';

/** ä¸šåŠ¡ç±»åž‹选项 */
const bizTypeOptions = [
  { label: 'å……å€¼', value: 1 },
  { label: 'æ¶ˆè´¹', value: 2 },
  { label: 'é€€æ¬¾', value: 3 },
  { label: 'æ´»åŠ¨èµ é€', value: 4 },
  { label: 'æçŽ°', value: 5 },
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
      fieldName: 'memberID',
      label: 'ä¼šå‘˜ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择ä¼šå‘˜ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'bizType',
      label: 'ä¸šåŠ¡ç±»åž‹',
      rules: 'selectRequired',
      componentProps: { options: bizTypeOptions, placeholder: '请选择ä¸šåŠ¡ç±»åž‹', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'bizID',
      label: 'å…³è”ä¸šåŠ¡ID',
      componentProps: { options: bizIDOptions, placeholder: '请选择å…³è”ä¸šåŠ¡ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'changeAmount',
      label: 'å˜åŠ¨é‡‘é¢ï¼ˆåˆ†ï¼‰',
      rules: 'required',
      componentProps: { placeholder: '请输入å˜åŠ¨é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'beforeBalance',
      label: 'å˜åŠ¨å‰ä½™é¢ï¼ˆåˆ†ï¼‰',
      rules: 'required',
      componentProps: { placeholder: '请输入å˜åŠ¨å‰ä½™é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'afterBalance',
      label: 'å˜åŠ¨åŽä½™é¢ï¼ˆåˆ†ï¼‰',
      rules: 'required',
      componentProps: { placeholder: '请输入å˜åŠ¨åŽä½™é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'remark',
      label: 'å¤‡æ³¨è¯´æ˜Ž',
      componentProps: { placeholder: '请输入å¤‡æ³¨è¯´æ˜Ž', maxlength: 200 },
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
        await updateBalanceLog({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createBalanceLog(values);
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
        modalApi.setState({ title: '编辑ä½™é¢æµæ°´è¡¨' });
        try {
          const detail = await getBalanceLogDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建ä½™é¢æµæ°´è¡¨' });
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
