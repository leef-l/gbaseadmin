<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getCoachDetail,
  createCoach,
  updateCoach,
} from '#/api/play/coach';

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
      label: '关联会员ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择关联会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'coachLevelID',
      label: '陪玩师等级ID',
      componentProps: { options: coachLevelIDOptions, placeholder: '请选择陪玩师等级ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'shopID',
      label: () => h('span', {}, ['所属店铺ID ', h(Tooltip, { title: '0表示无店铺' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { options: shopIDOptions, placeholder: '请选择所属店铺ID（0表示无店铺）', allowClear: true, class: 'w-full' },
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
      fieldName: 'intro',
      label: '个人简介',
      componentProps: { placeholder: '请输入个人简介', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'coverImage',
      label: '封面图',
      componentProps: { placeholder: '请输入封面图', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'totalOrders',
      label: '总接单数',
      componentProps: { placeholder: '请输入总接单数' },
    },
    {
      component: 'Input',
      fieldName: 'totalScore',
      label: () => h('span', {}, ['总评分 ', h(Tooltip, { title: '乘100，如 500=5.00' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入总评分（乘100，如 500=5.00）' },
    },
    {
      component: 'InputNumber',
      fieldName: 'scoreNum',
      label: '评分人数',
      componentProps: { placeholder: '请输入评分人数', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'incomeTotal',
      label: () => h('span', {}, ['累计收入 ', h(Tooltip, { title: '分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入累计收入（分）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'incomeBalance',
      label: () => h('span', {}, ['可提现余额 ', h(Tooltip, { title: '分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入可提现余额（分）', class: 'w-full' },
    },
    {
      component: 'Switch',
      fieldName: 'isOnline',
      label: '是否在线',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 0,
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
        await updateCoach({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createCoach(values);
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
        modalApi.setState({ title: '编辑陪玩师表' });
        try {
          const detail = await getCoachDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建陪玩师表' });
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
