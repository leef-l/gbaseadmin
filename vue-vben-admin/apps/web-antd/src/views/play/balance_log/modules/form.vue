<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getBalanceLogDetail,
  createBalanceLog,
  updateBalanceLog,
} from '#/api/play/balance_log';
import { getMemberList } from '#/api/play/member';

/** 业务类型选项 */
const bizTypeOptions = [
  { label: '充值', value: 1 },
  { label: '消费', value: 2 },
  { label: '退款', value: 3 },
  { label: '活动赠送', value: 4 },
  { label: '提现', value: 5 },
];

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
      fieldName: 'memberID',
      label: '会员ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'bizType',
      label: '业务类型',
      rules: 'selectRequired',
      componentProps: { options: bizTypeOptions, placeholder: '请选择业务类型', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'bizID',
      label: () => h('span', {}, ['关联业务ID ', h(Tooltip, { title: '订单ID/充值订单ID/活动ID' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入业务ID' },
    },
    {
      component: 'InputNumber',
      fieldName: 'changeAmount',
      label: () => h('span', {}, ['变动金额 ', h(Tooltip, { title: '分，正数增加负数减少' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      rules: 'required',
      componentProps: { placeholder: '请输入变动金额（分，正数增加负数减少）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'beforeBalance',
      label: () => h('span', {}, ['变动前余额 ', h(Tooltip, { title: '分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      rules: 'required',
      componentProps: { placeholder: '请输入变动前余额（分）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'afterBalance',
      label: () => h('span', {}, ['变动后余额 ', h(Tooltip, { title: '分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      rules: 'required',
      componentProps: { placeholder: '请输入变动后余额（分）', class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'remark',
      label: '备注说明',
      componentProps: { placeholder: '请输入备注说明', maxlength: 200 },
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
        await updateBalanceLog({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createBalanceLog(values);
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
        modalApi.setState({ title: '编辑余额流水表' });
        try {
          const detail = await getBalanceLogDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建余额流水表' });
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
