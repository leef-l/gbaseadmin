<script setup lang="ts">
import { h, ref } from 'vue';
import { useVbenModal } from '@vben/common-ui';
import { useVbenForm } from '#/adapter/form';
import { message, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {
  getOrderDetail,
  createOrder,
  updateOrder,
} from '#/api/play/order';
import { getMemberList } from '#/api/play/member';
import { getCoachList } from '#/api/play/coach';
import { getShopList } from '#/api/play/shop';
import { getGoodsList } from '#/api/play/goods';
import { getCouponMemberList } from '#/api/play/coupon_member';

/** 支付方式选项 */
const payTypeOptions = [
  { label: '未支付', value: 0 },
  { label: '微信支付', value: 1 },
  { label: '支付宝支付', value: 2 },
  { label: '余额支付', value: 3 },
];

/** 订单状态选项 */
const orderStatusOptions = [
  { label: '待支付', value: 0 },
  { label: '已支付', value: 1 },
  { label: '进行中', value: 2 },
  { label: '已完成', value: 3 },
  { label: '已取消', value: 4 },
  { label: '退款中', value: 5 },
  { label: '已退款', value: 6 },
];

const memberIDOptions = ref<{ label: string; value: string }[]>([]);
const coachIDOptions = ref<{ label: string; value: string }[]>([]);
const shopIDOptions = ref<{ label: string; value: string }[]>([]);
const goodsIDOptions = ref<{ label: string; value: string }[]>([]);
const couponMemberIDOptions = ref<{ label: string; value: string }[]>([]);

const emit = defineEmits<{ success: [] }>();
const isEdit = ref(false);
const editId = ref('');

/** 表单配置 */
const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  schema: [
    {
      component: 'Input',
      fieldName: 'orderNo',
      label: '订单编号',
      rules: 'required',
      componentProps: { placeholder: '请输入订单编号', maxlength: 32 },
    },
    {
      component: 'Select',
      fieldName: 'memberID',
      label: '下单会员ID',
      rules: 'selectRequired',
      componentProps: { options: memberIDOptions, placeholder: '请选择下单会员ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'coachID',
      label: '陪玩师ID',
      rules: 'selectRequired',
      componentProps: { options: coachIDOptions, placeholder: '请选择陪玩师ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'shopID',
      label: () => h('span', {}, ['店铺ID ', h(Tooltip, { title: '0表示无店铺' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { options: shopIDOptions, placeholder: '请选择店铺ID（0表示无店铺）', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'goodsID',
      label: '商品ID',
      rules: 'selectRequired',
      componentProps: { options: goodsIDOptions, placeholder: '请选择商品ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Input',
      fieldName: 'goodsTitle',
      label: () => h('span', {}, ['商品名称 ', h(Tooltip, { title: '冗余' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      rules: 'required',
      componentProps: { placeholder: '请输入商品名称（冗余）', maxlength: 100 },
    },
    {
      component: 'InputNumber',
      fieldName: 'goodsPrice',
      label: () => h('span', {}, ['商品单价 ', h(Tooltip, { title: '分，下单时快照' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      rules: 'required',
      componentProps: { placeholder: '请输入商品单价（分，下单时快照）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'quantity',
      label: '数量',
      componentProps: { placeholder: '请输入数量', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'totalAmount',
      label: () => h('span', {}, ['订单总额 ', h(Tooltip, { title: '分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入订单总额（分）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'discountAmount',
      label: () => h('span', {}, ['会员折扣金额 ', h(Tooltip, { title: '分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入会员折扣金额（分）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'couponAmount',
      label: () => h('span', {}, ['优惠券抵扣金额 ', h(Tooltip, { title: '分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入优惠券抵扣金额（分）', class: 'w-full' },
    },
    {
      component: 'InputNumber',
      fieldName: 'payAmount',
      label: () => h('span', {}, ['实付金额 ', h(Tooltip, { title: '分' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })]),
      componentProps: { placeholder: '请输入实付金额（分）', class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'couponMemberID',
      label: '使用的优惠券领取记录ID',
      componentProps: { options: couponMemberIDOptions, placeholder: '请选择使用的优惠券领取记录ID', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'payType',
      label: '支付方式',
      componentProps: { options: payTypeOptions, placeholder: '请选择支付方式', allowClear: true, class: 'w-full' },
    },
    {
      component: 'Select',
      fieldName: 'orderStatus',
      label: '订单状态',
      componentProps: { options: orderStatusOptions, placeholder: '请选择订单状态', allowClear: true, class: 'w-full' },
    },
    {
      component: 'DatePicker',
      fieldName: 'payAt',
      label: '支付时间',
      componentProps: { showTime: true, placeholder: '请选择支付时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'startAt',
      label: '服务开始时间',
      componentProps: { showTime: true, placeholder: '请选择服务开始时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'finishAt',
      label: '服务完成时间',
      componentProps: { showTime: true, placeholder: '请选择服务完成时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'DatePicker',
      fieldName: 'cancelAt',
      label: '取消时间',
      componentProps: { showTime: true, placeholder: '请选择取消时间', class: 'w-full', valueFormat: 'YYYY-MM-DD HH:mm:ss' },
    },
    {
      component: 'Input',
      fieldName: 'cancelReason',
      label: '取消原因',
      componentProps: { placeholder: '请输入取消原因', maxlength: 500 },
    },
    {
      component: 'Input',
      fieldName: 'remark',
      label: '订单备注',
      componentProps: { placeholder: '请输入订单备注', maxlength: 500 },
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
        await updateOrder({ id: editId.value, ...values });
        message.success('更新成功');
      } else {
        await createOrder(values);
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
      // 加载陪玩师选项
      try {
        const coachRes = await getCoachList({ pageNum: 1, pageSize: 1000 });
        coachIDOptions.value = (coachRes?.list ?? []).map((item: any) => ({
          label: item.realName || item.id,
          value: item.id,
        }));
      } catch {
        // ignore
      }
      // 加载店铺选项
      try {
        const shopRes = await getShopList({ pageNum: 1, pageSize: 1000 });
        shopIDOptions.value = (shopRes?.list ?? []).map((item: any) => ({
          label: item.title || item.id,
          value: item.id,
        }));
      } catch {
        // ignore
      }
      // 加载商品选项
      try {
        const goodsRes = await getGoodsList({ pageNum: 1, pageSize: 1000 });
        goodsIDOptions.value = (goodsRes?.list ?? []).map((item: any) => ({
          label: item.title || item.id,
          value: item.id,
        }));
      } catch {
        // ignore
      }
      // 加载会员优惠券选项
      try {
        const couponMemberRes = await getCouponMemberList({ pageNum: 1, pageSize: 1000 });
        couponMemberIDOptions.value = (couponMemberRes?.list ?? []).map((item: any) => ({
          label: item.couponTitle || item.id,
          value: item.id,
        }));
      } catch {
        // ignore
      }
      if (data?.id) {
        isEdit.value = true;
        editId.value = data.id;
        modalApi.setState({ title: '编辑订单表' });
        try {
          const detail = await getOrderDetail(data.id);
          if (detail) {
            formApi.setValues(detail);
          }
        } catch {
          message.error('获取详情失败');
        }
      } else {
        isEdit.value = false;
        editId.value = '';
        modalApi.setState({ title: '新建订单表' });
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
