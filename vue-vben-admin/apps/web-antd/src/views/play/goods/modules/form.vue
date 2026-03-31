<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getGoodsDetail,
  createGoods,
  updateGoods,
} from '#/api/play/goods';

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Select',
      fieldName: 'categoryID',
      label: '分类ID',
      rules: 'selectRequired',
      componentProps: { options: categoryIDOptions, placeholder: '请选择分类ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'coachID',
      label: '陪玩师ID',
      rules: 'selectRequired',
      componentProps: { options: coachIDOptions, placeholder: '请选择陪玩师ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'title',
      label: '商品名称',
      rules: 'required',
      componentProps: { placeholder: '请输入商品名称', maxlength: 100 },
    },
    {
      component: 'ImageUpload',
      fieldName: 'coverImage',
      label: '商品封面图',
      componentProps: { maxCount: 1 },
    },
    {
      component: 'JsonEditor',
      fieldName: 'descContent',
      label: '商品详情描述',
      formItemClass: 'col-span-full',
    },
    {
      component: 'InputNumber',
      fieldName: 'price',
      label: () => h('span', {}, ['单价 ', h(Tooltip, { title: '分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入单价（分）', class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'unit',
      label: () => h('span', {}, ['计量单位 ', h(Tooltip, { title: '如：小时、局、次' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入计量单位（如', maxlength: 20 },
    },
    {
      component: 'InputNumber',
      fieldName: 'salesNum',
      label: '销量',
      componentProps: { placeholder: '请输入销量', class: 'w-full' },
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
        await updateGoods({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createGoods(values);
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
        modalApi.setState({ title: '编辑商品表' });
        try {
          const detail = await getGoodsDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建商品表' });
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
