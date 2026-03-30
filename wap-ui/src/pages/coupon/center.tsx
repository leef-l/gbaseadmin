import { useState, useCallback } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad } from '@tarojs/taro';
import { getCouponCenter, receiveCoupon } from '../../api/coupon';
import './center.scss';

export default function CouponCenterPage() {
  const [coupons, setCoupons] = useState<any[]>([]);

  const fetchList = useCallback(async () => {
    try {
      const res = await getCouponCenter();
      setCoupons(res?.data || []);
    } catch (e) {
      console.error(e);
    }
  }, []);

  useLoad(() => { fetchList(); });

  const handleReceive = async (id: string) => {
    try {
      await receiveCoupon(id);
      Taro.showToast({ title: '领取成功', icon: 'success' });
      setCoupons((prev) =>
        prev.map((c) => (c.id === id ? { ...c, received: true } : c))
      );
    } catch (e) {
      Taro.showToast({ title: '领取失败', icon: 'none' });
    }
  };

  return (
    <View className="coupon-center">
      {coupons.length === 0 && (
        <View className="coupon-center__empty">
          <Text>暂无可领取的优惠券</Text>
        </View>
      )}
      {coupons.map((c) => (
        <View key={c.id} className="coupon-center__card">
          <View className="coupon-center__left">
            <Text className="coupon-center__symbol">¥</Text>
            <Text className="coupon-center__value">{(c.amount / 100).toFixed(0)}</Text>
          </View>
          <View className="coupon-center__right">
            <Text className="coupon-center__name">{c.name}</Text>
            <Text className="coupon-center__condition">
              {c.minAmount > 0 ? `满${(c.minAmount / 100).toFixed(0)}元可用` : '无门槛'}
            </Text>
            <Text className="coupon-center__time">{c.expireTime}</Text>
          </View>
          <View
            className={`coupon-center__btn ${c.received ? 'coupon-center__btn--disabled' : ''}`}
            onClick={() => !c.received && handleReceive(c.id)}
          >
            <Text>{c.received ? '已领取' : '领取'}</Text>
          </View>
        </View>
      ))}
    </View>
  );
}
