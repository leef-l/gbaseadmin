<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getDirDetail,
  createDir,
  updateDir,
  getDirTree,
} from '#/api/upload/dir';
import type { DirItem } from '#/api/upload/dir/types';

const treeData = ref<DirItem[]>([]);

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'TreeSelect',
      fieldName: 'parentID',
      label: 'ä¸Šçº§ç›®å½•',
      componentProps: {
        treeData: treeData.value,
        fieldNames: { label: 'name', value: 'id', children: 'children' },
        placeholder: '请选择ä¸Šçº§ç›®å½•',
        allowClear: true,
        treeDefaultExpandAll: true,
        class: 'w-full',
      },
    },
    {
      component: 'Input',
      fieldName: 'name',
      label: 'ç›®å½•åç§°',
      rules: 'required',
      componentProps: { placeholder: '请输入ç›®å½•åç§°', maxlength: 100 },
    },
    {
      component: 'Input',
      fieldName: 'path',
      label: 'ç›®å½•è·¯å¾„',
      rules: 'required',
      componentProps: { placeholder: '请输入ç›®å½•è·¯å¾„', maxlength: 500 },
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: 'æŽ’åº',
      componentProps: { placeholder: '请输入æŽ’åº', class: 'w-full' },
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
        await updateDir({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createDir(values);
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
      // 加载树形数据
      try {
        const res = await getDirTree();
        treeData.value = [
          { id: '0', title: '顶级节点', children: res ?? [] } as DirItem,
        ];
        formApi.updateSchema([
          {
            fieldName: 'parentID',
            componentProps: { treeData: treeData.value },
          },
        ]);
      } catch {
        // ignore
      }
      if (data?.id) {
        isEdit.value = true;
        editId.value = data.id;
        modalApi.setState({ title: '编辑æ–‡ä»¶ç›®å½•' });
        try {
          const detail = await getDirDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建æ–‡ä»¶ç›®å½•' });
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
