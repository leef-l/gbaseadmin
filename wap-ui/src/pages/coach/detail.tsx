import { useState, useRef, useCallback, useEffect } from 'react';
import { View, Text, Image } from '@tarojs/components';
import Taro, { useLoad, useRouter, useReachBottom } from '@tarojs/taro';
import { getCoachDetail } from '../../api/coach';
import { getGoodsList } from '../../api/goods';
import { getReviewList } from '../../api/review';
import LoadMore from '../../components/LoadMore';
import './detail.scss';

const PAGE_SIZE = 10;

export default function CoachDetailPage() {
  const { params } = useRouter();
  const [activeTab, setActiveTab] = useState(0);
  const [detail, setDetail] = useState<any>(null);

  // 服务项目分页
  const [goodsList, setGoodsList] = useState<any[]>([]);
  const [goodsHasMore, setGoodsHasMore] = useState(false);
  const [goodsLoading, setGoodsLoading] = useState(false);
  const goodsPageRef = useRef(1);

  // 评价分页
  const [reviews, setReviews] = useState<any[]>([]);
  const [reviewHasMore, setReviewHasMore] = useState(false);
  const [reviewLoading, setReviewLoading] = useState(false);
  const reviewPageRef = useRef(1);

  const fetchGoods = useCallback(async (reset = false) => {
    if (goodsLoading) return;
    if (reset) goodsPageRef.current = 1;
    setGoodsLoading(true);
    try {
      const data = await getGoodsList({ coachId: params.id, page: goodsPageRef.current, pageSize: PAGE_SIZE });
      const rows = data?.list || [];
      if (reset) {
        setGoodsList(rows);
      } else {
        setGoodsList((prev) => [...prev, ...rows]);
      }
      setGoodsHasMore(rows.length >= PAGE_SIZE);
      goodsPageRef.current += 1;
    } catch {
      Taro.showToast({ title: '加载失败', icon: 'none' });
    } finally {
      setGoodsLoading(false);
    }
  }, [params.id, goodsLoading]);

  const fetchReviews = useCallback(async (reset = false) => {
    if (reviewLoading) return;
    if (reset) reviewPageRef.current = 1;
    setReviewLoading(true);
    try {
      const data = await getReviewList({ coachId: params.id, page: reviewPageRef.current, pageSize: PAGE_SIZE });
      const rows = data?.list || [];
      if (reset) {
        setReviews(rows);
      } else {
        setReviews((prev) => [...prev, ...rows]);
      }
      setReviewHasMore(rows.length >= PAGE_SIZE);
      reviewPageRef.current += 1;
    } catch {} finally {
      setReviewLoading(false);
    }
  }, [params.id, reviewLoading]);

  useLoad(async () => {
    try {
      const data = await getCoachDetail(params.id || '');
      setDetail(data);
    } catch {
      Taro.showToast({ title: '加载失败', icon: 'none' });
    }
  });

  // tab 切换时重置并加载对应数据
  useEffect(() => {
    if (activeTab === 0) {
      fetchGoods(true);
    } else {
      fetchReviews(true);
    }
  }, [activeTab]);

  useReachBottom(() => {
    if (activeTab === 0) {
      if (goodsHasMore && !goodsLoading) fetchGoods();
    } else {
      if (reviewHasMore && !reviewLoading) fetchReviews();
    }
  });

  if (!detail) return <View />;

  return (
    <View className="coach-detail">
      <View className="coach-detail__cover">
        {detail.avatar && (
          <Image src={detail.avatar} className="coach-detail__cover-img" mode="aspectFill" />
        )}
        <View className="coach-detail__cover-mask" />
        <View className="coach-detail__back" onClick={() => Taro.navigateBack()}>←</View>
        <View className="coach-detail__cover-info">
          {detail.isOnline && <View className="coach-detail__online-dot" />}
          <Text className="coach-detail__cover-name">{detail.nickname}</Text>
        </View>
      </View>

      {/* 信息卡片 */}
      <View className="coach-detail__info-card card">
        <View className="coach-detail__score-row">
          <Text className="coach-detail__score-num">{detail.score}</Text>
          <Text className="coach-detail__score-stars">★★★★★</Text>
        </View>
        <View className="coach-detail__stats">
          <Text className="coach-detail__stat">接单 <Text className="coach-detail__stat-val">{detail.orderCount}</Text></Text>
        </View>
        <Text className="coach-detail__intro">{detail.intro}</Text>
      </View>

      {/* Tab 切换 */}
      <View className="coach-detail__tabs">
        <Text className={`coach-detail__tab ${activeTab === 0 ? 'coach-detail__tab--active' : ''}`} onClick={() => setActiveTab(0)}>服务项目</Text>
        <Text className={`coach-detail__tab ${activeTab === 1 ? 'coach-detail__tab--active' : ''}`} onClick={() => setActiveTab(1)}>用户评价</Text>
      </View>

      {activeTab === 0 ? (
        <View className="coach-detail__goods-list">
          {goodsList.map((g) => (
            <View key={g.goodsId} className="coach-detail__goods-item card">
              <View className="coach-detail__goods-cover" />
              <View className="coach-detail__goods-info">
                <Text className="coach-detail__goods-name">{g.title}</Text>
                <Text className="coach-detail__goods-desc">{g.description}</Text>
                <View className="coach-detail__goods-bottom">
                  <Text className="price">¥{(g.price / 100).toFixed(2)}</Text>
                  <View className="coach-detail__goods-btn" onClick={() => Taro.navigateTo({ url: `/pages/goods/detail?id=${g.goodsId}` })}>立即预约</View>
                </View>
              </View>
            </View>
          ))}
          {goodsList.length > 0 && <LoadMore hasMore={goodsHasMore} />}
        </View>
      ) : (
        <View className="coach-detail__reviews">
          {reviews.map((r) => (
            <View key={r.reviewId} className="coach-detail__review-item">
              <View className="coach-detail__review-header">
                <View className="coach-detail__review-avatar" />
                <Text className="coach-detail__review-name">{r.nickname}</Text>
                <Text className="coach-detail__review-score">{'★'.repeat(r.score)}</Text>
                <Text className="coach-detail__review-time">{r.createdAt}</Text>
              </View>
              <Text className="coach-detail__review-content">{r.content}</Text>
            </View>
          ))}
          {reviews.length > 0 && <LoadMore hasMore={reviewHasMore} />}
        </View>
      )}

      {/* 底部操作栏 */}
      <View className="bottom-bar">
        <View style={{ display: 'flex', gap: '16px' }}>
          <Text style={{ fontSize: '20px' }}>❤️</Text>
          <Text style={{ fontSize: '20px' }}>💬</Text>
        </View>
        <View className="btn-primary" onClick={() => {
          if (goodsList.length > 0) {
            Taro.navigateTo({ url: `/pages/goods/detail?id=${goodsList[0].goodsId}` });
          } else {
            Taro.showToast({ title: '暂无可预约的服务', icon: 'none' });
          }
        }}>立即预约</View>
      </View>
    </View>
  );
}
