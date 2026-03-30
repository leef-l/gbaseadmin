import { useState, useCallback } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad } from '@tarojs/taro';
import { getCouponList } from '../../api/coupon';
import './list.scss';

const tabs = [
  { label: '未使用', value: 0 },
  { label: '已使用', value: 1 },
  { label: '已过期', value: 2 },
];

export default function CouponListPage() {
  const [activeTab, setActiveTab] = useState(0);
  const [coupons, setCoupons] = useState<any[]>([]);

  const fetchList = useCallback(async (status: number) => {
    try {
      const res = await getCouponList({ status });
      setCoupons(res?.data || []);
    } catch (e) {
      console.error(e);
    }
  }, []);

  useLoad(() => { fetchList(0); });

  const handleTabChange = (val: number) => {
    setActiveTab(val);
    fetchList(val);
  };

  return (
    <View className="coupon-list">
      <View className="coupon-list__tabs">
        {tabs.map((t) => (
          <View
            key={t.value}
            className={`coupon-list__tab ${activeTab === t.value ? 'coupon-list__tab--active' : ''}`}
            onClick={() => handleTabChange(t.value)}
          >
            <Text>{t.label}</Text>
          </View>
        ))}
      </View>

      <View className="coupon-list__body">
        {coupons.length === 0 && (
          <View className="coupon-list__empty">
            <Text>暂无优惠券</Text>
          </View>
        )}
        {coupons.map((c) => (
          <View key={c.id} className={`coupon-card ${activeTab !== 0 ? 'coupon-card--disabled' : ''}`}>
            <View className="coupon-card__left">
              <Text className="coupon-card__symbol">¥</Text>
              <Text className="coupon-card__value">{(c.amount / 100).toFixed(0)}</Text>
            </View>
            <View className="coupon-card__right">
              <Text className="coupon-card__name">{c.name}</Text>
              <Text className="coupon-card__condition">
                {c.minAmount > 0 ? `满${(c.minAmount / 100).toFixed(0)}元可用` : '无门槛'}
              </Text>
              <Text className="coupon-card__time">{c.expireTime}</Text>
              {activeTab === 0 && (
                <View
                  className="coupon-card__btn"
                  onClick={() => Taro.switchTab({ url: '/pages/index/index' })}
                >
                  <Text>去使用</Text>
                </View>
              )}
            </View>
          </View>
        ))}
      </View>
    </View>
  );
}
