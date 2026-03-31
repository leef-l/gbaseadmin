<script setup lang="ts">
import type { UploadChangeParam, UploadFile, UploadProps } from 'ant-design-vue';

import { computed, ref, watch } from 'vue';

import { Upload, message } from 'ant-design-vue';

import { uploadFile } from '#/api/upload/uploader';

interface Props {
  value?: string;
  maxCount?: number;
  maxSize?: number;
  accept?: string;
  disabled?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  value: '',
  maxCount: 1,
  maxSize: 10,
  accept: '',
  disabled: false,
});

const emit = defineEmits<{
  'update:value': [val: string];
}>();

const fileList = ref<UploadFile[]>([]);

/** Convert URL string(s) to UploadFile array */
function urlsToFileList(val: string): UploadFile[] {
  if (!val) return [];
  return val
    .split(',')
    .filter(Boolean)
    .map((url, index) => ({
      uid: `${index}-${Date.now()}`,
      name: url.split('/').pop() || 'file',
      status: 'done' as const,
      url,
    }));
}

/** Convert UploadFile array to comma-separated URL string */
function fileListToUrls(list: UploadFile[]): string {
  return list
    .filter((f) => f.status === 'done' && f.url)
    .map((f) => f.url)
    .join(',');
}
// Sync from prop to internal fileList
watch(
  () => props.value,
  (val) => {
    const newUrls = val || '';
    const currentUrls = fileListToUrls(fileList.value);
    if (newUrls !== currentUrls) {
      fileList.value = urlsToFileList(newUrls);
    }
  },
  { immediate: true },
);

const customRequest: UploadProps['customRequest'] = async (options) => {
  const { file, onSuccess, onError } = options;
  try {
    const res = await uploadFile(file as File);
    const uploadedFile = fileList.value.find(
      (f) => (f.originFileObj as File) === file,
    );
    if (uploadedFile) {
      uploadedFile.status = 'done';
      uploadedFile.url = res.url;
    }
    onSuccess?.(res);
    emit('update:value', fileListToUrls(fileList.value));
  } catch (error: any) {
    const failedFile = fileList.value.find(
      (f) => (f.originFileObj as File) === file,
    );
    if (failedFile) {
      failedFile.status = 'error';
    }
    onError?.(error);
  }
};

function beforeUpload(file: File) {
  const isLtMax = file.size / 1024 / 1024 < props.maxSize;
  if (!isLtMax) {
    message.error(`文件大小不能超过 ${props.maxSize}MB`);
  }
  return isLtMax;
}

function handleChange(info: UploadChangeParam) {
  fileList.value = info.fileList;
}

function handleRemove(file: UploadFile) {
  fileList.value = fileList.value.filter((f) => f.uid !== file.uid);
  emit('update:value', fileListToUrls(fileList.value));
  return true;
}
const showUploadButton = computed(
  () => !props.disabled && fileList.value.length < props.maxCount,
);
</script>

<template>
  <Upload
    v-model:file-list="fileList"
    :accept="accept"
    :before-upload="beforeUpload"
    :custom-request="customRequest"
    :disabled="disabled"
    :max-count="maxCount"
    list-type="text"
    @change="handleChange"
    @remove="handleRemove"
  >
    <a-button v-if="showUploadButton" :disabled="disabled">
      选择文件
    </a-button>
  </Upload>
</template>
