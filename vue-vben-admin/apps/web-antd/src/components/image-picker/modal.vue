<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import ImagePicker from './index.vue';

const pickerRef = ref<InstanceType<typeof ImagePicker>>();
const multiple = ref(false);
const maxCount = ref(1);

const emit = defineEmits<{
  confirm: [images: { id: string; url: string; name: string; size: number; ext: string }[]];
}>();

const [Modal, modalApi] = useVbenModal({
  fullscreenButton: false,
  onCancel() {
    modalApi.close();
  },
  onOpenChange(isOpen: boolean) {
    if (isOpen) {
      const data = modalApi.getData<{ multiple?: boolean; maxCount?: number } | null>();
      multiple.value = data?.multiple ?? false;
      maxCount.value = data?.maxCount ?? 1;
      pickerRef.value?.reset();
      pickerRef.value?.loadFileList();
    }
  },
});

function handleConfirm(images: { id: string; url: string; name: string; size: number; ext: string }[]) {
  emit('confirm', images);
  modalApi.close();
}
</script>

<template>
  <Modal title="选择图片" class="w-[900px]">
    <ImagePicker
      ref="pickerRef"
      :multiple="multiple"
      :max-count="maxCount"
      @confirm="handleConfirm"
    />
  </Modal>
</template>
