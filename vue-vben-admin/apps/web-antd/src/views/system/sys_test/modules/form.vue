<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getSysTestDetail,
  createSysTest,
  updateSysTest,
  getSysTestTree,
} from '#/api/system/sys_test';
import type { SysTestItem } from '#/api/system/sys_test/types';

const treeData = ref<SysTestItem[]>([]);

/** ç±»åž‹选项 */
const typeOptions = [
  { label: 'æ™®é€š', value: 1 },
  { label: 'ç‰¹æ®Š', value: 2 },
  { label: 'é«˜çº§', value: 3 },
];

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
      label: 'ä¸Šçº§IDï¼Œ0è¡¨ç¤ºé¡¶çº§',
      componentProps: {
        treeData: treeData.value,
        fieldNames: { label: 'title', value: 'id', children: 'children' },
        placeholder: '请选择ä¸Šçº§IDï¼Œ0è¡¨ç¤ºé¡¶çº§',
        allowClear: true,
        treeDefaultExpandAll: true,
        class: 'w-full',
      },
    },
    {
      component: 'Input',
      fieldName: 'title',
      label: 'åç§°',
      componentProps: { placeholder: '请输入åç§°', maxlength: 100 },
    },
    {
      component: 'Input',
      fieldName: 'code',
      label: 'ç¼–ç ',
      componentProps: { placeholder: '请输入ç¼–ç ', maxlength: 50 },
    },
    {
      component: 'Select',
      fieldName: 'type',
      label: 'ç±»åž‹',
      componentProps: { options: typeOptions, placeholder: '请选择ç±»åž‹', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Switch',
      fieldName: 'status',
      label: 'çŠ¶æ€',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 1,
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: 'æŽ’åº',
      componentProps: { placeholder: '请输入æŽ’åº', class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'remark',
      label: 'å¤‡æ³¨',
      componentProps: { placeholder: '请输入å¤‡æ³¨', maxlength: 500 },
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
        await updateSysTest({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createSysTest(values);
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
        const res = await getSysTestTree();
        treeData.value = [
          { id: '0', title: '顶级节点', children: res ?? [] } as SysTestItem,
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
        modalApi.setState({ title: '编辑æµ‹è¯•è¡¨' });
        try {
          const detail = await getSysTestDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建æµ‹è¯•è¡¨' });
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
