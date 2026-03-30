<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getActivityDetail,
  createActivity,
  updateActivity,
} from '#/api/play/activity';

/** æ´»åŠ¨ç±»åž‹选项 */
const typeOptions = [
  { label: 'å……å€¼æ´»åŠ¨', value: 1 },
  { label: 'ä¸‹å•æ´»åŠ¨', value: 2 },
  { label: 'æ³¨å†Œæ´»åŠ¨', value: 3 },
  { label: 'å›¾æ–‡æ­¥éª¤æ´»åŠ¨', value: 4 },
  { label: 'è‡ªå®šä¹‰æ´»åŠ¨', value: 5 },
];

/** å‚ä¸Žæ¡ä»¶选项 */
const conditionTypeOptions = [
  { label: 'æ— æ¡ä»¶', value: 0 },
  { label: 'éœ€æŠ¥å', value: 1 },
  { label: 'å……å€¼æ»¡é¢', value: 2 },
  { label: 'ä¸‹å•æ»¡é¢', value: 3 },
  { label: 'å®Œæˆæ­¥éª¤', value: 4 },
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
      label: 'æ´»åŠ¨åç§°',
      rules: 'required',
      componentProps: { placeholder: '请输入æ´»åŠ¨åç§°', maxlength: 100 },
    },
    {
      component: 'Input',
      fieldName: 'coverImage',
      label: 'æ´»åŠ¨å°é¢å›¾',
      componentProps: { placeholder: '请输入æ´»åŠ¨å°é¢å›¾', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'descContent',
      label: 'æ´»åŠ¨è¯¦æƒ…æè¿°',
      componentProps: { placeholder: '请输入æ´»åŠ¨è¯¦æƒ…æè¿°', maxlength: 65535 },
    },
    {
      component: 'Select',
      fieldName: 'type',
      label: 'æ´»åŠ¨ç±»åž‹',
      componentProps: { options: typeOptions, placeholder: '请选择æ´»åŠ¨ç±»åž‹', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'conditionType',
      label: 'å‚ä¸Žæ¡ä»¶',
      componentProps: { options: conditionTypeOptions, placeholder: '请选择å‚ä¸Žæ¡ä»¶', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'conditionValue',
      label: 'æ¡ä»¶å€¼',
      componentProps: { placeholder: '请输入æ¡ä»¶å€¼' },
    },
    {
      component: 'Switch',
      fieldName: 'isAutoReward',
      label: 'æ˜¯å¦è‡ªåŠ¨å‘å¥–',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 1,
    },
    {
      component: 'DatePicker',
      fieldName: 'startAt',
      label: 'æ´»åŠ¨å¼€å§‹æ—¶é—´',
      rules: 'required',
      componentProps: { showTime: true, placeholder: '请选择æ´»åŠ¨å¼€å§‹æ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'endAt',
      label: 'æ´»åŠ¨ç»“æŸæ—¶é—´',
      rules: 'required',
      componentProps: { showTime: true, placeholder: '请选择æ´»åŠ¨ç»“æŸæ—¶é—´', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'InputNumber',
      fieldName: 'maxNum',
      label: 'å‚ä¸Žäººæ•°ä¸Šé™ï¼ˆ0ä¸é™ï¼‰',
      componentProps: { placeholder: '请输入å‚ä¸Žäººæ•°ä¸Šé™ï¼ˆ0ä¸é™ï¼‰', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'joinNum',
      label: 'å·²å‚ä¸Žäººæ•°',
      componentProps: { placeholder: '请输入å·²å‚ä¸Žäººæ•°', class: 'w-full' },
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
        await updateActivity({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createActivity(values);
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
        modalApi.setState({ title: '编辑æ´»åŠ¨è¡¨' });
        try {
          const detail = await getActivityDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建æ´»åŠ¨è¡¨' });
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
