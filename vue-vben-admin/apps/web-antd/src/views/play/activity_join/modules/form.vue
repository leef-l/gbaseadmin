<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getActivityJoinDetail,
  createActivityJoin,
  updateActivityJoin,
} from '#/api/play/activity_join';
import { getActivityList } from '#/api/play/activity';
import { getMemberList } from '#/api/play/member';

/** 参与状态选项 */
const joinStatusOptions = [
  { label: '已报名', value: 0 },
  { label: '进行中', value: 1 },
  { label: '已完成', value: 2 },
  { label: '已领奖', value: 3 },
];

const activityIDOptions = ref<{ label: string; value: string }[]>([]);
const memberIDOptions = ref<{ label: string; value: string }[]>([]);

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Select',
      fieldName: 'activityID',
      label: '活动ID',
      rules: 'selectRequired',
      componentProps: { options: activityIDOptions, placeholder: '请选择活动ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'memberID',
      label: '会员ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'joinStatus',
      label: '参与状态',
      componentProps: { options: joinStatusOptions, placeholder: '请选择参与状态', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'currentStep',
      label: () => h('span', {}, ['当前完成到第几步 ', h(Tooltip, { title: '步骤活动用' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入当前完成到第几步（步骤活动用）' },
    },
    {
      component: 'DatePicker',
      fieldName: 'finishAt',
      label: '完成时间',
      componentProps: { showTime: true, placeholder: '请选择完成时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'rewardAt',
      label: '领奖时间',
      componentProps: { showTime: true, placeholder: '请选择领奖时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'Input',
      fieldName: 'remark',
      label: '备注',
      componentProps: { placeholder: '请输入备注', maxlength: 500 },
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
        await updateActivityJoin({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createActivityJoin(values);
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
      // 加载活动选项
      try {
        const activityRes = await getActivityList({ pageNum: 1, pageSize: 1000 });
        activityIDOptions.value = (activityRes?.list ?? []).map((item: any) => ({
          label: item.title || item.id,
          value: item.id,
        }));
      } catch {
        // ignore
      }
      // 加载会员选项
      try {
        const memberRes = await getMemberList({ pageNum: 1, pageSize: 1000 });
        memberIDOptions.value = (memberRes?.list ?? []).map((item: any) => ({
          label: item.nickname || item.id,
          value: item.id,
        }));
      } catch {
        // ignore
      }
      if (data?.id) {
        isEdit.value = true;
        editId.value = data.id;
        modalApi.setState({ title: '编辑活动参与记录表' });
        try {
          const detail = await getActivityJoinDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建活动参与记录表' });
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
