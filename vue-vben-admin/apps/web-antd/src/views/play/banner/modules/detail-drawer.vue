<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { Descriptions, DescriptionsItem, Tag } from 'ant-design-vue';
import { getBannerDetail } from '#/api/play/banner';
import type { BannerItem } from '#/api/play/banner/types';

const detail = ref<BannerItem | null>(null);

const [Modal, modalApi] = useVbenModal({
  fullscreenButton: false,
  footer: false,
  async onOpenChange(isOpen: boolean) {
    if (isOpen) {
      const data = modalApi.getData<{ id: string }>();
      if (data?.id) {
        modalApi.setState({ title: '首页Banner轮播详情' });
        try {
          detail.value = await getBannerDetail(data.id);
        } catch {
          detail.value = null;
        }
      }
    } else {
      detail.value = null;
    }
  },
});
</script>

<template>
  <Modal class="w-[600px]">
    <Descriptions v-if="detail" bordered :column="1" size="small">
      <DescriptionsItem label="ID">{{ detail.id }}</DescriptionsItem>
      <DescriptionsItem label="Banner标题">{{ detail.title || '-' }}</DescriptionsItem>
      <DescriptionsItem label="图片URL">{{ detail.image || '-' }}</DescriptionsItem>
      <DescriptionsItem label="跳转类型">{{ detail.linkType || '-' }}</DescriptionsItem>
      <DescriptionsItem label="跳转值">{{ detail.linkValue || '-' }}</DescriptionsItem>
      <DescriptionsItem label="排序">{{ detail.sort || '-' }}</DescriptionsItem>
      <DescriptionsItem label="状态">{{ detail.status || '-' }}</DescriptionsItem>
      <DescriptionsItem label="生效开始时间">{{ detail.startTime || '-' }}</DescriptionsItem>
      <DescriptionsItem label="生效结束时间">{{ detail.endTime || '-' }}</DescriptionsItem>
      <DescriptionsItem label="备注">{{ detail.remark || '-' }}</DescriptionsItem>
      <DescriptionsItem label="创建时间">{{ detail.createdAt || '-' }}</DescriptionsItem>
    </Descriptions>
  </Modal>
</template>
