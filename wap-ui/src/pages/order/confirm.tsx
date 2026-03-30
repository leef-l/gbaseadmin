import { useState } from 'react';
import { View, Text, Input } from '@tarojs/components';
import Taro, { useLoad } from '@tarojs/taro';
import { createOrder } from '../../api/order';
import { getUsableCoupons } from '../../api/coupon';
import { useCartStore } from '../../store/cart';
import './confirm.scss';

export default function OrderConfirmPage() {
  const { goodsId, goodsName, coachId, coachName, price, quantity, couponId, couponAmount, remark, setOrder } = useCartStore();
  const [coupons, setCoupons] = useState<any[]>([]);
  const [showCoupon, setShowCoupon] = useState(false);
  const totalAmount = price * quantity;
  const payAmount = Math.max(0, totalAmount - couponAmount);

  useLoad(async () => {
    try {
      const data = await getUsableCoupons(totalAmount);
      setCoupons(Array.isArray(data) ? data : data?.list || []);
    } catch {}
  });

  const selectCoupon = (c: any) => {
    setOrder({ couponId: c.couponMemberId, couponAmount: c.faceValue });
    setShowCoupon(false);
  };

  const handleSubmit = async () => {
    try {
      Taro.showLoading({ title: '提交中...' });
      const data = await createOrder({
        goodsId,
        coachId,
        quantity,
        couponId: couponId || undefined,
        remark: remark || undefined,
      });
      Taro.hideLoading();
      Taro.showToast({ title: '下单成功', icon: 'success' });
      setTimeout(() => Taro.redirectTo({ url: `/pages/order/pay?orderId=${data?.orderId}` }), 1500);
    } catch {
      Taro.hideLoading();
    }
  };

  return (
    <View className="order-confirm">
      <View className="order-confirm__goods card">
        <View className="order-confirm__cover" />
        <View className="order-confirm__info">
          <Text className="order-confirm__name">{goodsName}</Text>
          <Text className="order-confirm__coach">{coachName}</Text>
          <View className="order-confirm__price-row">
            <Text className="price">¥{(price / 100).toFixed(2)}</Text>
            <View className="order-confirm__quantity">
              <View className="order-confirm__qty-btn" onClick={() => quantity > 1 && setOrder({ quantity: quantity - 1 })}>-</View>
              <Text className="order-confirm__qty-num">{quantity}</Text>
              <View className="order-confirm__qty-btn" onClick={() => setOrder({ quantity: quantity + 1 })}>+</View>
            </View>
          </View>
        </View>
      </View>

      <View className="order-confirm__coupon card" onClick={() => setShowCoupon(!showCoupon)}>
        <Text className="order-confirm__coupon-label">优惠券</Text>
        <Text className="order-confirm__coupon-value">{couponAmount > 0 ? `-¥${(couponAmount / 100).toFixed(2)}` : coupons.length > 0 ? `${coupons.length}张可用` : '暂无可用'} &gt;</Text>
      </View>

      {showCoupon && coupons.length > 0 && (
        <View className="order-confirm__coupon-list card">
          {coupons.map((c: any) => (
            <View key={c.couponMemberId} className={`order-confirm__coupon-item ${couponId === c.couponMemberId ? 'order-confirm__coupon-item--active' : ''}`} onClick={() => selectCoupon(c)}>
              <Text className="order-confirm__coupon-amount">¥{(c.faceValue / 100).toFixed(2)}</Text>
              <Text className="order-confirm__coupon-name">{c.title || '优惠券'}</Text>
            </View>
          ))}
        </View>
      )}

      <View className="order-confirm__remark card">
        <Input className="order-confirm__remark-input" placeholder="给陪玩师留言（选填）" value={remark} onInput={(e) => setOrder({ remark: e.detail.value })} />
      </View>

      <View className="order-confirm__amount card">
        <View className="order-confirm__amount-row">
          <Text>商品金额</Text>
          <Text>¥{(totalAmount / 100).toFixed(2)}</Text>
        </View>
        {couponAmount > 0 && (
          <View className="order-confirm__amount-row">
            <Text>优惠券</Text>
            <Text className="order-confirm__amount-discount">-¥{(couponAmount / 100).toFixed(2)}</Text>
          </View>
        )}
        <View className="order-confirm__amount-total">
          <Text>实付金额</Text>
          <Text className="order-confirm__amount-total-price">¥{(payAmount / 100).toFixed(2)}</Text>
        </View>
      </View>

      <View className="bottom-bar">
        <Text>合计 <Text className="price" style={{ fontSize: '18px' }}>¥{(payAmount / 100).toFixed(2)}</Text></Text>
        <View className="btn-primary" onClick={handleSubmit}>提交订单</View>
      </View>
    </View>
  );
}
