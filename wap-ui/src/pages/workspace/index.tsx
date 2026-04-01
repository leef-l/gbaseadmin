import { useState, useCallback } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh } from '@tarojs/taro';
import { useAuthStore } from '../../store/auth';
import { getIncome, getCoachOrders, setOnline } from '../../api/coach';
import './index.scss';

const shortcuts = [
  { icon: '📋', text: '接单管理', url: '/pages/workspace/orders' },
  { icon: '💰', text: '收入统计', url: '/pages/workspace/income' },
  { icon: '🛍️', text: '商品管理', url: '/pages/workspace/goods' },
  { icon: '⭐', text: '评价管理', url: '/pages/workspace/reviews' },
];

export default function WorkspacePage() {
  const { userInfo } = useAuthStore();
  const [stats, setStats] = useState({ todayOrders: 0, todayIncome: 0, totalIncome: 0, totalOrders: 0 });
  const [recentOrders, setRecentOrders] = useState<any[]>([]);
  const [isOnline, setIsOnline] = useState(false);
  const [onlineLoading, setOnlineLoading] = useState(false);

  const fetchData = useCallback(async () => {
    try {
      const info = await getIncome();
      setStats({
        todayOrders: info.todayOrders ?? 0,
        todayIncome: info.todayIncome ?? 0,
        totalIncome: info.totalIncome ?? 0,
        totalOrders: info.totalOrders ?? 0,
      });
    } catch { /* ignore */ }
    try {
      const res = await getCoachOrders({ page: 1, pageSize: 5 });
      setRecentOrders(res.list || []);
    } catch { /* ignore */ }
  }, []);

  const handleToggleOnline = useCallback(async () => {
    if (onlineLoading) return;
    setOnlineLoading(true);
    const next = isOnline ? 0 : 1;
    try {
      await setOnline(next);
      setIsOnline(next === 1);
      Taro.showToast({ title: next === 1 ? '已上线，开始接单' : '已下线', icon: 'success' });
    } catch { /* ignore */ }
    setOnlineLoading(false);
  }, [isOnline, onlineLoading]);

  useLoad(() => { fetchData(); });
  usePullDownRefresh(() => { fetchData().then(() => Taro.stopPullDownRefresh()); });

  return (
    <View className="workspace">
      <View className="workspace__header">
        <View className="workspace__profile">
          <View className="workspace__avatar">👤</View>
          <View className="workspace__info">
            <Text className="workspace__name">{userInfo?.nickname || '陪玩师'}</Text>
            <Text className="workspace__level">{userInfo?.levelTitle || 'Lv.1'}</Text>
          </View>
          <View
            className={`workspace__online-toggle ${isOnline ? 'workspace__online-toggle--on' : ''} ${onlineLoading ? 'workspace__online-toggle--loading' : ''}`}
            onClick={handleToggleOnline}
          >
            <View className="workspace__online-dot" />
            <Text className="workspace__online-text">{onlineLoading ? '...' : (isOnline ? '在线' : '离线')}</Text>
          </View>
        </View>
      </View>

      {/* 今日数据 */}
      <View className="workspace__stats card">
        <View className="workspace__stat">
          <Text className="workspace__stat-value">{stats.todayOrders}</Text>
          <Text className="workspace__stat-label">今日订单</Text>
        </View>
        <View className="workspace__stat">
          <Text className="workspace__stat-value">¥{(stats.todayIncome / 100).toFixed(2)}</Text>
          <Text className="workspace__stat-label">今日收入</Text>
        </View>
        <View className="workspace__stat">
          <Text className="workspace__stat-value">¥{(stats.totalIncome / 100).toFixed(2)}</Text>
          <Text className="workspace__stat-label">总收入</Text>
        </View>
        <View className="workspace__stat">
          <Text className="workspace__stat-value">{stats.totalOrders}</Text>
          <Text className="workspace__stat-label">总订单</Text>
        </View>
      </View>

      {/* 快捷入口 */}
      <View className="workspace__shortcuts">
        {shortcuts.map((s, i) => (
          <View key={i} className="workspace__shortcut card" onClick={() => Taro.navigateTo({ url: s.url })}>
            <Text className="workspace__shortcut-icon">{s.icon}</Text>
            <Text className="workspace__shortcut-text">{s.text}</Text>
          </View>
        ))}
      </View>

      {/* 最近订单 */}
      <View className="workspace__recent">
        <View className="workspace__recent-header">
          <Text className="workspace__recent-title">最近订单</Text>
          <Text className="workspace__recent-more" onClick={() => Taro.navigateTo({ url: '/pages/workspace/orders' })}>查看全部 &gt;</Text>
        </View>
        {recentOrders.length === 0 ? (
          <View className="workspace__empty">
            <Text className="workspace__empty-text">暂无订单</Text>
          </View>
        ) : (
          recentOrders.map((o) => (
            <View key={o.orderId} className="workspace__order-item card" onClick={() => Taro.navigateTo({ url: `/pages/order/detail?id=${o.orderId}` })}>
              <View className="workspace__order-top">
                <Text className="workspace__order-name">{o.goodsTitle || '服务订单'}</Text>
                <Text className="workspace__order-amount">¥{(o.payAmount / 100).toFixed(2)}</Text>
              </View>
              <View className="workspace__order-bottom">
                <Text className="workspace__order-user">{o.coachName || '用户'}</Text>
                <Text className="workspace__order-time">{o.createdAt || ''}</Text>
              </View>
            </View>
          ))
        )}
      </View>
    </View>
  );
}