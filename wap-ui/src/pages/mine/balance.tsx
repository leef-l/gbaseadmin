import { useState, useCallback } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, useReachBottom } from '@tarojs/taro';
import { getBalanceLog } from '../../api/member';
import { useAuthStore } from '../../store/auth';
import './balance.scss';

export default function BalancePage() {
  const { userInfo } = useAuthStore();
  const [logs, setLogs] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);
  const [loading, setLoading] = useState(false);

  const fetchLogs = useCallback(async (pageNum: number, reset = false) => {
    if (loading) return;
    setLoading(true);
    try {
      const res = await getBalanceLog({ page: pageNum, pageSize: 20 });
      const list = res?.data?.list || [];
      setLogs((prev) => reset ? list : [...prev, ...list]);
      setHasMore(list.length >= 20);
      setPage(pageNum);
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  }, [loading]);

  useLoad(() => { fetchLogs(1, true); });

  useReachBottom(() => {
    if (hasMore && !loading) fetchLogs(page + 1);
  });

  return (
    <View className="balance">
      <View className="balance__header">
        <Text className="balance__label">当前余额（元）</Text>
        <Text className="balance__amount">
          {((userInfo?.balance || 0) / 100).toFixed(2)}
        </Text>
        <View
          className="balance__recharge-btn"
          onClick={() => Taro.navigateTo({ url: '/pages/recharge/index' })}
        >
          <Text>充值</Text>
        </View>
      </View>

      <View className="balance__list">
        <Text className="balance__list-title">余额明细</Text>
        {logs.length === 0 && !loading && (
          <View className="balance__empty"><Text>暂无记录</Text></View>
        )}
        {logs.map((item, i) => (
          <View key={i} className="balance__item">
            <View className="balance__item-left">
              <Text className="balance__item-type">{item.typeName || item.type}</Text>
              <Text className="balance__item-time">{item.createdAt}</Text>
            </View>
            <View className="balance__item-right">
              <Text className={`balance__item-amount ${item.amount > 0 ? 'balance__item-amount--in' : 'balance__item-amount--out'}`}>
                {item.amount > 0 ? '+' : ''}{(item.amount / 100).toFixed(2)}
              </Text>
              <Text className="balance__item-balance">余额 ¥{(item.balance / 100).toFixed(2)}</Text>
            </View>
          </View>
        ))}
        {loading && <View className="balance__loading"><Text>加载中...</Text></View>}
        {!hasMore && logs.length > 0 && <View className="balance__nomore"><Text>没有更多了</Text></View>}
      </View>
    </View>
  );
}
