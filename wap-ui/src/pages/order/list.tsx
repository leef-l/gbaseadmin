import { useState, useCallback } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, useRouter, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import { getOrderList, cancelOrder } from '../../api/order';
import OrderCard from '../../components/OrderCard';
import EmptyState from '../../components/EmptyState';
import LoadMore from '../../components/LoadMore';
import './list.scss';

const tabs = [
  { text: '全部', status: -1 },
  { text: '待支付', status: 0 },
  { text: '已支付', status: 1 },
  { text: '进行中', status: 2 },
  { text: '已完成', status: 3 },
  { text: '退款', status: 5 },
];

export default function OrderListPage() {
  const { params } = useRouter();
  const initTab = params.status ? tabs.findIndex(t => t.status === Number(params.status)) : 0;
  const [activeTab, setActiveTab] = useState(initTab >= 0 ? initTab : 0);
  const [list, setList] = useState<any[]>([]);
  const [hasMore, setHasMore] = useState(false);
  const [page, setPage] = useState(1);

  const fetchList = useCallback(async (p: number, status: number, reset = false) => {
    try {
      const reqParams: any = { page: p, pageSize: 10 };
      if (status >= 0) reqParams.status = status;
      const data = await getOrderList(reqParams);
      const items = data?.list || [];
      setList(prev => reset ? items : [...prev, ...items]);
      setHasMore(items.length >= 10);
      setPage(p);
    } catch {}
  }, []);

  useLoad(() => {
    fetchList(1, tabs[activeTab].status, true);
  });

  const switchTab = (i: number) => {
    setActiveTab(i);
    setList([]);
    fetchList(1, tabs[i].status, true);
  };

  usePullDownRefresh(() => {
    fetchList(1, tabs[activeTab].status, true).then(() => Taro.stopPullDownRefresh());
  });

  useReachBottom(() => {
    if (hasMore) fetchList(page + 1, tabs[activeTab].status);
  });

  const handleCancel = async (id: string) => {
    try {
      await cancelOrder(id);
      Taro.showToast({ title: '已取消', icon: 'success' });
      fetchList(1, tabs[activeTab].status, true);
    } catch {}
  };

  const getActions = (status: number, id: string) => {
    switch (status) {
      case 0: return [
        { label: '取消订单', onClick: () => handleCancel(id) },
        { label: '去支付', type: 'primary' as const, onClick: () => Taro.navigateTo({ url: `/pages/order/pay?orderId=${id}` }) },
      ];
      case 3: return [
        { label: '再来一单', onClick: () => Taro.navigateTo({ url: `/pages/order/detail?id=${id}` }) },
        { label: '去评价', type: 'primary' as const, onClick: () => Taro.navigateTo({ url: `/pages/order/review?orderId=${id}` }) },
      ];
      default: return [];
    }
  };

  return (
    <View className="order-list">
      <View className="order-list__tabs">
        {tabs.map((t, i) => (
          <Text key={i} className={`order-list__tab ${activeTab === i ? 'order-list__tab--active' : ''}`} onClick={() => switchTab(i)}>{t.text}</Text>
        ))}
      </View>
      <View className="order-list__content">
        {list.length === 0 ? <EmptyState text="暂无订单" /> : list.map((o) => (
          <OrderCard key={o.orderId} {...o} actions={getActions(o.orderStatus, o.orderId)} />
        ))}
        {list.length > 0 && <LoadMore hasMore={hasMore} />}
      </View>
    </View>
  );
}
