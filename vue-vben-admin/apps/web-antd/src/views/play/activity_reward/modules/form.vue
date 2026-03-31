<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getActivityRewardDetail,
  createActivityReward,
  updateActivityReward,
} from '#/api/play/activity_reward';

/** 奖励类型选项 */
const rewardTypeOptions = [
  { label: '余额', value: 1 },
  { label: '优惠券', value: 2 },
  { label: '经验值', value: 3 },
  { label: '会员等级天数', value: 4 },
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
      label: '活动ID',
      rules: 'selectRequired',
      componentProps: { options: activityIDOptions, placeholder: '请选择活动ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'rewardType',
      label: '奖励类型',
      componentProps: { options: rewardTypeOptions, placeholder: '请选择奖励类型', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'rewardValue',
      label: () => h('span', {}, ['奖励数值 ', h(Tooltip, { title: '余额=分，优惠券=coupon_id，经验=值，等级天数=天' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入奖励数值（余额=分，优惠券=coupon_id，经验=值，等级天数=天）' },
    },
    {
      component: 'Input',
      fieldName: 'rewardName',
      label: () => h('span', {}, ['奖励名称 ', h(Tooltip, { title: '展示用，如"送50元余额"' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      rules: 'required',
      componentProps: { placeholder: '请输入奖励名称（展示用，如"送50元余额"）', maxlength: 100 },
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
        modalApi.setState({ title: '编辑活动奖励表' });
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
        modalApi.setState({ title: '新建活动奖励表' });
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
