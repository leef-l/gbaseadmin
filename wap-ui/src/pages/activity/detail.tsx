import { useState, useCallback } from 'react';
import { View, Text, Image, RichText } from '@tarojs/components';
import Taro, { useLoad, useRouter } from '@tarojs/taro';
import { getActivityDetail, joinActivity, completeStep } from '../../api/activity';
import './detail.scss';

const rewardTypeIcon: Record<number, string> = { 1: '💰', 2: '🎫', 3: '⭐', 4: '👑' };
const stepTypeLabel: Record<number, string> = { 1: '文字', 2: '链接', 3: '图片' };

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
    if (!detail || joining || detail.joined) return;
    setJoining(true);
    try {
      await joinActivity(detail.id);
      Taro.showToast({ title: '参与成功', icon: 'success' });
      fetchDetail(detail.id);
    } catch {
      Taro.showToast({ title: '参与失败', icon: 'none' });
    } finally {
      setJoining(false);
    }
  };

  const handleCompleteStep = async (stepId: string) => {
    try {
      await completeStep({ id: detail.id, stepId });
      Taro.showToast({ title: '步骤完成', icon: 'success' });
      fetchDetail(detail.id);
    } catch {
      Taro.showToast({ title: '操作失败', icon: 'none' });
    }
  };
  const handleCopy = (text: string) => {
    Taro.setClipboardData({ data: text });
  };

  const handleOpenLink = (url: string) => {
    Taro.setClipboardData({ data: url });
    Taro.showToast({ title: '链接已复制', icon: 'success' });
  };

  const isStepDone = (stepId: string) => detail.completedSteps?.includes(stepId);

  const isStepActive = (index: number) => {
    if (!detail.joined) return false;
    if (index === 0) return true;
    return isStepDone(detail.steps[index - 1]?.stepId);
  };

  const renderStepBody = (s: any) => {
    const t = s.stepType || 1;
    if (t === 3 && s.stepImage) {
      return (
        <View className="step-body">
          <Image className="step-body__image" src={s.stepImage} mode="widthFix" />
        </View>
      );
    }
    if (t === 2 && s.exampleText) {
      return (
        <View className="step-body">
          <Text className="step-body__text step-body__text--link">{s.exampleText}</Text>
          <View className="step-body__actions">
            <View className="step-body__action" onClick={() => handleCopy(s.exampleText)}>复制链接</View>
            <View className="step-body__action step-body__action--primary" onClick={() => handleOpenLink(s.exampleText)}>立即跳转</View>
          </View>
        </View>
      );
    }
    if (t === 1 && s.exampleText) {
      return (
        <View className="step-body">
          <Text className="step-body__text">{s.exampleText}</Text>
          <View className="step-body__actions">
            <View className="step-body__action" onClick={() => handleCopy(s.exampleText)}>复制文字</View>
          </View>
        </View>
      );
    }
    return null;
  };

  if (!detail) return <View className="activity-detail__loading"><Text>加载中...</Text></View>;
  return (
    <View className="activity-detail">
      <View className="activity-detail__hero">
        {detail.cover && <Image className="activity-detail__cover" src={detail.cover} mode="aspectFill" />}
        <View className="activity-detail__hero-mask" />
        <View className="activity-detail__hero-content">
          <Text className="activity-detail__hero-title">{detail.title}</Text>
          <View className="activity-detail__hero-meta">
            <Text>{detail.startTime?.slice(0, 10)} ~ {detail.endTime?.slice(0, 10)}</Text>
            <Text>{detail.joinCount || 0}人参与</Text>
          </View>
        </View>
      </View>

      {detail.content && (
        <View className="activity-detail__section card">
          <Text className="activity-detail__section-title">活动说明</Text>
          <View className="activity-detail__richtext">
            <RichText nodes={detail.content} />
          </View>
        </View>
      )}
      {detail.rewards?.length > 0 && (
        <View className="activity-detail__section card">
          <Text className="activity-detail__section-title">活动奖励</Text>
          <View className="activity-detail__rewards">
            {detail.rewards.map((r: any) => (
              <View key={r.rewardId} className="reward-card">
                <Text className="reward-card__icon">{rewardTypeIcon[r.rewardType] || '🎁'}</Text>
                <View className="reward-card__info">
                  <Text className="reward-card__name">{r.rewardName}</Text>
                  <Text className="reward-card__value">x{r.rewardValue}</Text>
                </View>
              </View>
            ))}
          </View>
        </View>
      )}
      {detail.steps?.length > 0 && (
        <View className="activity-detail__section card">
          <Text className="activity-detail__section-title">活动步骤</Text>
          <View className="activity-detail__steps">
            {detail.steps.map((s: any, i: number) => {
              const done = isStepDone(s.stepId);
              const active = isStepActive(i);
              const hasAction = s.exampleText || s.stepImage;
              return (
                <View key={s.stepId} className={`step-card ${done ? 'step-card--done' : ''} ${active && !done ? 'step-card--active' : ''}`}>
                  <View className="step-card__header">
                    <View className={`step-card__num ${done ? 'step-card__num--done' : ''}`}>
                      {done ? '✓' : (s.stepNo || i + 1)}
                    </View>
                    <View className="step-card__info">
                      <View className="step-card__title-row">
                        <Text className="step-card__title">{s.title}</Text>
                        {hasAction && <Text className="step-card__tag">{stepTypeLabel[s.stepType] || '步骤'}</Text>}
                      </View>
                      {s.description && <Text className="step-card__desc">{s.description}</Text>}
                    </View>
                  </View>
                  {(active || done) && renderStepBody(s)}
                  {detail.joined && !done && active && (
                    <View className="step-card__complete" onClick={() => handleCompleteStep(s.stepId)}>完成此步骤</View>
                  )}
                  {!active && !done && (
                    <View className="step-card__locked">请先完成上一步骤</View>
                  )}
                </View>
              );
            })}
          </View>
        </View>
      )}

      <View className="activity-detail__bottom safe-bottom">
        <View
          className={`activity-detail__btn ${detail.joined || joining ? 'activity-detail__btn--disabled' : ''}`}
          onClick={handleJoin}
        >
          {detail.joined ? '已参与' : joining ? '处理中...' : '立即参与'}
        </View>
      </View>
    </View>
  );
}
