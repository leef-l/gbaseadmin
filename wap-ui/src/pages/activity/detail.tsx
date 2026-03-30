import { useState, useCallback } from 'react';
import { View, Text, Image } from '@tarojs/components';
import Taro, { useLoad, useRouter } from '@tarojs/taro';
import { getActivityDetail, joinActivity, completeStep } from '../../api/activity';
import './detail.scss';

export default function ActivityDetailPage() {
  const router = useRouter();
  const [detail, setDetail] = useState<any>(null);
  const [joining, setJoining] = useState(false);

  const fetchDetail = useCallback(async (id: string) => {
    try {
      const res = await getActivityDetail(id);
      setDetail(res || null);
    } catch (e) {
      console.error(e);
    }
  }, []);

  useLoad(() => {
    const id = router.params.id;
    if (id) fetchDetail(id);
  });

  const handleJoin = async () => {
    if (!detail || joining) return;
    setJoining(true);
    try {
      await joinActivity(detail.activityId);
      Taro.showToast({ title: '参与成功', icon: 'success' });
      fetchDetail(detail.activityId);
    } catch (e) {
      Taro.showToast({ title: '参与失败', icon: 'none' });
    } finally {
      setJoining(false);
    }
  };

  const handleCompleteStep = async (stepId: string) => {
    try {
      await completeStep({ activityId: detail.activityId, stepId });
      Taro.showToast({ title: '步骤完成', icon: 'success' });
      fetchDetail(detail.activityId);
    } catch (e) {
      Taro.showToast({ title: '操作失败', icon: 'none' });
    }
  };

  if (!detail) return <View className="activity-detail__loading"><Text>加载中...</Text></View>;

  return (
    <View className="activity-detail">
      {detail.cover && (
        <Image className="activity-detail__cover" src={detail.cover} mode="aspectFill" />
      )}

      <View className="activity-detail__info">
        <Text className="activity-detail__title">{detail.title}</Text>
        <Text className="activity-detail__desc">{detail.content}</Text>

        <View className="activity-detail__meta">
          <View className="activity-detail__meta-item">
            <Text className="activity-detail__meta-label">活动时间</Text>
            <Text className="activity-detail__meta-value">{detail.startTime} ~ {detail.endTime}</Text>
          </View>
          <View className="activity-detail__meta-item">
            <Text className="activity-detail__meta-label">参与人数</Text>
            <Text className="activity-detail__meta-value">{detail.joinCount || 0}人</Text>
          </View>
        </View>
      </View>

      {detail.rewards && detail.rewards.length > 0 && (
        <View className="activity-detail__section">
          <Text className="activity-detail__section-title">活动奖励</Text>
          {detail.rewards.map((r: any) => (
            <View key={r.rewardId} className="activity-detail__reward">
              <Text className="activity-detail__reward-name">{r.rewardName}</Text>
              <Text className="activity-detail__reward-value">{r.rewardValue}</Text>
            </View>
          ))}
        </View>
      )}

      {detail.steps && detail.steps.length > 0 && (
        <View className="activity-detail__section">
          <Text className="activity-detail__section-title">活动步骤</Text>
          {detail.steps.map((s: any, i: number) => (
            <View key={s.stepId} className="activity-detail__step">
              <View className="activity-detail__step-num">
                {s.stepNo || i + 1}
              </View>
              <View className="activity-detail__step-content">
                <Text className="activity-detail__step-title">{s.title}</Text>
                {s.description && <Text className="activity-detail__step-desc">{s.description}</Text>}
              </View>
              {detail.joined && (
                <View className="activity-detail__step-btn" onClick={() => handleCompleteStep(s.stepId)}>
                  <Text>完成</Text>
                </View>
              )}
            </View>
          ))}
        </View>
      )}

      <View className="activity-detail__bottom safe-bottom">
        <View
          className={`activity-detail__btn ${detail.joined || joining ? 'activity-detail__btn--disabled' : ''}`}
          onClick={handleJoin}
        >
          <Text>{detail.joined ? '已参与' : joining ? '处理中...' : '立即参与'}</Text>
        </View>
      </View>
    </View>
  );
}
