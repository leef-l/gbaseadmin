<script setup lang="ts">
import { ref, reactive } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { Tabs, Table, Button, Modal, Form, Input, InputNumber, Select, message, Popconfirm } from 'ant-design-vue';
import { getActivityRewardList, createActivityReward, updateActivityReward, deleteActivityReward } from '#/api/play/activity_reward';
import { getActivityStepList, createActivityStep, updateActivityStep, deleteActivityStep } from '#/api/play/activity_step';
import type { ActivityRewardItem } from '#/api/play/activity_reward/types';
import type { ActivityStepItem } from '#/api/play/activity_step/types';

const rewards = ref<ActivityRewardItem[]>([]);
const steps = ref<ActivityStepItem[]>([]);
const loading = ref(false);
const activityId = ref('');
const activityTitle = ref('');
const activityType = ref(0);
const activeTab = ref('rewards');

/** 奖励编辑状态 */
const rewardModalVisible = ref(false);
const rewardForm = reactive({ id: '', rewardName: '', rewardType: 1, rewardValue: '', sort: 0 });
const rewardSaving = ref(false);

/** 步骤编辑状态 */
const stepModalVisible = ref(false);
const stepForm = reactive({ id: '', title: '', stepNum: 1, stepType: 1, exampleText: '', descContent: '', stepImage: '', sort: 0 });
const stepSaving = ref(false);

const stepTypeMap: Record<number, string> = { 1: '文字', 2: '链接', 3: '图片' };
const stepTypeOptions = [
  { label: '文字', value: 1 },
  { label: '链接', value: 2 },
  { label: '图片', value: 3 },
];

const rewardTypeMap: Record<number, string> = { 1: '余额', 2: '优惠券', 3: '经验值', 4: '会员等级天数' };
const rewardTypeOptions = [
  { label: '余额', value: 1 },
  { label: '优惠券', value: 2 },
  { label: '经验值', value: 3 },
  { label: '会员等级天数', value: 4 },
];

const rewardColumns = [
  { title: '奖励名称', dataIndex: 'rewardName', key: 'rewardName' },
  { title: '奖励类型', dataIndex: 'rewardType', key: 'rewardType', customRender: ({ text }: any) => rewardTypeMap[text] || text },
  { title: '奖励数值', dataIndex: 'rewardValue', key: 'rewardValue' },
  { title: '排序', dataIndex: 'sort', key: 'sort', width: 80 },
  { title: '操作', key: 'action', width: 150 },
];

const stepColumns = [
  { title: '步骤序号', dataIndex: 'stepNum', key: 'stepNum', width: 80 },
  { title: '步骤标题', dataIndex: 'title', key: 'title' },
  { title: '步骤类型', dataIndex: 'stepType', key: 'stepType', width: 80, customRender: ({ text }: any) => stepTypeMap[text] || '文字' },
  { title: '示例内容', dataIndex: 'exampleText', key: 'exampleText', ellipsis: true },
  { title: '步骤说明', dataIndex: 'descContent', key: 'descContent', ellipsis: true },
  { title: '排序', dataIndex: 'sort', key: 'sort', width: 80 },
  { title: '操作', key: 'action', width: 150 },
];
async function loadData(id: string) {
  loading.value = true;
  try {
    const [rewardRes, stepRes] = await Promise.all([
      getActivityRewardList({ pageNum: 1, pageSize: 100, activityID: id }),
      getActivityStepList({ pageNum: 1, pageSize: 100, activityID: id }),
    ]);
    rewards.value = rewardRes?.list ?? [];
    steps.value = stepRes?.list ?? [];
  } finally {
    loading.value = false;
  }
}

/** 奖励 CRUD */
function handleAddReward() {
  Object.assign(rewardForm, { id: '', rewardName: '', rewardType: 1, rewardValue: '', sort: 0 });
  rewardModalVisible.value = true;
}
function handleEditReward(row: ActivityRewardItem) {
  Object.assign(rewardForm, { id: row.id, rewardName: row.rewardName, rewardType: row.rewardType ?? 1, rewardValue: row.rewardValue ?? '', sort: row.sort ?? 0 });
  rewardModalVisible.value = true;
}
async function handleSaveReward() {
  if (!rewardForm.rewardName.trim()) { message.warning('请输入奖励名称'); return; }
  rewardSaving.value = true;
  try {
    const payload = { activityID: activityId.value, rewardName: rewardForm.rewardName, rewardType: rewardForm.rewardType, rewardValue: rewardForm.rewardValue, sort: rewardForm.sort };
    if (rewardForm.id) {
      await updateActivityReward({ id: rewardForm.id, ...payload });
    } else {
      await createActivityReward(payload);
    }
    message.success('保存成功');
    rewardModalVisible.value = false;
    loadData(activityId.value);
  } catch { message.error('保存失败'); } finally { rewardSaving.value = false; }
}
async function handleDeleteReward(row: ActivityRewardItem) {
  await deleteActivityReward(row.id);
  message.success('删除成功');
  loadData(activityId.value);
}

/** 步骤 CRUD */
function handleAddStep() {
  Object.assign(stepForm, { id: '', title: '', stepNum: steps.value.length + 1, stepType: 1, exampleText: '', descContent: '', stepImage: '', sort: 0 });
  stepModalVisible.value = true;
}
function handleEditStep(row: ActivityStepItem) {
  Object.assign(stepForm, { id: row.id, title: row.title, stepNum: row.stepNum ?? 1, stepType: row.stepType ?? 1, exampleText: row.exampleText ?? '', descContent: row.descContent ?? '', stepImage: row.stepImage ?? '', sort: row.sort ?? 0 });
  stepModalVisible.value = true;
}
async function handleSaveStep() {
  if (!stepForm.title.trim()) { message.warning('请输入步骤标题'); return; }
  stepSaving.value = true;
  try {
    const payload = { activityID: activityId.value, title: stepForm.title, stepNum: stepForm.stepNum, stepType: stepForm.stepType, exampleText: stepForm.exampleText, descContent: stepForm.descContent, stepImage: stepForm.stepImage, sort: stepForm.sort };
    if (stepForm.id) {
      await updateActivityStep({ id: stepForm.id, ...payload });
    } else {
      await createActivityStep(payload);
    }
    message.success('保存成功');
    stepModalVisible.value = false;
    loadData(activityId.value);
  } catch { message.error('保存失败'); } finally { stepSaving.value = false; }
}
async function handleDeleteStep(row: ActivityStepItem) {
  await deleteActivityStep(row.id);
  message.success('删除成功');
  loadData(activityId.value);
}

const [DrawerModal, modalApi] = useVbenModal({
  onOpenChange(isOpen: boolean) {
    if (isOpen) {
      const data = modalApi.getData<{ id: string; title: string; type: number; tab: string }>();
      if (data) {
        activityId.value = data.id;
        activityTitle.value = data.title;
        activityType.value = data.type;
        activeTab.value = data.tab || 'rewards';
        loadData(data.id);
      }
    }
  },
});
</script>

<template>
  <DrawerModal :title="`活动管理 - ${activityTitle}`" :footer="null">
    <Tabs v-model:activeKey="activeTab">
      <Tabs.TabPane key="rewards" tab="奖励管理">
        <div style="margin-bottom: 12px">
          <Button type="primary" size="small" @click="handleAddReward">新增奖励</Button>
        </div>
        <Table :columns="rewardColumns" :data-source="rewards" :loading="loading" :pagination="false" row-key="id" size="small">
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'action'">
              <Button type="link" size="small" @click="handleEditReward(record)">编辑</Button>
              <Popconfirm title="确定删除？" @confirm="handleDeleteReward(record)">
                <Button type="link" danger size="small">删除</Button>
              </Popconfirm>
            </template>
          </template>
        </Table>
      </Tabs.TabPane>
      <Tabs.TabPane v-if="activityType === 4" key="steps" tab="步骤管理">
        <div style="margin-bottom: 12px">
          <Button type="primary" size="small" @click="handleAddStep">新增步骤</Button>
        </div>
        <Table :columns="stepColumns" :data-source="steps" :loading="loading" :pagination="false" row-key="id" size="small">
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'action'">
              <Button type="link" size="small" @click="handleEditStep(record)">编辑</Button>
              <Popconfirm title="确定删除？" @confirm="handleDeleteStep(record)">
                <Button type="link" danger size="small">删除</Button>
              </Popconfirm>
            </template>
          </template>
        </Table>
      </Tabs.TabPane>
    </Tabs>

    <!-- 奖励编辑弹窗 -->
    <Modal v-model:open="rewardModalVisible" :title="rewardForm.id ? '编辑奖励' : '新增奖励'" :confirm-loading="rewardSaving" @ok="handleSaveReward">
      <Form layout="vertical">
        <Form.Item label="奖励名称" required>
          <Input v-model:value="rewardForm.rewardName" placeholder="请输入奖励名称" />
        </Form.Item>
        <Form.Item label="奖励类型" required>
          <Select v-model:value="rewardForm.rewardType" :options="rewardTypeOptions" placeholder="请选择奖励类型" style="width: 100%" />
        </Form.Item>
        <Form.Item label="奖励数值">
          <Input v-model:value="rewardForm.rewardValue" placeholder="请输入奖励数值" />
        </Form.Item>
        <Form.Item label="排序">
          <InputNumber v-model:value="rewardForm.sort" :min="0" style="width: 100%" />
        </Form.Item>
      </Form>
    </Modal>
    <!-- 步骤编辑弹窗 -->
    <Modal v-model:open="stepModalVisible" :title="stepForm.id ? '编辑步骤' : '新增步骤'" :confirm-loading="stepSaving" @ok="handleSaveStep">
      <Form layout="vertical">
        <Form.Item label="步骤标题" required>
          <Input v-model:value="stepForm.title" placeholder="请输入步骤标题" />
        </Form.Item>
        <Form.Item label="步骤序号">
          <InputNumber v-model:value="stepForm.stepNum" :min="1" style="width: 100%" />
        </Form.Item>
        <Form.Item label="步骤类型" required>
          <Select v-model:value="stepForm.stepType" :options="stepTypeOptions" placeholder="请选择步骤类型" />
        </Form.Item>
        <Form.Item v-if="stepForm.stepType === 1" label="示例文字">
          <Input.TextArea v-model:value="stepForm.exampleText" placeholder="用户可复制的示例文字" :rows="3" />
        </Form.Item>
        <Form.Item v-if="stepForm.stepType === 2" label="链接URL">
          <Input v-model:value="stepForm.exampleText" placeholder="用户可复制并打开的链接地址" />
        </Form.Item>
        <Form.Item v-if="stepForm.stepType === 3" label="示例图片">
          <Input v-model:value="stepForm.stepImage" placeholder="请输入示例图片URL" />
        </Form.Item>
        <Form.Item label="步骤说明">
          <Input.TextArea v-model:value="stepForm.descContent" placeholder="请输入步骤说明" :rows="3" />
        </Form.Item>
        <Form.Item label="排序">
          <InputNumber v-model:value="stepForm.sort" :min="0" style="width: 100%" />
        </Form.Item>
      </Form>
    </Modal>
  </DrawerModal>
</template>
