<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { Descriptions, DescriptionsItem, Tag } from 'ant-design-vue';
import { getWithdrawDetail } from '#/api/play/withdraw';
import type { WithdrawItem } from '#/api/play/withdraw/types';

const detail = ref<WithdrawItem | null>(null);

const [Modal, modalApi] = useVbenModal({
  fullscreenButton: false,
  footer: false,
  async onOpenChange(isOpen: boolean) {
    if (isOpen) {
      const data = modalApi.getData<{ id: string }>();
      if (data?.id) {
        modalApi.setState({ title: '陪玩师提现记录详情' });
        try {
          detail.value = await getWithdrawDetail(data.id);
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
      <DescriptionsItem label="陪玩师ID">{{ detail.coachRealName || '-' }}</DescriptionsItem>
      <DescriptionsItem label="会员ID">{{ detail.memberNickname || '-' }}</DescriptionsItem>
      <DescriptionsItem label="提现金额">{{ detail.amount != null ? (detail.amount / 100).toFixed(2) : '-' }}</DescriptionsItem>
      <DescriptionsItem label="状态 0=待审核 1=已打款 2=已拒绝">{{ detail.status || '-' }}</DescriptionsItem>
      <DescriptionsItem label="拒绝原因">{{ detail.reason || '-' }}</DescriptionsItem>
      <DescriptionsItem label="审核时间">{{ detail.auditedAt || '-' }}</DescriptionsItem>
      <DescriptionsItem label="审核时间">{{ detail.auditedAt || '-' }}</DescriptionsItem>
      <DescriptionsItem label="创建时间">{{ detail.createdAt || '-' }}</DescriptionsItem>
      <DescriptionsItem label="更新时间">{{ detail.updatedAt || '-' }}</DescriptionsItem>
    </Descriptions>
  </Modal>
</template>
