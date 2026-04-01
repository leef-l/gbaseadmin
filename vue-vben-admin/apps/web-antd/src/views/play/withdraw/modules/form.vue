<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getWithdrawDetail,
  createWithdraw,
  updateWithdraw,
} from '#/api/play/withdraw';
import { getCoachList } from '#/api/play/coach';
import { getMemberList } from '#/api/play/member';

const coachIDOptions = ref<{ label: string; value: string }[]>([]);
const memberIDOptions = ref<{ label: string; value: string }[]>([]);
/** 渲染带 Tooltip 的表单 label */
function tooltipLabel(label: string, tip: string) {
  return () => h('span', {}, [
    label + ' ',
    h(Tooltip, { title: tip }, {
      default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }),
    }),
  ]);
}

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Select',
      fieldName: 'coachID',
      label: '陪玩师ID',
      rules: 'selectRequired',
      componentProps: { options: coachIDOptions, placeholder: '请选择陪玩师ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'memberID',
      label: '会员ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'amount',
      label: tooltipLabel('提现金额', '分'),
      componentProps: { placeholder: '请输入提现金额(分)', class: 'w-full' },
    },
    {
      component: 'Switch',
      fieldName: 'status',
      label: '状态 0=待审核 1=已打款 2=已拒绝',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 0,
    },
    {
      component: 'Input',
      fieldName: 'reason',
      label: '拒绝原因',
      componentProps: { placeholder: '请输入拒绝原因', maxlength: 500 },
    },
    {
      component: 'DatePicker',
      fieldName: 'auditedAt',
      label: '审核时间',
      componentProps: { showTime: true, placeholder: '请选择审核时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
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
        await updateWithdraw({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createWithdraw(values);
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
      // 加载陪玩师ID选项
      try {
        const coachRes = await getCoachList({ pageNum: 1, pageSize: 1000 });
        coachIDOptions.value = (coachRes?.list ?? []).map((item: any) => ({
          label: item.realName || item.id,
          value: item.id,
        }));
      } catch {
        // ignore
      }
      // 加载会员ID选项
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
        modalApi.setState({ title: '编辑陪玩师提现记录' });
        try {
          const detail = await getWithdrawDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建陪玩师提现记录' });
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
