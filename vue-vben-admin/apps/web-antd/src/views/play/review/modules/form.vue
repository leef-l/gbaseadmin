<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getReviewDetail,
  createReview,
  updateReview,
} from '#/api/play/review';

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Select',
      fieldName: 'orderID',
      label: '订单ID',
      rules: 'selectRequired',
      componentProps: { options: orderIDOptions, placeholder: '请选择订单ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'memberID',
      label: '评价会员ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择评价会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'coachID',
      label: '被评陪玩师ID',
      rules: 'selectRequired',
      componentProps: { options: coachIDOptions, placeholder: '请选择被评陪玩师ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'score',
      label: () => h('span', {}, ['评分 ', h(Tooltip, { title: '乘100，如 500=5.00分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入评分（乘100，如 500=5.00分）' },
    },
    {
      component: 'Input',
      fieldName: 'reviewContent',
      label: '评价内容',
      componentProps: { placeholder: '请输入评价内容', maxlength: 65535 },
    },
    {
      component: 'Input',
      fieldName: 'reviewImage',
      label: () => h('span', {}, ['评价图片 ', h(Tooltip, { title: '多张逗号分隔' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入评价图片（多张逗号分隔）', maxlength: 2000 },
    },
    {
      component: 'Input',
      fieldName: 'replyContent',
      label: '陪玩师回复内容',
      componentProps: { placeholder: '请输入陪玩师回复内容', maxlength: 65535 },
    },
    {
      component: 'DatePicker',
      fieldName: 'replyAt',
      label: '回复时间',
      componentProps: { showTime: true, placeholder: '请选择回复时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'Switch',
      fieldName: 'isAnonymous',
      label: '是否匿名',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 0,
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
        await updateReview({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createReview(values);
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
        modalApi.setState({ title: '编辑评价表' });
        try {
          const detail = await getReviewDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建评价表' });
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
