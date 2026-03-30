import { useState, useCallback } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import { getCoachOrders } from '../../api/coach';
import { acceptOrder, finishOrder } from '../../api/order';
import EmptyState from '../../components/EmptyState';
import LoadMore from '../../components/LoadMore';
import './orders.scss';

const tabs = [
  { text: '待接单', status: 1 },
  { text: '进行中', status: 2 },
  { text: '已完成', status: 3 },
];

export default function WorkspaceOrdersPage() {
  const [activeTab, setActiveTab] = useState(0);
  const [list, setList] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(false);
  const [loading, setLoading] = useState(false);

  const fetchList = useCallback(async (p = 1, status?: number) => {
    setLoading(true);
    try {
      const s = status ?? tabs[activeTab].status;
      const res = await getCoachOrders({ page: p, pageSize: 10, status: s });
      const items = res.list || [];
      setList(p === 1 ? items : (prev) => [...prev, ...items]);
      setHasMore(items.length >= 10);
      setPage(p);
    } catch { /* ignore */ }
    setLoading(false);
  }, [activeTab]);

  useLoad(() => { fetchList(1); });
  usePullDownRefresh(() => { fetchList(1).then(() => Taro.stopPullDownRefresh()); });
  useReachBottom(() => { if (hasMore && !loading) fetchList(page + 1); });

  const switchTab = (idx: number) => {
    setActiveTab(idx);
    setList([]);
    fetchList(1, tabs[idx].status);
  };

  const handleAccept = async (id: string) => {
    try {
      await acceptOrder(id);
      Taro.showToast({ title: '接单成功', icon: 'success' });
      fetchList(1);
    } catch { /* ignore */ }
  };

  const handleFinish = async (id: string) => {
    try {
      await finishOrder(id);
      Taro.showToast({ title: '已完成服务', icon: 'success' });
      fetchList(1);
    } catch { /* ignore */ }
  };

  return (
    <View className="ws-orders">
      <View className="ws-orders__tabs">
        {tabs.map((t, i) => (
          <Text key={i} className={`ws-orders__tab ${activeTab === i ? 'ws-orders__tab--active' : ''}`} onClick={() => switchTab(i)}>{t.text}</Text>
        ))}
      </View>
      <View className="ws-orders__content">
        {list.length === 0 && !loading ? <EmptyState text="暂无订单" /> : (
          list.map((o) => (
            <View key={o.orderId} className="ws-orders__card card">
              <View className="ws-orders__card-top">
                <Text className="ws-orders__card-user">{o.coachName || '用户'}</Text>
                <Text className="ws-orders__card-time">{o.createdAt || ''}</Text>
              </View>
              <View className="ws-orders__card-body">
                <Text className="ws-orders__card-goods">{o.goodsTitle || '服务订单'}</Text>
                <Text className="ws-orders__card-amount">¥{(o.payAmount / 100).toFixed(2)}</Text>
              </View>
              {activeTab === 0 && (
                <View className="ws-orders__card-actions">
                  <View className="ws-orders__btn ws-orders__btn--primary" onClick={() => handleAccept(o.orderId)}>接单</View>
                </View>
              )}
              {activeTab === 1 && (
                <View className="ws-orders__card-actions">
                  <View className="ws-orders__btn ws-orders__btn--primary" onClick={() => handleFinish(o.orderId)}>完成服务</View>
                </View>
              )}
            </View>
          ))
        )}
        {list.length > 0 && <LoadMore hasMore={hasMore} loading={loading} />}
      </View>
    </View>
  );
}
