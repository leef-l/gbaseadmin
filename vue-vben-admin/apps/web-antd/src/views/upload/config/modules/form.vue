<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message } from 'ant-design-vue';
import {
  getConfigDetail,
  createConfig,
  updateConfig,
} from '#/api/upload/config';

/** 存储类型选项 */
const storageOptions = [
  { label: '本地', value: 1 },
  { label: '阿里云OSS', value: 2 },
  { label: '腾讯云COS', value: 3 },
];

/** 存储类型依赖: 控制字段显隐 */
const localDeps = {
  triggerFields: ['storage'],
  if: (values: Record<string, any>) => values.storage === 1,
};
const ossDeps = {
  triggerFields: ['storage'],
  if: (values: Record<string, any>) => values.storage === 2,
};
const cosDeps = {
  triggerFields: ['storage'],
  if: (values: Record<string, any>) => values.storage === 3,
};

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: '配置名称',
      rules: 'required',
      componentProps: { placeholder: '请输入配置名称', maxlength: 100 },
    },
    {
      component: 'RadioGroup',
      fieldName: 'storage',
      label: '存储类型',
      rules: 'required',
      defaultValue: 1,
      componentProps: { options: storageOptions, optionType: 'button', buttonStyle: 'solid' },
    },
    {
      component: 'Switch',
      fieldName: 'isDefault',
      label: '是否默认',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 0,
    },
    // ---- 本地存储字段 ----
    {
      component: 'Input',
      fieldName: 'localPath',
      label: '本地存储路径',
      componentProps: { placeholder: '请输入本地存储路径', maxlength: 500 },
      dependencies: localDeps,
    },
    // ---- 阿里云 OSS 字段 ----
    {
      component: 'Input',
      fieldName: 'ossEndpoint',
      label: 'OSS Endpoint',
      componentProps: { placeholder: '请输入OSS Endpoint', maxlength: 255 },
      dependencies: ossDeps,
    },
    {
      component: 'Input',
      fieldName: 'ossBucket',
      label: 'OSS Bucket',
      componentProps: { placeholder: '请输入OSS Bucket', maxlength: 255 },
      dependencies: ossDeps,
    },
    {
      component: 'Input',
      fieldName: 'ossAccessKey',
      label: 'OSS AccessKey',
      componentProps: { placeholder: '请输入OSS AccessKey', maxlength: 255 },
      dependencies: ossDeps,
    },
    {
      component: 'Input',
      fieldName: 'ossSecretKey',
      label: 'OSS SecretKey',
      componentProps: { placeholder: '请输入OSS SecretKey', maxlength: 255 },
      dependencies: ossDeps,
    },
    // ---- 腾讯云 COS 字段 ----
    {
      component: 'Input',
      fieldName: 'cosRegion',
      label: 'COS Region',
      componentProps: { placeholder: '请输入COS Region', maxlength: 100 },
      dependencies: cosDeps,
    },
    {
      component: 'Input',
      fieldName: 'cosBucket',
      label: 'COS Bucket',
      componentProps: { placeholder: '请输入COS Bucket', maxlength: 255 },
      dependencies: cosDeps,
    },
    {
      component: 'Input',
      fieldName: 'cosSecretID',
      label: 'COS SecretId',
      componentProps: { placeholder: '请输入COS SecretId', maxlength: 255 },
      dependencies: cosDeps,
    },
    {
      component: 'Input',
      fieldName: 'cosSecretKey',
      label: 'COS SecretKey',
      componentProps: { placeholder: '请输入COS SecretKey', maxlength: 255 },
      dependencies: cosDeps,
    },
    // ---- 通用字段 ----
    {
      component: 'InputNumber',
      fieldName: 'maxSize',
      label: '最大文件大小(MB)',
      componentProps: { placeholder: '请输入最大文件大小(MB)', min: 0 },
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
        await updateConfig({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createConfig(values);
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
        modalApi.setState({ title: '编辑上传配置' });
        try {
          const detail = await getConfigDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建上传配置' });
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
