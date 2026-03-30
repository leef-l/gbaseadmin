<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getShopDetail,
  createShop,
  updateShop,
} from '#/api/play/shop';

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
      label: '店铺名称',
      rules: 'required',
      componentProps: { placeholder: '请输入店铺名称', maxlength: 100 },
    },
    {
      component: 'Input',
      fieldName: 'logoImage',
      label: '店铺LOGO',
      componentProps: { placeholder: '请输入店铺LOGO', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'coverImage',
      label: '封面图',
      componentProps: { placeholder: '请输入封面图', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'contactName',
      label: '联系人姓名',
      componentProps: { placeholder: '请输入联系人姓名', maxlength: 50 },
    },
    {
      component: 'Input',
      fieldName: 'contactPhone',
      label: '联系电话',
      componentProps: { placeholder: '请输入联系电话', maxlength: 20 },
    },
    {
      component: 'Input',
      fieldName: 'intro',
      label: '店铺简介',
      componentProps: { placeholder: '请输入店铺简介', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'commissionRate',
      label: '店铺抽成比例（百分比，如 10 表示 10%）',
      componentProps: { placeholder: '请输入店铺抽成比例（百分比，如 10 表示 10%）' },
    },
    {
      component: 'InputNumber',
      fieldName: 'coachNum',
      label: '陪玩师数量',
      componentProps: { placeholder: '请输入陪玩师数量', class: 'w-full' },
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
        await updateShop({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createShop(values);
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
        modalApi.setState({ title: '编辑店铺表' });
        try {
          const detail = await getShopDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建店铺表' });
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
