<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
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
      label: 'åˆ†ç±»ID',
      rules: 'selectRequired',
      componentProps: { options: categoryIDOptions, placeholder: '请选择åˆ†ç±»ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'coachID',
      label: 'é™ªçŽ©å¸ˆID',
      rules: 'selectRequired',
      componentProps: { options: coachIDOptions, placeholder: '请选择é™ªçŽ©å¸ˆID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'title',
      label: 'å•†å“åç§°',
      rules: 'required',
      componentProps: { placeholder: '请输入å•†å“åç§°', maxlength: 100 },
    },
    {
      component: 'Input',
      fieldName: 'coverImage',
      label: 'å•†å“å°é¢å›¾',
      componentProps: { placeholder: '请输入å•†å“å°é¢å›¾', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'descContent',
      label: 'å•†å“è¯¦æƒ…æè¿°',
      componentProps: { placeholder: '请输入å•†å“è¯¦æƒ…æè¿°', maxlength: 65535 },
    },
    {
      component: 'Input',
      fieldName: 'price',
      label: 'å•ä»·ï¼ˆåˆ†ï¼‰',
      componentProps: { placeholder: '请输入å•ä»·ï¼ˆåˆ†ï¼‰' },
    },
    {
      component: 'Input',
      fieldName: 'unit',
      label: 'è®¡é‡å•ä½',
      componentProps: { placeholder: '请输入è®¡é‡å•ä½', maxlength: 20 },
    },
    {
      component: 'InputNumber',
      fieldName: 'salesNum',
      label: 'é”€é‡',
      componentProps: { placeholder: '请输入é”€é‡', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: 'æŽ’åºï¼ˆå‡åºï¼‰',
      componentProps: { placeholder: '请输入æŽ’åºï¼ˆå‡åºï¼‰', class: 'w-full' },
    },
    {
      component: 'Switch',
      fieldName: 'status',
      label: 'çŠ¶æ€',
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
        modalApi.setState({ title: '编辑å•†å“è¡¨' });
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
        modalApi.setState({ title: '新建å•†å“è¡¨' });
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
