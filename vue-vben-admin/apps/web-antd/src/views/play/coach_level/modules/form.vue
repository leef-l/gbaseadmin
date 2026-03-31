<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getCoachLevelDetail,
  createCoachLevel,
  updateCoachLevel,
} from '#/api/play/coach_level';

/** 等级选项 */
const levelOptions = [
  { label: '青铜', value: 1 },
  { label: '白银', value: 2 },
  { label: '黄金', value: 3 },
  { label: '铂金', value: 4 },
  { label: '钻石', value: 5 },
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
      label: '等级名称',
      rules: 'required',
      componentProps: { placeholder: '请输入等级名称', maxlength: 50 },
    },
    {
      component: 'Select',
      fieldName: 'level',
      label: '等级',
      componentProps: { options: levelOptions, placeholder: '请选择等级', allowClear: true, class: 'w-full' },
    },
    {
      component: 'ImageUpload',
      fieldName: 'icon',
      label: '等级图标',
      componentProps: { maxCount: 1 },
    },
    {
      component: 'Input',
      fieldName: 'minOrders',
      label: '所需最低接单数',
      componentProps: { placeholder: '请输入所需最低接单数' },
    },
    {
      component: 'Input',
      fieldName: 'minScore',
      label: () => h('span', {}, ['所需最低评分 ', h(Tooltip, { title: '乘100存储，如 450=4.50分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入所需最低评分（乘100存储，如 450=4.50分）' },
    },
    {
      component: 'Input',
      fieldName: 'commissionRate',
      label: () => h('span', {}, ['平台抽成比例 ', h(Tooltip, { title: '百分比，如 20 表示 20%' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入平台抽成比例（百分比，如 20 表示 20%）' },
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
        await updateCoachLevel({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createCoachLevel(values);
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
        modalApi.setState({ title: '编辑陪玩师等级表' });
        try {
          const detail = await getCoachLevelDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建陪玩师等级表' });
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
