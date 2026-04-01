<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getBannerDetail,
  createBanner,
  updateBanner,
} from '#/api/play/banner';

/** 渲染带 Tooltip 的表单 label */
function tooltipLabel(label: string, tip: string) {
  return () => h('span', {}, [
    label + ' ',
    h(Tooltip, { title: tip }, {
      default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }),
    }),
  ]);
}

/** 跳转类型 placeholder 映射 */
const linkTypePlaceholderMap: Record<number, string> = {
  1: '请输入页面路径，如 /pages/recharge/index',
  2: '请输入完整URL，如 https://example.com',
  3: '请输入活动ID',
  4: '请输入商品ID',
  5: '请输入陪玩师ID',
  6: '请输入App Scheme，如 weixin:// 或 alipays://',
};

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'title',
      label: 'Banner标题',
      componentProps: { placeholder: '请输入Banner标题', maxlength: 100 },
    },
    {
      component: 'ImageUpload',
      fieldName: 'image',
      label: 'Banner图片',
      rules: 'required',
      componentProps: { maxCount: 1 },
    },
    {
      component: 'Select',
      fieldName: 'linkType',
      label: '跳转类型',
      rules: 'required',
      defaultValue: 1,
      componentProps: {
        options: [
          { label: '内页', value: 1 },
          { label: '外链', value: 2 },
          { label: '活动页', value: 3 },
          { label: '商品页', value: 4 },
          { label: '陪玩师页', value: 5 },
          { label: '唤醒App', value: 6 },
        ],
        class: 'w-full',
      },
    },
    {
      component: 'Input',
      fieldName: 'linkValue',
      label: tooltipLabel('跳转值', '页面路径/URL/业务ID/App Scheme'),
      rules: 'required',
      componentProps: { placeholder: '请输入跳转值', maxlength: 500 },
      dependencies: {
        triggerFields: ['linkType'],
        componentProps: (values: Record<string, any>) => {
          return {
            placeholder: linkTypePlaceholderMap[values.linkType as number] ?? '请输入跳转值',
          };
        },
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: tooltipLabel('排序', '越大越前'),
      defaultValue: 0,
      componentProps: { placeholder: '请输入排序(越大越前)', class: 'w-full' },
    },
    {
      component: 'Switch',
      fieldName: 'status',
      label: '状态',
      componentProps: { checkedValue: 1, unCheckedValue: 0 },
      defaultValue: 1,
    },
    {
      component: 'DatePicker',
      fieldName: 'startTime',
      label: '生效开始时间',
      componentProps: { showTime: true, format: 'YYYY-MM-DD HH:mm:ss', valueFormat: 'YYYY-MM-DD HH:mm:ss', class: 'w-full' },
    },
    {
      component: 'DatePicker',
      fieldName: 'endTime',
      label: '生效结束时间',
      componentProps: { showTime: true, format: 'YYYY-MM-DD HH:mm:ss', valueFormat: 'YYYY-MM-DD HH:mm:ss', class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'remark',
      label: '备注',
      componentProps: { placeholder: '请输入备注', maxlength: 255 },
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
        await updateBanner({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createBanner(values);
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
      if (data?.id) {
        isEdit.value = true;
        editId.value = data.id;
        modalApi.setState({ title: '编辑首页Banner轮播' });
        try {
          const detail = await getBannerDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建首页Banner轮播' });
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
