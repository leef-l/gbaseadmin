<script setup lang="ts">
import { ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { Tabs, Table, Tag } from 'ant-design-vue';
import { getActivityRewardList } from '#/api/play/activity_reward';
import { getActivityStepList } from '#/api/play/activity_step';
import type { ActivityRewardItem } from '#/api/play/activity_reward/types';
import type { ActivityStepItem } from '#/api/play/activity_step/types';

const rewards = ref<ActivityRewardItem[]>([]);
const steps = ref<ActivityStepItem[]>([]);
const loading = ref(false);
const activityTitle = ref('');

const rewardTypeMap: Record<number, string> = {
  1: '余额', 2: '优惠券', 3: '经验', 4: '等级天数',
};
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple'];

const rewardColumns = [
  { title: '奖励名称', dataIndex: 'rewardName', key: 'rewardName' },
  { title: '奖励类型', dataIndex: 'rewardType', key: 'rewardType', customRender: ({ text }: any) => rewardTypeMap[text] || text },
  { title: '奖励数值', dataIndex: 'rewardValue', key: 'rewardValue' },
  { title: '排序', dataIndex: 'sort', key: 'sort' },
];

const stepColumns = [
  { title: '步骤序号', dataIndex: 'stepNum', key: 'stepNum' },
  { title: '步骤标题', dataIndex: 'title', key: 'title' },
  { title: '步骤说明', dataIndex: 'descContent', key: 'descContent', ellipsis: true },
  { title: '排序', dataIndex: 'sort', key: 'sort' },
];

async function loadData(activityID: string, title: string) {
  loading.value = true;
  activityTitle.value = title;
  try {
    const [rewardRes, stepRes] = await Promise.all([
      getActivityRewardList({ pageNum: 1, pageSize: 100, activityID }),
      getActivityStepList({ pageNum: 1, pageSize: 100, activityID }),
    ]);
    rewards.value = rewardRes?.list ?? [];
    steps.value = stepRes?.list ?? [];
  } finally {
    loading.value = false;
  }
}

const [Modal, modalApi] = useVbenModal({
  onOpenChange(isOpen: boolean) {
    if (isOpen) {
      const data = modalApi.getData<{ id: string; title: string }>();
      if (data) {
        loadData(data.id, data.title);
      }
    }
  },
});
</script>

<template>
  <Modal :title="`活动详情 - ${activityTitle}`" :footer="null">
    <Tabs default-active-key="rewards">
      <Tabs.TabPane key="rewards" tab="奖励配置">
        <Table
          :columns="rewardColumns"
          :data-source="rewards"
          :loading="loading"
          :pagination="false"
          row-key="id"
          size="small"
        />
      </Tabs.TabPane>
      <Tabs.TabPane key="steps" tab="步骤配置">
        <Table
          :columns="stepColumns"
          :data-source="steps"
          :loading="loading"
          :pagination="false"
          row-key="id"
          size="small"
        />
      </Tabs.TabPane>
    </Tabs>
  </Modal>
</template>
