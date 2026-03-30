<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getActivityRewardDetail,
  createActivityReward,
  updateActivityReward,
} from '#/api/play/activity_reward';

/** å¥–åŠ±ç±»åž‹选项 */
const rewardTypeOptions = [
  { label: 'ä½™é¢', value: 1 },
  { label: 'ä¼˜æƒ åˆ¸', value: 2 },
  { label: 'ç»éªŒå€¼', value: 3 },
  { label: 'ä¼šå‘˜ç­‰çº§å¤©æ•°', value: 4 },
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
      fieldName: 'activityID',
      label: 'æ´»åŠ¨ID',
      rules: 'selectRequired',
      componentProps: { options: activityIDOptions, placeholder: '请选择æ´»åŠ¨ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'rewardType',
      label: 'å¥–åŠ±ç±»åž‹',
      componentProps: { options: rewardTypeOptions, placeholder: '请选择å¥–åŠ±ç±»åž‹', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'rewardValue',
      label: 'å¥–åŠ±æ•°å€¼',
      componentProps: { placeholder: '请输入å¥–åŠ±æ•°å€¼' },
    },
    {
      component: 'Input',
      fieldName: 'rewardName',
      label: 'å¥–åŠ±åç§°',
      rules: 'required',
      componentProps: { placeholder: '请输入å¥–åŠ±åç§°', maxlength: 100 },
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
        await updateActivityReward({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createActivityReward(values);
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
        modalApi.setState({ title: '编辑æ´»åŠ¨å¥–åŠ±è¡¨' });
        try {
          const detail = await getActivityRewardDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建æ´»åŠ¨å¥–åŠ±è¡¨' });
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
