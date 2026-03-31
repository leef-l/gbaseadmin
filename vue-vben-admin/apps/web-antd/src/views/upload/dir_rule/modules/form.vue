<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getDirRuleDetail,
  createDirRule,
  updateDirRule,
} from '#/api/upload/dir_rule';
import { getDirTree } from '#/api/upload/dir';
import type { DirItem } from '#/api/upload/dir/types';

/** ç±»åˆ«选项 */
const categoryOptions = [
  { label: 'é»˜è®¤', value: 1 },
  { label: 'ç±»åž‹', value: 2 },
  { label: 'æŽ¥å£', value: 3 },
];

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 目录下拉选项 */
const dirIDOptions = ref<{ label: string; value: string }[]>([]);

/** 将树形目录打平为选项 */
function flattenDirTree(
  items: DirItem[],
  prefix = '',
): { label: string; value: string }[] {
  const result: { label: string; value: string }[] = [];
  for (const item of items) {
    const label = prefix ? `${prefix} / ${item.name}` : item.name;
    result.push({ label, value: item.id });
    if (item.children?.length) {
      result.push(...flattenDirTree(item.children, label));
    }
  }
  return result;
}

/** 加载目录选项 */
async function loadDirOptions() {
  try {
    const list = await getDirTree();
    dirIDOptions.value = flattenDirTree(list);
  } catch {
    dirIDOptions.value = [];
  }
}

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Select',
      fieldName: 'dirID',
      label: 'ç›®å½•ID',
      rules: 'selectRequired',
      componentProps: { options: dirIDOptions, placeholder: '请选择ç›®å½•ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'category',
      label: 'ç±»åˆ«',
      componentProps: { options: categoryOptions, placeholder: '请选择ç±»åˆ«', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'savePath',
      label: 'ä¿å­˜ç›®å½•',
      componentProps: { placeholder: '请输入ä¿å­˜ç›®å½•', maxlength: 500 },
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
        await updateDirRule({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createDirRule(values);
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
      await loadDirOptions();
      const data = modalApi.getData<{ id?: string } | null>();
      if (data?.id) {
        isEdit.value = true;
        editId.value = data.id;
        modalApi.setState({ title: '编辑æ–‡ä»¶ç›®å½•è§„åˆ™' });
        try {
          const detail = await getDirRuleDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建æ–‡ä»¶ç›®å½•è§„åˆ™' });
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
