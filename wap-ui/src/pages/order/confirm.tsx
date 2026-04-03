import { useState } from 'react';
import { View, Text, Textarea, Button } from '@tarojs/components';
import Taro, { useLoad, useRouter } from '@tarojs/taro';
import { createOrder } from '../../api/order';
import { getUsableCoupons } from '../../api/coupon';
import { useCartStore } from '../../store/cart';
import './confirm.scss';

export default function OrderConfirmPage() {
  const { params } = useRouter();
  const { goodsId, goodsName, coachId, coachName, price, quantity, couponId, couponAmount, remark, setOrder, startOrder, reset } = useCartStore();
  const [coupons, setCoupons] = useState<any[]>([]);
  const [showCoupon, setShowCoupon] = useState(false);
  const [couponLoaded, setCouponLoaded] = useState(false);
  const totalAmount = price * quantity;
  const payAmount = Math.max(0, totalAmount - couponAmount);

  useLoad(() => {
    if (!params.goodsId) {
      if (goodsId && price > 0) return;
      Taro.showToast({ title: '订单信息缺失', icon: 'none' });
      setTimeout(() => Taro.navigateBack(), 1500);
      return;
    }
    if (params.goodsId === goodsId && price > 0) return;
    startOrder({
      goodsId: params.goodsId,
      goodsName: params.goodsName ? decodeURIComponent(params.goodsName) : goodsName,
      coachId,
      coachName: params.coachName ? decodeURIComponent(params.coachName) : coachName,
      price: Number(params.price || 0),
      quantity,
    });
  });

  const changeQuantity = (nextQuantity: number) => {
    const normalizedQuantity = Math.max(1, nextQuantity);
    setOrder({
      quantity: normalizedQuantity,
      couponId: '',
      couponAmount: 0,
    });
    setCoupons([]);
    setCouponLoaded(false);
    setShowCoupon(false);
  };

  const fetchCoupons = async () => {
    if (couponLoaded || totalAmount <= 0) return;
    try {
      const data = await getUsableCoupons(totalAmount);
      setCoupons(Array.isArray(data) ? data : data?.list || []);
    } catch {
      setCoupons([]);
    } finally {
      setCouponLoaded(true);
    }
  };

  const selectCoupon = (c: any) => {
    setOrder({ couponId: c.couponMemberId, couponAmount: c.faceValue });
    setShowCoupon(false);
  };

  const handleSubmit = async () => {
    if (!goodsId || quantity <= 0 || price <= 0) {
      Taro.showToast({ title: '订单信息不完整', icon: 'none' });
      return;
    }
    try {
      Taro.showLoading({ title: '提交中...' });
      const data = await createOrder({
        goodsId,
        quantity,
        couponMemberId: couponId || undefined,
        remark: remark || undefined,
      });
      Taro.hideLoading();
      Taro.showToast({ title: '下单成功', icon: 'success' });
      setTimeout(() => {
        reset();
        Taro.redirectTo({ url: `/pages/order/pay?orderId=${data?.orderId}` });
      }, 1500);
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
               <View className="order-confirm__qty-btn" onClick={() => changeQuantity(quantity - 1)}>-</View>
               <Text className="order-confirm__qty-num">{quantity}</Text>
               <View className="order-confirm__qty-btn" onClick={() => changeQuantity(quantity + 1)}>+</View>
             </View>
           </View>
         </View>
      </View>

      <View
        className="order-confirm__coupon card"
        onClick={async () => {
          const next = !showCoupon;
          setShowCoupon(next);
          if (next && !couponLoaded) {
            await fetchCoupons();
          }
        }}
      >
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
        <Textarea
          className="order-confirm__remark-input"
          placeholder="给陪玩师留言（选填）"
          value={remark}
          autoHeight
          maxlength={200}
          onInput={(e) => setOrder({ remark: e.detail.value })}
        />
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
        <Button className="btn-primary" onClick={handleSubmit}>提交订单</Button>
      </View>
    </View>
  );
}
