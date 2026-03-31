import { useVbenModal } from '@vben/common-ui';
import ImagePickerModal from './modal.vue';

export function useImagePicker() {
  const [ModalComp, modalApi] = useVbenModal({
    connectedComponent: ImagePickerModal,
  });

  function open(options?: { multiple?: boolean; maxCount?: number }) {
    modalApi.setData(options || {});
    modalApi.open();
  }

  return { ModalComp, modalApi, open };
}
