<script setup lang="ts">
import { ref, reactive, computed } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { Tabs, Table, Button, Modal, Form, Input, InputNumber, Select, message, Popconfirm, Tag } from 'ant-design-vue';
import { getActivityRewardList, createActivityReward, updateActivityReward, deleteActivityReward } from '#/api/play/activity_reward';
import { getActivityStepList, createActivityStep, updateActivityStep, deleteActivityStep } from '#/api/play/activity_step';
import { getCouponList } from '#/api/play/coupon';
import { getMemberLevelList } from '#/api/play/member_level';
import type { ActivityRewardItem } from '#/api/play/activity_reward/types';
import type { ActivityStepItem } from '#/api/play/activity_step/types';
import ImageUpload from '#/components/upload/image-upload.vue';

const rewards = ref<ActivityRewardItem[]>([]);
const steps = ref<ActivityStepItem[]>([]);
const loading = ref(false);
const activityId = ref('');
const activityTitle = ref('');
const activityType = ref(0);
const activeTab = ref('rewards');

/** 奖励编辑状态 */
const rewardModalVisible = ref(false);
const rewardForm = reactive({ id: '', rewardName: '', rewardType: 1, rewardValue: '', rewardLevelId: '', rewardDays: 1, sort: 0 });
const rewardSaving = ref(false);

/** 优惠券下拉选项 */
const couponOptions = ref<{ label: string; value: string }[]>([]);
const couponSearching = ref(false);
async function handleCouponSearch(keyword: string) {
  couponSearching.value = true;
  try {
    const res = await getCouponList({ pageNum: 1, pageSize: 50, title: keyword } as any);
    couponOptions.value = (res?.list ?? []).map((c: any) => ({ label: `${c.title}（面值${c.faceValue}分）`, value: String(c.id) }));
  } finally { couponSearching.value = false; }
}

/** 会员等级下拉选项 */
const levelOptions = ref<{ label: string; value: string }[]>([]);
async function loadLevelOptions() {
  const res = await getMemberLevelList({ pageNum: 1, pageSize: 50 } as any);
  levelOptions.value = (res?.list ?? []).map((l: any) => ({ label: l.title, value: String(l.id) }));
}

/** 奖励数值的计算（会员等级天数时 = levelId:days） */
const rewardValueLabel = computed(() => {
  if (rewardForm.rewardType === 1) return '余额（单位：分）';
  if (rewardForm.rewardType === 2) return '选择优惠券';
  if (rewardForm.rewardType === 3) return '经验值';
  if (rewardForm.rewardType === 4) return '会员等级天数';
  return '奖励数值';
});

/** 步骤编辑状态 */
const stepModalVisible = ref(false);
const stepForm = reactive({ id: '', title: '', stepNum: 1, stepType: 1, exampleText: '', descContent: '', stepImage: '', isRequired: 0, sort: 0 });
const stepSaving = ref(false);

const stepTypeMap: Record<number, string> = { 1: '文字', 2: '链接', 3: '图片' };
const stepTypeColor: Record<number, string> = { 1: 'blue', 2: 'green', 3: 'orange' };
const stepTypeOptions = [
  { label: '文字（用户需复制示例文字）', value: 1 },
  { label: '链接（用户需前往目标链接）', value: 2 },
  { label: '图片（用户需上传截图）', value: 3 },
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
  { title: '奖励类型', dataIndex: 'rewardType', key: 'rewardType', width: 120, customRender: ({ text }: any) => rewardTypeMap[text] || text },
  { title: '奖励数值', dataIndex: 'rewardValue', key: 'rewardValue', width: 100 },
  { title: '排序', dataIndex: 'sort', key: 'sort', width: 70 },
  { title: '操作', key: 'action', width: 120 },
];

const stepColumns = [
  { title: '序号', dataIndex: 'stepNum', key: 'stepNum', width: 60 },
  { title: '步骤标题', dataIndex: 'title', key: 'title' },
  { title: '类型', dataIndex: 'stepType', key: 'stepType', width: 80 },
  { title: '是否填写', dataIndex: 'isRequired', key: 'isRequired', width: 100 },
  { title: '示例内容', dataIndex: 'exampleText', key: 'exampleText', ellipsis: true },
  { title: '示例图片', dataIndex: 'stepImage', key: 'stepImage', width: 80 },
  { title: '步骤说明', dataIndex: 'descContent', key: 'descContent', ellipsis: true },
  { title: '排序', dataIndex: 'sort', key: 'sort', width: 70 },
  { title: '操作', key: 'action', width: 120 },
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
    // 加载会员等级选项
    loadLevelOptions();
  } finally {
    loading.value = false;
  }
}

/** 奖励 CRUD */
function handleAddReward() {
  Object.assign(rewardForm, { id: '', rewardName: '', rewardType: 1, rewardValue: '', rewardLevelId: '', rewardDays: 1, sort: 0 });
  rewardModalVisible.value = true;
}
function handleEditReward(row: ActivityRewardItem) {
  let levelId = '', days = 1;
  if (row.rewardType === 4 && row.rewardValue) {
    const parts = String(row.rewardValue).split(':');
    levelId = parts[0] || '';
    days = parseInt(parts[1] || '1', 10);
  }
  Object.assign(rewardForm, {
    id: row.id,
    rewardName: row.rewardName,
    rewardType: row.rewardType ?? 1,
    rewardValue: row.rewardType === 4 ? '' : (row.rewardValue ?? ''),
    rewardLevelId: levelId,
    rewardDays: days,
    sort: row.sort ?? 0,
  });
  rewardModalVisible.value = true;
}
async function handleSaveReward() {
  if (!rewardForm.rewardName.trim()) { message.warning('请输入奖励名称'); return; }
  let finalValue = rewardForm.rewardValue;
  if (rewardForm.rewardType === 4) {
    if (!rewardForm.rewardLevelId) { message.warning('请选择会员等级'); return; }
    finalValue = `${rewardForm.rewardLevelId}:${rewardForm.rewardDays}`;
  } else if (rewardForm.rewardType === 2) {
    if (!rewardForm.rewardValue) { message.warning('请选择优惠券'); return; }
  }
  rewardSaving.value = true;
  try {
    const payload = { activityID: activityId.value, rewardName: rewardForm.rewardName, rewardType: rewardForm.rewardType, rewardValue: finalValue, sort: rewardForm.sort };
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
  Object.assign(stepForm, { id: '', title: '', stepNum: steps.value.length + 1, stepType: 1, exampleText: '', descContent: '', stepImage: '', sort: steps.value.length });
  stepModalVisible.value = true;
}
function handleEditStep(row: ActivityStepItem) {
  Object.assign(stepForm, {
    id: row.id,
    title: row.title,
    stepNum: row.stepNum ?? 1,
    stepType: row.stepType ?? 1,
    exampleText: row.exampleText ?? '',
    descContent: row.descContent ?? '',
    stepImage: row.stepImage ?? '',
    isRequired: row.isRequired ?? 0,
    sort: row.sort ?? 0,
  });
  stepModalVisible.value = true;
}
async function handleSaveStep() {
  if (!stepForm.title.trim()) { message.warning('请输入步骤标题'); return; }
  if (stepForm.stepType === 3 && !stepForm.stepImage) { message.warning('图片类型步骤请上传示例图片'); return; }
  stepSaving.value = true;
  try {
    const payload = {
      activityID: activityId.value,
      title: stepForm.title,
      stepNum: stepForm.stepNum,
      stepType: stepForm.stepType,
      exampleText: stepForm.exampleText,
      descContent: stepForm.descContent,
      stepImage: stepForm.stepImage,
      isRequired: stepForm.isRequired,
      sort: stepForm.sort,
    };
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
        activityType.value = data.type ?? 0;
        activeTab.value = data.tab || 'rewards';
        loadData(data.id);
      }
    }
  },
});
</script>

<template>
  <DrawerModal :title="`活动管理 - ${activityTitle}`" :footer="null" width="70%">
    <Tabs v-model:activeKey="activeTab">

      <!-- ====== 奖励管理 ====== -->
      <Tabs.TabPane key="rewards" tab="奖励管理">
        <div style="margin-bottom: 12px">
          <Button type="primary" size="small" @click="handleAddReward">新增奖励</Button>
        </div>
        <Table :columns="rewardColumns" :data-source="rewards" :loading="loading" :pagination="false" row-key="id" size="small">
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'action'">
              <Button type="link" size="small" @click="handleEditReward(record)">编辑</Button>
              <Popconfirm title="确定删除该奖励？" @confirm="handleDeleteReward(record)">
                <Button type="link" danger size="small">删除</Button>
              </Popconfirm>
            </template>
          </template>
        </Table>
      </Tabs.TabPane>

      <!-- ====== 步骤管理（仅图文步骤活动 type=4 显示） ====== -->
      <Tabs.TabPane v-if="activityType === 4" key="steps" tab="步骤管理">
        <div style="margin-bottom: 12px; display: flex; align-items: center; gap: 8px;">
          <Button type="primary" size="small" @click="handleAddStep">新增步骤</Button>
          <span style="font-size: 12px; color: #999;">步骤按序号顺序执行，用户需依次完成</span>
        </div>
        <Table :columns="stepColumns" :data-source="steps" :loading="loading" :pagination="false" row-key="id" size="small">
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'stepType'">
              <Tag :color="stepTypeColor[record.stepType] || 'default'">
                {{ stepTypeMap[record.stepType] || '文字' }}
              </Tag>
            </template>
            <template v-if="column.key === 'isRequired'">
              <Tag :color="record.isRequired === 1 ? 'blue' : 'default'">
                {{ record.isRequired === 1 ? '需要填写' : '不需要' }}
              </Tag>
            </template>
            <template v-if="column.key === 'stepImage'">
              <img
                v-if="record.stepImage"
                :src="record.stepImage"
                style="width:40px;height:40px;object-fit:cover;border-radius:4px;cursor:pointer;"
                @click="() => {}"
              />
              <span v-else style="color:#ccc;">-</span>
            </template>
            <template v-if="column.key === 'action'">
              <Button type="link" size="small" @click="handleEditStep(record)">编辑</Button>
              <Popconfirm title="确定删除该步骤？" @confirm="handleDeleteStep(record)">
                <Button type="link" danger size="small">删除</Button>
              </Popconfirm>
            </template>
          </template>
        </Table>
      </Tabs.TabPane>

    </Tabs>

    <!-- ====== 奖励编辑弹窗 ====== -->
    <Modal
      v-model:open="rewardModalVisible"
      :title="rewardForm.id ? '编辑奖励' : '新增奖励'"
      :confirm-loading="rewardSaving"
      @ok="handleSaveReward"
      width="50%"
    >
      <Form layout="vertical" style="margin-top: 16px;">
        <Form.Item label="奖励名称" required>
          <Input v-model:value="rewardForm.rewardName" placeholder="请输入奖励展示名称" />
        </Form.Item>
        <Form.Item label="奖励类型" required>
          <Select v-model:value="rewardForm.rewardType" :options="rewardTypeOptions" style="width: 100%" />
        </Form.Item>
        <!-- 余额 / 经验值 -->
        <Form.Item v-if="rewardForm.rewardType === 1 || rewardForm.rewardType === 3" :label="rewardValueLabel">
          <InputNumber v-model:value="rewardForm.rewardValue" :min="0" style="width: 100%" :placeholder="rewardForm.rewardType === 1 ? '请输入金额（单位：分）' : '请输入经验值'" />
        </Form.Item>

        <!-- 优惠券：远程搜索下拉 -->
        <Form.Item v-if="rewardForm.rewardType === 2" label="选择优惠券">
          <Select
            v-model:value="rewardForm.rewardValue"
            show-search
            :filter-option="false"
            :options="couponOptions"
            :loading="couponSearching"
            placeholder="请输入优惠券名称搜索"
            @search="handleCouponSearch"
            style="width: 100%"
          />
        </Form.Item>

        <!-- 会员等级天数：选等级 + 输入天数 -->
        <Form.Item v-if="rewardForm.rewardType === 4" label="会员等级">
          <Select
            v-model:value="rewardForm.rewardLevelId"
            :options="levelOptions"
            placeholder="请选择会员等级"
            style="width: 100%"
          />
        </Form.Item>
        <Form.Item v-if="rewardForm.rewardType === 4" label="天数">
          <InputNumber v-model:value="rewardForm.rewardDays" :min="1" style="width: 100%" placeholder="请输入天数" />
        </Form.Item>
        <Form.Item label="排序">
          <InputNumber v-model:value="rewardForm.sort" :min="0" style="width: 100%" />
        </Form.Item>
      </Form>
    </Modal>

    <!-- ====== 步骤编辑弹窗 ====== -->
    <Modal
      v-model:open="stepModalVisible"
      :title="stepForm.id ? '编辑步骤' : '新增步骤'"
      :confirm-loading="stepSaving"
      @ok="handleSaveStep"
      width="50%"
    >
      <Form layout="vertical" style="margin-top: 16px;">
        <Form.Item label="步骤标题" required>
          <Input v-model:value="stepForm.title" placeholder="请输入步骤标题" />
        </Form.Item>
        <Form.Item label="步骤序号" help="决定用户完成顺序，序号小的先完成">
          <InputNumber v-model:value="stepForm.stepNum" :min="1" style="width: 100%" />
        </Form.Item>
        <Form.Item label="步骤类型" required>
          <Select v-model:value="stepForm.stepType" :options="stepTypeOptions" style="width: 100%" />
        </Form.Item>

        <!-- 文字类型：示例文字 -->
        <Form.Item v-if="stepForm.stepType === 1" label="示例文字" help="WAP 端展示给用户参考，用户可一键复制">
          <Input.TextArea v-model:value="stepForm.exampleText" placeholder="请输入用户需要复制/参考的文字内容" :rows="3" />
        </Form.Item>

        <!-- 链接类型：目标链接 -->
        <Form.Item v-if="stepForm.stepType === 2" label="目标链接" help="WAP 端展示给用户，用户可一键前往">
          <Input v-model:value="stepForm.exampleText" placeholder="请输入 http(s):// 开头的完整链接地址" />
        </Form.Item>

        <!-- 图片类型：示例图片 -->
        <Form.Item v-if="stepForm.stepType === 3" label="示例图片" required help="WAP 端左侧展示，右侧为用户上传区域">
          <ImageUpload v-model:value="stepForm.stepImage" />
        </Form.Item>

        <Form.Item label="是否需要填写" help="开启后 WAP 端用户需根据步骤类型填写内容或上传图片">
          <Select
            v-model:value="stepForm.isRequired"
            :options="[{ label: '不需要', value: 0 }, { label: '需要填写', value: 1 }]"
            style="width: 100%"
          />
        </Form.Item>
        <Form.Item label="步骤说明" help="对该步骤的补充说明，展示在步骤标题下方">
          <Input.TextArea v-model:value="stepForm.descContent" placeholder="请输入步骤说明（可选）" :rows="3" />
        </Form.Item>
        <Form.Item label="排序">
          <InputNumber v-model:value="stepForm.sort" :min="0" style="width: 100%" />
        </Form.Item>
      </Form>
    </Modal>
  </DrawerModal>
</template>
