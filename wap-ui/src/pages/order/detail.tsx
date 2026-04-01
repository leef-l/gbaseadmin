import { useState } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, useRouter } from '@tarojs/taro';
import { getOrderDetail, cancelOrder, refundOrder } from '../../api/order';
import './detail.scss';

const statusConfig: Record<number, { text: string; bg: string; desc: string }> = {
  0: { text: '待支付', bg: 'linear-gradient(135deg, #e17055, #fab1a0)', desc: '请尽快完成支付' },
  1: { text: '已支付', bg: 'linear-gradient(135deg, #0984e3, #74b9ff)', desc: '等待陪玩师接单' },
  2: { text: '进行中', bg: 'linear-gradient(135deg, #7c3aed, #6d28d9)', desc: '服务进行中' },
  3: { text: '已完成', bg: 'linear-gradient(135deg, #00b894, #55efc4)', desc: '服务已完成' },
  4: { text: '已取消', bg: 'linear-gradient(135deg, #b2bec3, #dfe6e9)', desc: '订单已取消' },
  5: { text: '退款中', bg: 'linear-gradient(135deg, #d63031, #ff7675)', desc: '退款处理中' },
  6: { text: '已退款', bg: 'linear-gradient(135deg, #b2bec3, #dfe6e9)', desc: '退款已完成' },
};

export default function OrderDetailPage() {
  const { params } = useRouter();
  const [order, setOrder] = useState<any>(null);

  useLoad(async () => {
    try {
      const data = await getOrderDetail(params.id || '');
      setOrder(data);
    } catch {
      Taro.showToast({ title: '加载失败', icon: 'none' });
    }
  });

  const handleCancel = async () => {
    try {
      await cancelOrder(order.orderId);
      Taro.showToast({ title: '已取消', icon: 'success' });
      setOrder({ ...order, orderStatus: 4 });
    } catch {}
  };

  const handleRefund = () => {
    Taro.showModal({
      title: '申请退款',
      content: '确认申请退款吗？退款将在1-3个工作日内处理。',
      confirmText: '确认退款',
      cancelText: '取消',
      success: async ({ confirm }) => {
        if (!confirm) return;
        try {
          await refundOrder(order.orderId, '用户申请退款');
          Taro.showToast({ title: '退款申请已提交', icon: 'success' });
          const data = await getOrderDetail(params.id || '');
          setOrder(data);
        } catch {}
      },
    });
  };

  if (!order) return <View />;
  const sc = statusConfig[order.orderStatus] || statusConfig[4];

  return (
    <View className="order-detail">
      <View className="order-detail__status" style={{ background: sc.bg }}>
        <Text className="order-detail__status-text">{sc.text}</Text>
        <Text className="order-detail__status-desc">{sc.desc}</Text>
      </View>

      <View className="order-detail__goods card">
        <View className="order-detail__cover" />
        <View className="order-detail__info">
          <Text className="order-detail__name">{order.goodsTitle}</Text>
          <Text className="order-detail__coach">{order.coachName}</Text>
          <Text className="order-detail__spec">{order.quantity} x ¥{(order.goodsPrice / 100).toFixed(2)}</Text>
        </View>
      </View>

      <View className="order-detail__amount card">
        <View className="order-detail__amount-row"><Text>商品金额</Text><Text>¥{(order.goodsPrice * order.quantity / 100).toFixed(2)}</Text></View>
        {order.couponAmount > 0 && (
          <View className="order-detail__amount-row"><Text>优惠券</Text><Text style={{ color: 'var(--success)' }}>-¥{(order.couponAmount / 100).toFixed(2)}</Text></View>
        )}
        <View className="order-detail__amount-total"><Text>实付金额</Text><Text className="price">¥{(order.payAmount / 100).toFixed(2)}</Text></View>
      </View>

      <View className="order-detail__meta card">
        <View className="order-detail__meta-row"><Text className="order-detail__meta-label">订单编号</Text><Text className="order-detail__meta-value">{order.orderNo}</Text></View>
        <View className="order-detail__meta-row"><Text className="order-detail__meta-label">下单时间</Text><Text className="order-detail__meta-value">{order.createdAt}</Text></View>
      </View>

      <View className="bottom-bar">
        {order.orderStatus === 0 && <>
          <View style={{ flex: 1 }} />
          <View style={{ border: '1px solid var(--border)', borderRadius: '20px', padding: '8px 20px', fontSize: '13px', marginRight: '8px' }} onClick={handleCancel}>取消订单</View>
          <View className="btn-primary" onClick={() => Taro.navigateTo({ url: `/pages/order/pay?orderId=${order.orderId}` })}>去支付</View>
        </>}
        {order.orderStatus === 2 && <>
          <View style={{ flex: 1 }} />
          <View style={{ border: '1px solid var(--border)', borderRadius: '20px', padding: '8px 20px', fontSize: '13px', marginRight: '8px' }} onClick={handleRefund}>申请退款</View>
        </>}
        {order.orderStatus === 3 && <>
          <View style={{ flex: 1 }} />
          <View style={{ border: '1px solid var(--border)', borderRadius: '20px', padding: '8px 20px', fontSize: '13px', marginRight: '8px' }} onClick={() => order.goodsId && Taro.navigateTo({ url: `/pages/goods/detail?id=${order.goodsId}` })}>再来一单</View>
          <View className="btn-primary" onClick={() => Taro.navigateTo({ url: `/pages/order/review?orderId=${order.orderId}` })}>去评价</View>
        </>}
      </View>
    </View>
  );
}
