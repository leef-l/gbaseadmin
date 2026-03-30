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
      label: '方案名称',
      rules: 'required',
      componentProps: { placeholder: '请输入方案名称', maxlength: 50 },
    },
    {
      component: 'InputNumber',
      fieldName: 'amount',
      label: '充值金额（分）',
      rules: 'required',
      componentProps: { placeholder: '请输入充值金额（分）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'giftAmount',
      label: '赠送金额（分）',
      componentProps: { placeholder: '请输入赠送金额（分）', class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'coverImage',
      label: '方案封面图',
      componentProps: { placeholder: '请输入方案封面图', maxlength: 500 },
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: '排序（升序）',
      componentProps: { placeholder: '请输入排序（升序）', class: 'w-full' },
    },
    {
      component: 'Switch',
      fieldName: 'status',
      label: '状态',
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
        modalApi.setState({ title: '编辑充值方案表' });
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
        modalApi.setState({ title: '新建充值方案表' });
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
