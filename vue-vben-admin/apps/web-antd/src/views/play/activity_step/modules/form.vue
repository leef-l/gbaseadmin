<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getActivityStepDetail,
  createActivityStep,
  updateActivityStep,
} from '#/api/play/activity_step';
import { getActivityList } from '#/api/play/activity';

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
      component: 'JsonEditor',
      fieldName: 'descContent',
      label: '步骤说明',
      formItemClass: 'col-span-full',
    },
    {
      component: 'ImageUpload',
      fieldName: 'stepImage',
      label: '步骤示例图片',
      componentProps: { maxCount: 1 },
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: () => h('span', {}, ['排序 ', h(Tooltip, { title: '升序' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
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
      // 加载活动选项
      try {
        const activityRes = await getActivityList({ pageNum: 1, pageSize: 1000 });
        activityIDOptions.value = (activityRes?.list ?? []).map((item: any) => ({
          label: item.title || item.id,
          value: item.id,
        }));
      } catch {
        // ignore
      }

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
