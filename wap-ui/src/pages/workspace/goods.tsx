import { useState, useCallback } from 'react';
import { View, Text, Input } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import { getMyGoodsList, updateGoodsStatus, createGoods } from '../../api/coach';
import EmptyState from '../../components/EmptyState';
import LoadMore from '../../components/LoadMore';
import './goods.scss';

const CATEGORIES = ['王者荣耀', '英雄联盟', '和平精英', '原神', '其他'];

export default function WorkspaceGoodsPage() {
  const [list, setList] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(false);
  const [loading, setLoading] = useState(false);

  // 发布商品弹窗状态
  const [showPublish, setShowPublish] = useState(false);
  const [publishing, setPublishing] = useState(false);
  const [form, setForm] = useState({ title: '', price: '', category: '', description: '' });

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

  const openPublish = () => {
    setForm({ title: '', price: '', category: '', description: '' });
    setShowPublish(true);
  };

  const handlePublish = async () => {
    if (!form.title.trim()) {
      Taro.showToast({ title: '请填写商品名称', icon: 'none' });
      return;
    }
    const priceNum = parseFloat(form.price);
    if (!form.price || isNaN(priceNum) || priceNum <= 0) {
      Taro.showToast({ title: '请填写正确的价格', icon: 'none' });
      return;
    }
    if (!form.category) {
      Taro.showToast({ title: '请选择分类', icon: 'none' });
      return;
    }
    setPublishing(true);
    try {
      await createGoods({
        title: form.title.trim(),
        price: Math.round(priceNum * 100),
        category: form.category,
        description: form.description.trim(),
      });
      Taro.showToast({ title: '发布成功', icon: 'success' });
      setShowPublish(false);
      fetchList(1);
    } catch { /* ignore */ }
    setPublishing(false);
  };

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
        <View className="ws-goods__publish" onClick={openPublish}>+ 发布商品</View>
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

      {/* 发布商品弹窗 */}
      {showPublish && (
        <View className="ws-goods__modal-mask" onClick={() => setShowPublish(false)}>
          <View className="ws-goods__modal" onClick={(e) => e.stopPropagation()}>
            <View className="ws-goods__modal-header">
              <Text className="ws-goods__modal-title">发布商品</Text>
              <Text className="ws-goods__modal-close" onClick={() => setShowPublish(false)}>✕</Text>
            </View>
            <View className="ws-goods__modal-body">
              <View className="ws-goods__field">
                <Text className="ws-goods__field-label">商品名称</Text>
                <Input
                  className="ws-goods__field-input"
                  placeholder="请输入商品名称"
                  value={form.title}
                  onInput={(e) => setForm((f) => ({ ...f, title: e.detail.value }))}
                />
              </View>
              <View className="ws-goods__field">
                <Text className="ws-goods__field-label">价格（元）</Text>
                <Input
                  className="ws-goods__field-input"
                  placeholder="如：88.00"
                  type="digit"
                  value={form.price}
                  onInput={(e) => setForm((f) => ({ ...f, price: e.detail.value }))}
                />
              </View>
              <View className="ws-goods__field">
                <Text className="ws-goods__field-label">游戏分类</Text>
                <View className="ws-goods__cats">
                  {CATEGORIES.map((c) => (
                    <Text
                      key={c}
                      className={`ws-goods__cat ${form.category === c ? 'ws-goods__cat--active' : ''}`}
                      onClick={() => setForm((f) => ({ ...f, category: c }))}
                    >{c}</Text>
                  ))}
                </View>
              </View>
              <View className="ws-goods__field">
                <Text className="ws-goods__field-label">简介（选填）</Text>
                <Input
                  className="ws-goods__field-input"
                  placeholder="简单介绍一下你的服务"
                  value={form.description}
                  onInput={(e) => setForm((f) => ({ ...f, description: e.detail.value }))}
                />
              </View>
            </View>
            <View
              className={`ws-goods__modal-submit ${publishing ? 'ws-goods__modal-submit--loading' : ''}`}
              onClick={publishing ? undefined : handlePublish}
            >
              {publishing ? '发布中...' : '确认发布'}
            </View>
          </View>
        </View>
      )}
    </View>
  );
}
