import { useState, useRef, useCallback, useEffect } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import { getAvailableCoupons, receiveCoupon } from '../../api/coupon';
import LoadMore from '../../components/LoadMore';
import './center.scss';

const PAGE_SIZE = 10;

export default function CouponCenterPage() {
  const [coupons, setCoupons] = useState<any[]>([]);
  const [hasMore, setHasMore] = useState(false);
  const [loading, setLoading] = useState(false);
  const pageRef = useRef(1);

  const fetchList = useCallback(async (reset = false) => {
    if (loading) return;
    if (reset) pageRef.current = 1;
    setLoading(true);
    try {
      const res = await getAvailableCoupons({ page: pageRef.current, pageSize: PAGE_SIZE });
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

  useLoad(() => { fetchList(true); });

  usePullDownRefresh(async () => {
    await fetchList(true);
    Taro.stopPullDownRefresh();
  });

  useReachBottom(() => {
    if (hasMore && !loading) fetchList();
  });

  const handleReceive = async (couponId: string) => {
    try {
      await receiveCoupon(couponId);
      Taro.showToast({ title: '领取成功', icon: 'success' });
      setCoupons((prev) =>
        prev.map((c) => (c.couponId === couponId ? { ...c, received: true } : c))
      );
    } catch (e) {
      Taro.showToast({ title: '领取失败', icon: 'none' });
    }
  };

  return (
    <View className="coupon-center">
      {coupons.length === 0 && !loading && (
        <View className="coupon-center__empty">
          <Text>暂无可领取的优惠券</Text>
        </View>
      )}
      {coupons.map((c) => (
        <View key={c.couponId} className="coupon-center__card">
          <View className="coupon-center__left">
            <Text className="coupon-center__symbol">¥</Text>
            <Text className="coupon-center__value">{(c.faceValue / 100).toFixed(0)}</Text>
          </View>
          <View className="coupon-center__right">
            <Text className="coupon-center__name">{c.title}</Text>
            <Text className="coupon-center__condition">
              {c.minAmount > 0 ? `满${(c.minAmount / 100).toFixed(0)}元可用` : '无门槛'}
            </Text>
            <Text className="coupon-center__time">{c.endTime}</Text>
          </View>
          <View
            className={`coupon-center__btn ${c.received ? 'coupon-center__btn--disabled' : ''}`}
            onClick={() => !c.received && handleReceive(c.couponId)}
          >
            <Text>{c.received ? '已领取' : '领取'}</Text>
          </View>
        </View>
      ))}
      {coupons.length > 0 && <LoadMore hasMore={hasMore} />}
    </View>
  );
}
