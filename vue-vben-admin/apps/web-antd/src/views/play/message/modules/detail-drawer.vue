<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { Descriptions, DescriptionsItem, Tag } from 'ant-design-vue';
import { getMessageDetail } from '#/api/play/message';
import type { MessageItem } from '#/api/play/message/types';

const detail = ref<MessageItem | null>(null);

const [Modal, modalApi] = useVbenModal({
  fullscreenButton: false,
  footer: false,
  async onOpenChange(isOpen: boolean) {
    if (isOpen) {
      const data = modalApi.getData<{ id: string }>();
      if (data?.id) {
        modalApi.setState({ title: '会员消息详情' });
        try {
          detail.value = await getMessageDetail(data.id);
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
      <DescriptionsItem label="接收者会员ID">{{ detail.memberNickname || '-' }}</DescriptionsItem>
      <DescriptionsItem label="消息标题">{{ detail.title || '-' }}</DescriptionsItem>
      <DescriptionsItem label="消息内容">{{ detail.content || '-' }}</DescriptionsItem>
      <DescriptionsItem label="消息类型 1=系统通知 2=订单消息 3=活动消息">{{ detail.msgType || '-' }}</DescriptionsItem>
      <DescriptionsItem label="关联业务ID">{{ detail.bizID || '-' }}</DescriptionsItem>
      <DescriptionsItem label="是否已读 0=未读 1=已读">{{ detail.isRead || '-' }}</DescriptionsItem>
      <DescriptionsItem label="状态 1=正常 0=禁用">{{ detail.status || '-' }}</DescriptionsItem>
      <DescriptionsItem label="创建时间">{{ detail.createdAt || '-' }}</DescriptionsItem>
      <DescriptionsItem label="更新时间">{{ detail.updatedAt || '-' }}</DescriptionsItem>
    </Descriptions>
  </Modal>
</template>
