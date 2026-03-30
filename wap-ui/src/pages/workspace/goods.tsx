import { useState, useCallback } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import { getMyGoodsList, updateGoodsStatus } from '../../api/coach';
import EmptyState from '../../components/EmptyState';
import LoadMore from '../../components/LoadMore';
import './goods.scss';

export default function WorkspaceGoodsPage() {
  const [list, setList] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(false);
  const [loading, setLoading] = useState(false);

  const fetchList = useCallback(async (p = 1) => {
    setLoading(true);
    try {
      const res = await getMyGoodsList({ page: p, pageSize: 10 });
      const items = res.list || [];
      setList(p === 1 ? items : (prev) => [...prev, ...items]);
      setHasMore(items.length >= 10);
      setPage(p);
    } catch { /* ignore */ }
    setLoading(false);
  }, []);

  useLoad(() => { fetchList(1); });
  usePullDownRefresh(() => { fetchList(1).then(() => Taro.stopPullDownRefresh()); });
  useReachBottom(() => { if (hasMore && !loading) fetchList(page + 1); });

  const toggleStatus = async (id: string, currentStatus: number) => {
    const newStatus = currentStatus === 1 ? 0 : 1;
    try {
      await updateGoodsStatus({ goodsId: id, status: newStatus });
      Taro.showToast({ title: newStatus === 1 ? '已上架' : '已下架', icon: 'success' });
      setList((prev) => prev.map((g) => g.goodsId === id ? { ...g, status: newStatus } : g));
    } catch { /* ignore */ }
  };

  return (
    <View className="ws-goods">
      <View className="ws-goods__header">
        <Text className="ws-goods__title">我的商品</Text>
        <View className="ws-goods__publish" onClick={() => Taro.showToast({ title: '发布功能即将开放', icon: 'none' })}>+ 发布商品</View>
      </View>
      <View className="ws-goods__content">
        {list.length === 0 && !loading ? <EmptyState text="暂无商品" /> : (
          list.map((g) => (
            <View key={g.goodsId} className="ws-goods__card card">
              <View className="ws-goods__card-cover">
                <View className="ws-goods__card-img ws-goods__card-img--empty">🛍️</View>
              </View>
              <View className="ws-goods__card-info">
                <Text className="ws-goods__card-name">{g.title || '商品'}</Text>
                <Text className="ws-goods__card-price">¥{(g.price / 100).toFixed(2)}</Text>
                <View className="ws-goods__card-bottom">
                  <Text className={`ws-goods__card-status ${g.status === 1 ? 'ws-goods__card-status--on' : ''}`}>
                    {g.status === 1 ? '已上架' : '已下架'}
                  </Text>
                  <View className="ws-goods__card-toggle" onClick={() => toggleStatus(g.goodsId, g.status)}>
                    {g.status === 1 ? '下架' : '上架'}
                  </View>
                </View>
              </View>
            </View>
          ))
        )}
        {list.length > 0 && <LoadMore hasMore={hasMore} loading={loading} />}
      </View>
    </View>
  );
}
