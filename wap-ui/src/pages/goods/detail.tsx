import { useState } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, useRouter } from '@tarojs/taro';
import { getGoodsDetail } from '../../api/goods';
import { getReviewList } from '../../api/review';
import { useCartStore } from '../../store/cart';
import './detail.scss';

export default function GoodsDetailPage() {
  const { params } = useRouter();
  const [detail, setDetail] = useState<any>(null);
  const [reviews, setReviews] = useState<any[]>([]);
  const setOrder = useCartStore((s) => s.setOrder);

  useLoad(async () => {
    try {
      const data = await getGoodsDetail(params.id || '');
      setDetail(data);
    } catch {
      Taro.showToast({ title: '加载失败', icon: 'none' });
    }
    try {
      const data = await getReviewList({ goodsId: params.id, page: 1, pageSize: 5 });
      setReviews(data?.list || []);
    } catch {}
  });

  const handleOrder = () => {
    if (!detail) return;
    setOrder({
      goodsId: detail.goodsId,
      goodsName: detail.title,
      coachId: detail.coachId,
      coachName: detail.coachName,
      price: detail.price,
      quantity: 1,
    });
    Taro.navigateTo({ url: '/pages/order/confirm' });
  };

  if (!detail) return <View />;

  return (
    <View className="goods-detail">
      <View className="goods-detail__cover" />
      <View className="goods-detail__info">
        <Text className="goods-detail__price">¥{(detail.price / 100).toFixed(2)}<Text className="goods-detail__unit"> /{detail.unit}</Text></Text>
        <Text className="goods-detail__name">{detail.title}</Text>
        <Text className="goods-detail__sales">已售 {detail.salesNum} 单</Text>
      </View>

      <View className="goods-detail__coach card" onClick={() => Taro.navigateTo({ url: `/pages/coach/detail?id=${detail.coachId}` })}>
        <View className="goods-detail__coach-avatar" />
        <View className="goods-detail__coach-info">
          <Text className="goods-detail__coach-name">{detail.coachName}</Text>
          <Text className="goods-detail__coach-meta">★ {detail.coachScore}</Text>
        </View>
        <Text className="goods-detail__coach-link">查看主页 &gt;</Text>
      </View>

      <View className="goods-detail__desc">
        <Text className="goods-detail__desc-title">商品详情</Text>
        <Text className="goods-detail__desc-content">{detail.description}</Text>
      </View>

      <View className="goods-detail__reviews">
        <View className="section-header">
          <Text className="title">用户评价</Text>
          <Text className="more">查看全部 &gt;</Text>
        </View>
        {reviews.map((r) => (
          <View key={r.reviewId} className="goods-detail__review-item">
            <View className="goods-detail__review-header">
              <View className="goods-detail__review-avatar" />
              <Text className="goods-detail__review-name">{r.nickname}</Text>
              <Text className="goods-detail__review-score">{'★'.repeat(r.score)}</Text>
            </View>
            <Text className="goods-detail__review-content">{r.content}</Text>
          </View>
        ))}
      </View>

      <View className="bottom-bar">
        <View style={{ display: 'flex', gap: '16px' }}>
          <Text style={{ fontSize: '20px' }}>💬</Text>
          <Text style={{ fontSize: '20px' }}>❤️</Text>
        </View>
        <View className="btn-primary" onClick={handleOrder}>立即下单</View>
      </View>
    </View>
  );
}
