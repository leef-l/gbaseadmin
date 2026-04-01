<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getActivityStepLogDetail,
  createActivityStepLog,
  updateActivityStepLog,
} from '#/api/play/activity_step_log';
import { getActivityList } from '#/api/play/activity';

/** 步骤类型选项 */
const stepTypeOptions = [
  { label: '文字 2=链接 3=图片', value: ��1 },
];

/** 审核状态选项 */
const auditStatusOptions = [
  { label: '待审核 1=通过 2=驳回', value: ��0 },
];

const activityIDOptions = ref<{ label: string; value: string }[]>([]);
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
      componentProps: { options: activityIDOptions, placeholder: '请选择活动ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'stepID',
      label: '步骤ID',
      componentProps: { options: stepIDOptions, placeholder: '请选择步骤ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'joinID',
      label: '参与记录ID',
      componentProps: { options: joinIDOptions, placeholder: '请选择参与记录ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'memberID',
      label: '会员ID',
      componentProps: { options: memberIDOptions, placeholder: '请选择会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'stepType',
      label: '步骤类型',
      componentProps: { options: stepTypeOptions, placeholder: '请选择步骤类型', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Textarea',
      fieldName: 'submitText',
      label: '用户提交的文字或链接',
      componentProps: { placeholder: '请输入用户提交的文字或链接', rows: 4, maxlength: 65535 },
    },
    {
      component: 'ImageUpload',
      fieldName: 'submitImage',
      label: '用户提交的图片URL',
      componentProps: { maxCount: 1 },
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
      component: 'Input',
      fieldName: 'auditBy',
      label: '审核人ID',
      componentProps: { placeholder: '请输入审核人ID' },
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
        await updateActivityStepLog({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createActivityStepLog(values);
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
      // 加载活动ID选项
      try {
        const activityRes = await getActivityList({ pageNum: 1, pageSize: 1000 });
        activityIDOptions.value = (activityRes?.list ?? []).map((item: any) => ({
          label: item.title || item.id,
          value: item.id,
        }));
      } catch {
        // ignore
      }
      if (data?.id) {
        isEdit.value = true;
        editId.value = data.id;
        modalApi.setState({ title: '编辑活动步骤提交记录' });
        try {
          const detail = await getActivityStepLogDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建活动步骤提交记录' });
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
