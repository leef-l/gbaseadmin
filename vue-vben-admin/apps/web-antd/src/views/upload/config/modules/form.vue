<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
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

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 根据存储类型显示/隐藏字段 */
function getStorageDependencies(storageVal: number) {
  return {
    localPath: storageVal === 1,
    ossEndpoint: storageVal === 2,
    ossBucket: storageVal === 2,
    ossAccessKey: storageVal === 2,
    ossSecretKey: storageVal === 2,
    cosRegion: storageVal === 3,
    cosBucket: storageVal === 3,
    cosSecretID: storageVal === 3,
    cosSecretKey: storageVal === 3,
  };
}

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
      component: 'Select',
      fieldName: 'storage',
      label: '存储类型',
      rules: 'selectRequired',
      componentProps: {
        options: storageOptions,
        placeholder: '请选择存储类型',
        allowClear: true,
        class: 'w-full',
        onChange: (val: number) => {
          const vis = getStorageDependencies(val);
          formApi.updateSchema(
            Object.entries(vis).map(([fieldName, show]) => ({
              fieldName,
              dependencies: { show },
            })),
          );
        },
      },
    },
    {
      component: 'Switch',
      fieldName: 'isDefault',
      label: '是否默认',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 0,
    },
    {
      component: 'Input',
      fieldName: 'localPath',
      label: '本地存储路径',
      dependencies: { show: false },
      componentProps: { placeholder: '请输入本地存储路径', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'ossEndpoint',
      label: 'OSS Endpoint',
      dependencies: { show: false },
      componentProps: { placeholder: '请输入OSS Endpoint', maxlength: 255 },
    },
    {
      component: 'Input',
      fieldName: 'ossBucket',
      label: 'OSS Bucket',
      dependencies: { show: false },
      componentProps: { placeholder: '请输入OSS Bucket', maxlength: 255 },
    },
    {
      component: 'Input',
      fieldName: 'ossAccessKey',
      label: 'OSS AccessKey',
      dependencies: { show: false },
      componentProps: { placeholder: '请输入OSS AccessKey', maxlength: 255 },
    },
    {
      component: 'Input',
      fieldName: 'ossSecretKey',
      label: 'OSS SecretKey',
      dependencies: { show: false },
      componentProps: { placeholder: '请输入OSS SecretKey', maxlength: 255 },
    },
    {
      component: 'Input',
      fieldName: 'cosRegion',
      label: 'COS Region',
      dependencies: { show: false },
      componentProps: { placeholder: '请输入COS Region', maxlength: 100 },
    },
    {
      component: 'Input',
      fieldName: 'cosBucket',
      label: 'COS Bucket',
      dependencies: { show: false },
      componentProps: { placeholder: '请输入COS Bucket', maxlength: 255 },
    },
    {
      component: 'Input',
      fieldName: 'cosSecretID',
      label: 'COS SecretId',
      dependencies: { show: false },
      componentProps: { placeholder: '请输入COS SecretId', maxlength: 255 },
    },
    {
      component: 'Input',
      fieldName: 'cosSecretKey',
      label: 'COS SecretKey',
      dependencies: { show: false },
      componentProps: { placeholder: '请输入COS SecretKey', maxlength: 255 },
    },
    {
      component: 'InputNumber',
      fieldName: 'maxSize',
      label: () => h('span', {}, ['最大文件大小 ', h(Tooltip, { title: 'MB' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入最大文件大小(MB)', min: 1, class: 'w-full' },
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
            // 根据存储类型显示对应字段
            if (detail.storage) {
              const vis = getStorageDependencies(detail.storage);
              formApi.updateSchema(
                Object.entries(vis).map(([fieldName, show]) => ({
                  fieldName,
                  dependencies: { show },
                })),
              );
            }
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
