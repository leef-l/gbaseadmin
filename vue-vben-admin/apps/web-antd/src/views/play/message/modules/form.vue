<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getMessageDetail,
  createMessage,
  updateMessage,
} from '#/api/play/message';
import { getMemberList } from '#/api/play/member';

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
      fieldName: 'memberID',
      label: '接收者会员ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择接收者会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'title',
      label: '消息标题',
      componentProps: { placeholder: '请输入消息标题', maxlength: 200 },
    },
    {
      component: 'Textarea',
      fieldName: 'content',
      label: '消息内容',
      componentProps: { placeholder: '请输入消息内容', rows: 4, maxlength: 65535 },
    },
    {
      component: 'Input',
      fieldName: 'msgType',
      label: '消息类型 1=系统通知 2=订单消息 3=活动消息',
      componentProps: { placeholder: '请输入消息类型 1=系统通知 2=订单消息 3=活动消息' },
    },
    {
      component: 'Input',
      fieldName: 'bizID',
      label: tooltipLabel('关联业务ID', '订单ID/活动ID等'),
      componentProps: { placeholder: '请输入关联业务ID（订单ID/活动ID等）' },
    },
    {
      component: 'Switch',
      fieldName: 'isRead',
      label: '是否已读 0=未读 1=已读',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 0,
    },
    {
      component: 'Switch',
      fieldName: 'status',
      label: '状态 1=正常 0=禁用',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 1,
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
        await updateMessage({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createMessage(values);
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
      // 加载接收者会员ID选项
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
        modalApi.setState({ title: '编辑会员消息' });
        try {
          const detail = await getMessageDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建会员消息' });
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
