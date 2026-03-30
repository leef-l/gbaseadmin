<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getRechargePlanDetail,
  createRechargePlan,
  updateRechargePlan,
} from '#/api/play/recharge_plan';

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
      label: 'æ–¹æ¡ˆåç§°',
      rules: 'required',
      componentProps: { placeholder: '请输入æ–¹æ¡ˆåç§°', maxlength: 50 },
    },
    {
      component: 'Input',
      fieldName: 'amount',
      label: 'å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰',
      rules: 'required',
      componentProps: { placeholder: '请输入å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'giftAmount',
      label: 'èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'coverImage',
      label: 'æ–¹æ¡ˆå°é¢å›¾',
      componentProps: { placeholder: '请输入æ–¹æ¡ˆå°é¢å›¾', maxlength: 500 },
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: 'æŽ’åºï¼ˆå‡åºï¼‰',
      componentProps: { placeholder: '请输入æŽ’åºï¼ˆå‡åºï¼‰', class: 'w-full' },
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
        await updateRechargePlan({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createRechargePlan(values);
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
        modalApi.setState({ title: '编辑å……å€¼æ–¹æ¡ˆè¡¨' });
        try {
          const detail = await getRechargePlanDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建å……å€¼æ–¹æ¡ˆè¡¨' });
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
