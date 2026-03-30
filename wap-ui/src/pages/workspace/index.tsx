import { useState, useCallback } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh } from '@tarojs/taro';
import { useAuthStore } from '../../store/auth';
import { getIncome, getCoachOrders } from '../../api/coach';
import './index.scss';

const shortcuts = [
  { icon: '📋', text: '接单管理', url: '/pages/workspace/orders' },
  { icon: '💰', text: '收入统计', url: '/pages/workspace/income' },
  { icon: '🛍️', text: '商品管理', url: '/pages/workspace/goods' },
];

export default function WorkspacePage() {
  const { userInfo } = useAuthStore();
  const [stats, setStats] = useState({ todayOrders: 0, todayIncome: 0, totalIncome: 0, balance: 0 });
  const [recentOrders, setRecentOrders] = useState<any[]>([]);

  const fetchData = useCallback(async () => {
    try {
      const info = await getIncome();
      setStats({
        todayOrders: info.todayOrders ?? 0,
        todayIncome: info.todayIncome ?? 0,
        totalIncome: info.totalIncome ?? 0,
        balance: info.balance ?? 0,
      });
    } catch { /* ignore */ }
    try {
      const res = await getCoachOrders({ page: 1, pageSize: 5 });
      setRecentOrders(res.list || []);
    } catch { /* ignore */ }
  }, []);

  useLoad(() => { fetchData(); });
  usePullDownRefresh(() => { fetchData().then(() => Taro.stopPullDownRefresh()); });

  return (
    <View className="workspace">
      <View className="workspace__header">
        <View className="workspace__profile">
          <View className="workspace__avatar">👤</View>
          <View className="workspace__info">
            <Text className="workspace__name">{userInfo?.nickname || '陪玩师'}</Text>
            <Text className="workspace__level">{userInfo?.levelName || 'Lv.1'}</Text>
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
          <Text className="workspace__stat-value">¥{(stats.balance / 100).toFixed(2)}</Text>
          <Text className="workspace__stat-label">总余额</Text>
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
            <View key={o.id} className="workspace__order-item card" onClick={() => Taro.navigateTo({ url: `/pages/order/detail?id=${o.id}` })}>
              <View className="workspace__order-top">
                <Text className="workspace__order-name">{o.goodsName || '服务订单'}</Text>
                <Text className="workspace__order-amount">¥{(o.payAmount / 100).toFixed(2)}</Text>
              </View>
              <View className="workspace__order-bottom">
                <Text className="workspace__order-user">{o.memberName || '用户'}</Text>
                <Text className="workspace__order-time">{o.createdAt || ''}</Text>
              </View>
            </View>
          ))
        )}
      </View>
    </View>
  );
}