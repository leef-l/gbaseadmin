<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getActivityDetail,
  createActivity,
  updateActivity,
} from '#/api/play/activity';

/** 活动类型选项 */
const typeOptions = [
  { label: '充值活动', value: 1 },
  { label: '下单活动', value: 2 },
  { label: '注册活动', value: 3 },
  { label: '图文步骤活动', value: 4 },
  { label: '自定义活动', value: 5 },
];

/** 参与条件选项 */
const conditionTypeOptions = [
  { label: '无条件', value: 0 },
  { label: '需报名', value: 1 },
  { label: '充值满额', value: 2 },
  { label: '下单满额', value: 3 },
  { label: '完成步骤', value: 4 },
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
      label: '活动名称',
      rules: 'required',
      componentProps: { placeholder: '请输入活动名称', maxlength: 100 },
    },
    {
      component: 'ImageUpload',
      fieldName: 'coverImage',
      label: '活动封面图',
      componentProps: { maxCount: 1 },
    },
    {
      component: 'RichText',
      fieldName: 'descContent',
      label: '活动详情描述',
      formItemClass: 'col-span-full',
    },
    {
      component: 'Select',
      fieldName: 'type',
      label: '活动类型',
      componentProps: { options: typeOptions, placeholder: '请选择活动类型', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'conditionType',
      label: '参与条件',
      componentProps: { options: conditionTypeOptions, placeholder: '请选择参与条件', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'conditionValue',
      label: () => h('span', {}, ['条件值 ', h(Tooltip, { title: '分/次，如充值满5000分、下单满3次' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入条件值（分/次，如充值满5000分、下单满3次）' },
    },
    {
      component: 'Switch',
      fieldName: 'isAutoReward',
      label: '是否自动发奖',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 1,
    },
    {
      component: 'DatePicker',
      fieldName: 'startAt',
      label: '活动开始时间',
      rules: 'required',
      componentProps: { showTime: true, placeholder: '请选择活动开始时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'endAt',
      label: '活动结束时间',
      rules: 'required',
      componentProps: { showTime: true, placeholder: '请选择活动结束时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'InputNumber',
      fieldName: 'maxNum',
      label: () => h('span', {}, ['参与人数上限 ', h(Tooltip, { title: '0表示不限' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入参与人数上限（0表示不限）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'joinNum',
      label: '已参与人数',
      componentProps: { placeholder: '请输入已参与人数', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: () => h('span', {}, ['排序 ', h(Tooltip, { title: '升序' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
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
        modalApi.setState({ title: '编辑活动表' });
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
        modalApi.setState({ title: '新建活动表' });
        formApi.resetForm();
      }
    }
  },
});
</script>

<template>
  <Modal class="w-[860px]">
    <Form />
  </Modal>
</template>
