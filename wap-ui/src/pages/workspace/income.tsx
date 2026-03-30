import { useState, useCallback } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import { getCoachInfo, getProfitLogList } from '../../api/coach';
import EmptyState from '../../components/EmptyState';
import LoadMore from '../../components/LoadMore';
import './income.scss';

export default function IncomePage() {
  const [summary, setSummary] = useState({ totalIncome: 0, balance: 0, monthIncome: 0 });
  const [list, setList] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(false);
  const [loading, setLoading] = useState(false);

  const fetchSummary = useCallback(async () => {
    try {
      const info = await getCoachInfo();
      setSummary({
        totalIncome: info.totalIncome ?? 0,
        balance: info.balance ?? 0,
        monthIncome: info.monthIncome ?? 0,
      });
    } catch { /* ignore */ }
  }, []);

  const fetchList = useCallback(async (p = 1) => {
    setLoading(true);
    try {
      const res = await getProfitLogList({ page: p, pageSize: 15 });
      const items = res.list || [];
      setList(p === 1 ? items : (prev) => [...prev, ...items]);
      setHasMore(items.length >= 15);
      setPage(p);
    } catch { /* ignore */ }
    setLoading(false);
  }, []);

  useLoad(() => { fetchSummary(); fetchList(1); });
  usePullDownRefresh(() => {
    Promise.all([fetchSummary(), fetchList(1)]).then(() => Taro.stopPullDownRefresh());
  });
  useReachBottom(() => { if (hasMore && !loading) fetchList(page + 1); });

  return (
    <View className="ws-income">
      <View className="ws-income__header">
        <View className="ws-income__summary">
          <View className="ws-income__summary-item">
            <Text className="ws-income__summary-value">¥{(summary.totalIncome / 100).toFixed(2)}</Text>
            <Text className="ws-income__summary-label">总收入</Text>
          </View>
          <View className="ws-income__summary-item">
            <Text className="ws-income__summary-value">¥{(summary.balance / 100).toFixed(2)}</Text>
            <Text className="ws-income__summary-label">可提现余额</Text>
          </View>
        </View>
        <View className="ws-income__withdraw" onClick={() => Taro.showToast({ title: '提现功能即将开放', icon: 'none' })}>提现</View>
      </View>

      <View className="ws-income__month card">
        <Text className="ws-income__month-label">本月收入</Text>
        <Text className="ws-income__month-value">¥{(summary.monthIncome / 100).toFixed(2)}</Text>
      </View>

      <View className="ws-income__list">
        <Text className="ws-income__list-title">收入明细</Text>
        {list.length === 0 && !loading ? <EmptyState text="暂无收入记录" /> : (
          list.map((item, i) => (
            <View key={i} className="ws-income__item card">
              <View className="ws-income__item-left">
                <Text className="ws-income__item-order">订单 {item.orderNo || ''}</Text>
                <Text className="ws-income__item-time">{item.createdAt || ''}</Text>
              </View>
              <Text className="ws-income__item-amount">+¥{(item.amount / 100).toFixed(2)}</Text>
            </View>
          ))
        )}
        {list.length > 0 && <LoadMore hasMore={hasMore} loading={loading} />}
      </View>
    </View>
  );
}
