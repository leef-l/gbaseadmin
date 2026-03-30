import { useState } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, useRouter } from '@tarojs/taro';
import { getCoachDetail } from '../../api/coach';
import { getGoodsList } from '../../api/goods';
import { getReviewList } from '../../api/review';
import './detail.scss';

export default function CoachDetailPage() {
  const { params } = useRouter();
  const [activeTab, setActiveTab] = useState(0);
  const [detail, setDetail] = useState<any>(null);
  const [goodsList, setGoodsList] = useState<any[]>([]);
  const [reviews, setReviews] = useState<any[]>([]);

  useLoad(async () => {
    try {
      const data = await getCoachDetail(params.id || '');
      setDetail(data);
    } catch {
      Taro.showToast({ title: '加载失败', icon: 'none' });
    }
    try {
      const data = await getGoodsList({ coachId: params.id, page: 1, pageSize: 20 });
      setGoodsList(data?.list || []);
    } catch {}
    try {
      const data = await getReviewList({ coachId: params.id, page: 1, pageSize: 10 });
      setReviews(data?.list || []);
    } catch {}
  });

  if (!detail) return <View />;

  return (
    <View className="coach-detail">
      <View className="coach-detail__cover">
        <View className="coach-detail__cover-mask" />
        <View className="coach-detail__back" onClick={() => Taro.navigateBack()}>←</View>
        <View className="coach-detail__cover-info">
          {detail.online && <View className="coach-detail__online-dot" />}
          <Text className="coach-detail__cover-name">{detail.nickname}</Text>
          <Text className="coach-detail__cover-level">{detail.level}</Text>
        </View>
      </View>

      {/* 信息卡片 */}
      <View className="coach-detail__info-card card">
        <View className="coach-detail__score-row">
          <Text className="coach-detail__score-num">{detail.score}</Text>
          <Text className="coach-detail__score-stars">★★★★★</Text>
          <Text className="coach-detail__score-count">{detail.scoreCount}人评分</Text>
        </View>
        <View className="coach-detail__stats">
          <Text className="coach-detail__stat">接单 <Text className="coach-detail__stat-val">{detail.orderCount}</Text></Text>
          <Text className="coach-detail__stat">好评率 <Text className="coach-detail__stat-val">{detail.goodRate}</Text></Text>
          <Text className="coach-detail__stat">响应 <Text className="coach-detail__stat-val">{detail.responseTime}</Text></Text>
        </View>
        <Text className="coach-detail__intro">{detail.intro}</Text>
        <View className="coach-detail__skill-tags">
          {detail.tags.map((t: string, i: number) => <Text key={i} className="tag">{t}</Text>)}
        </View>
      </View>

      {/* Tab 切换 */}
      <View className="coach-detail__tabs">
        <Text className={`coach-detail__tab ${activeTab === 0 ? 'coach-detail__tab--active' : ''}`} onClick={() => setActiveTab(0)}>服务项目</Text>
        <Text className={`coach-detail__tab ${activeTab === 1 ? 'coach-detail__tab--active' : ''}`} onClick={() => setActiveTab(1)}>用户评价</Text>
      </View>

      {activeTab === 0 ? (
        <View className="coach-detail__goods-list">
          {goodsList.map((g) => (
            <View key={g.id} className="coach-detail__goods-item card">
              <View className="coach-detail__goods-cover" />
              <View className="coach-detail__goods-info">
                <Text className="coach-detail__goods-name">{g.name}</Text>
                <Text className="coach-detail__goods-desc">{g.desc}</Text>
                <View className="coach-detail__goods-bottom">
                  <Text className="price">¥{(g.price / 100).toFixed(2)}</Text>
                  <View className="coach-detail__goods-btn" onClick={() => Taro.navigateTo({ url: `/pages/goods/detail?id=${g.id}` })}>立即预约</View>
                </View>
              </View>
            </View>
          ))}
        </View>
      ) : (
        <View className="coach-detail__reviews">
          {reviews.map((r) => (
            <View key={r.id} className="coach-detail__review-item">
              <View className="coach-detail__review-header">
                <View className="coach-detail__review-avatar" />
                <Text className="coach-detail__review-name">{r.name}</Text>
                <Text className="coach-detail__review-score">{'★'.repeat(r.score)}</Text>
                <Text className="coach-detail__review-time">{r.time}</Text>
              </View>
              <Text className="coach-detail__review-content">{r.content}</Text>
            </View>
          ))}
        </View>
      )}

      {/* 底部操作栏 */}
      <View className="bottom-bar">
        <View style={{ display: 'flex', gap: '16px' }}>
          <Text style={{ fontSize: '20px' }}>❤️</Text>
          <Text style={{ fontSize: '20px' }}>💬</Text>
        </View>
        <View className="btn-primary" onClick={() => Taro.navigateTo({ url: `/pages/goods/detail?coachId=${detail.id}` })}>立即预约</View>
      </View>
    </View>
  );
}