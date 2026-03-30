<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getCoachApplyDetail,
  createCoachApply,
  updateCoachApply,
} from '#/api/play/coach_apply';

/** 审核状态选项 */
const auditStatusOptions = [
  { label: '待审核', value: 0 },
  { label: '通过', value: 1 },
  { label: '拒绝', value: 2 },
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
      label: '会员ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'realName',
      label: '真实姓名',
      rules: 'required',
      componentProps: { placeholder: '请输入真实姓名', maxlength: 50 },
    },
    {
      component: 'Input',
      fieldName: 'idCard',
      label: '身份证号',
      rules: 'required',
      componentProps: { placeholder: '请输入身份证号', maxlength: 30 },
    },
    {
      component: 'Input',
      fieldName: 'idCardFrontImage',
      label: '身份证正面照',
      rules: 'required',
      componentProps: { placeholder: '请输入身份证正面照', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'idCardBackImage',
      label: '身份证反面照',
      rules: 'required',
      componentProps: { placeholder: '请输入身份证反面照', maxlength: 500 },
    },
    {
      component: 'Textarea',
      fieldName: 'skillDesc',
      label: '技能描述',
      componentProps: { placeholder: '请输入技能描述', rows: 4, maxlength: 65535 },
    },
    {
      component: 'Select',
      fieldName: 'auditStatus',
      label: '审核状态',
      componentProps: { options: auditStatusOptions, placeholder: '请选择审核状态', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'auditRemark',
      label: '审核备注',
      componentProps: { placeholder: '请输入审核备注', maxlength: 500 },
    },
    {
      component: 'DatePicker',
      fieldName: 'auditAt',
      label: '审核时间',
      componentProps: { showTime: true, placeholder: '请选择审核时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
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
        await updateCoachApply({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createCoachApply(values);
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
        modalApi.setState({ title: '编辑陪玩师申请表' });
        try {
          const detail = await getCoachApplyDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建陪玩师申请表' });
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
