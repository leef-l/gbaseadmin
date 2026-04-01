import { useState, useRef, useCallback, useEffect } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import { getMyCoupons } from '../../api/coupon';
import LoadMore from '../../components/LoadMore';
import './list.scss';

const PAGE_SIZE = 10;

const tabs = [
  { label: '未使用', value: 0 },
  { label: '已使用', value: 1 },
  { label: '已过期', value: 2 },
];

export default function CouponListPage() {
  const [activeTab, setActiveTab] = useState(0);
  const [coupons, setCoupons] = useState<any[]>([]);
  const [hasMore, setHasMore] = useState(false);
  const [loading, setLoading] = useState(false);
  const pageRef = useRef(1);

  const fetchList = useCallback(async (status: number, reset = false) => {
    if (loading) return;
    if (reset) pageRef.current = 1;
    setLoading(true);
    try {
      const res = await getMyCoupons({ status, page: pageRef.current, pageSize: PAGE_SIZE });
      const rows = res?.list || [];
      if (reset) {
        setCoupons(rows);
      } else {
        setCoupons((prev) => [...prev, ...rows]);
      }
      setHasMore(rows.length >= PAGE_SIZE);
      pageRef.current += 1;
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  }, [loading]);

  useLoad(() => { fetchList(0, true); });

  usePullDownRefresh(async () => {
    await fetchList(activeTab, true);
    Taro.stopPullDownRefresh();
  });

  useReachBottom(() => {
    if (hasMore && !loading) fetchList(activeTab);
  });

  // tab 切换时重置列表
  useEffect(() => {
    fetchList(activeTab, true);
  }, [activeTab]);

  const handleTabChange = (val: number) => {
    if (activeTab !== val) setActiveTab(val);
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
        {coupons.length === 0 && !loading && (
          <View className="coupon-list__empty">
            <Text>暂无优惠券</Text>
          </View>
        )}
        {coupons.map((c) => (
          <View key={c.couponMemberId} className={`coupon-card ${activeTab !== 0 ? 'coupon-card--disabled' : ''}`}>
            <View className="coupon-card__left">
              <Text className="coupon-card__symbol">¥</Text>
              <Text className="coupon-card__value">{(c.faceValue / 100).toFixed(0)}</Text>
            </View>
            <View className="coupon-card__right">
              <Text className="coupon-card__name">{c.title}</Text>
              <Text className="coupon-card__condition">
                {c.minAmount > 0 ? `满${(c.minAmount / 100).toFixed(0)}元可用` : '无门槛'}
              </Text>
              <Text className="coupon-card__time">{c.endTime}</Text>
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
        {coupons.length > 0 && <LoadMore hasMore={hasMore} />}
      </View>
    </View>
  );
}
