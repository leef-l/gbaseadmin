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
      label: '活动ID',
      rules: 'selectRequired',
      componentProps: { options: activityIDOptions, placeholder: '请选择活动ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'stepNum',
      label: '步骤序号',
      componentProps: { placeholder: '请输入步骤序号', class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'title',
      label: '步骤标题',
      rules: 'required',
      componentProps: { placeholder: '请输入步骤标题', maxlength: 100 },
    },
    {
      component: 'Input',
      fieldName: 'descContent',
      label: '步骤说明（富文本，支持图文）',
      componentProps: { placeholder: '请输入步骤说明（富文本，支持图文）', maxlength: 65535 },
    },
    {
      component: 'Input',
      fieldName: 'stepImage',
      label: '步骤示例图片',
      componentProps: { placeholder: '请输入步骤示例图片', maxlength: 500 },
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: '排序（升序）',
      componentProps: { placeholder: '请输入排序（升序）', class: 'w-full' },
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
        modalApi.setState({ title: '编辑活动步骤表' });
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
        modalApi.setState({ title: '新建活动步骤表' });
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
