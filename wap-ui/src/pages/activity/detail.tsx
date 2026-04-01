import { useState, useCallback } from 'react';
import { View, Text, Image, RichText } from '@tarojs/components';
import Taro, { useLoad, useRouter } from '@tarojs/taro';
import { getActivityDetail, joinActivity, completeStep } from '../../api/activity';
import { useAuthStore } from '../../store/auth';
import './detail.scss';

const BASE_URL = process.env.TARO_APP_API || '';

const rewardTypeIcon: Record<number, string> = { 1: '💰', 2: '🎫', 3: '⭐', 4: '👑' };
const stepTypeLabel: Record<number, string> = { 1: '文字', 2: '链接', 3: '图片' };

async function uploadImage(filePath: string): Promise<string> {
  const token = useAuthStore.getState().token;
  return new Promise((resolve, reject) => {
    Taro.uploadFile({
      url: `${BASE_URL}/api/upload/uploader/upload`,
      filePath,
      name: 'file',
      header: token ? { Authorization: `Bearer ${token}` } : {},
      success: (res) => {
        try {
          const body = JSON.parse(res.data);
          if (body.code === 0 && body.data?.url) {
            resolve(body.data.url);
          } else {
            reject(new Error(body.message || '上传失败'));
          }
        } catch {
          reject(new Error('上传响应解析失败'));
        }
      },
      fail: () => reject(new Error('上传失败')),
    });
  });
}

export default function ActivityDetailPage() {
  const router = useRouter();
  const [detail, setDetail] = useState<any>(null);
  const [joining, setJoining] = useState(false);
  // 每个步骤的用户上传图片，key 为 stepId
  const [stepImages, setStepImages] = useState<Record<string, string>>({});
  const [uploadingStepId, setUploadingStepId] = useState<string | null>(null);

  const fetchDetail = useCallback(async (id: string) => {
    try {
      const res = await getActivityDetail(id);
      setDetail(res || null);
    } catch (e) {
      console.error(e);
    }
  }, []);

  useLoad(() => {
    const activityId = router.params.activityId;
    if (activityId) fetchDetail(activityId);
  });

  const handleJoin = async () => {
    if (!detail || joining || detail.joined) return;
    setJoining(true);
    try {
      await joinActivity(detail.activityId);
      Taro.showToast({ title: '参与成功', icon: 'success' });
      fetchDetail(detail.activityId);
    } catch {
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
    } catch {
      Taro.showToast({ title: '操作失败', icon: 'none' });
    }
  };

  const handleCopy = (text: string) => {
    Taro.setClipboardData({ data: text });
    Taro.showToast({ title: '已复制', icon: 'success' });
  };

  const handleOpenLink = (url: string) => {
    Taro.setClipboardData({ data: url });
    Taro.showToast({ title: '链接已复制，请在浏览器打开', icon: 'none', duration: 2000 });
  };

  const handleChooseImage = async (stepId: string) => {
    if (uploadingStepId) return;
    try {
      const res = await Taro.chooseImage({
        count: 1,
        sizeType: ['compressed'],
        sourceType: ['album', 'camera'],
      });
      const tempPath = res.tempFilePaths[0];
      // 先显示本地预览
      setStepImages((prev) => ({ ...prev, [stepId]: tempPath }));
      setUploadingStepId(stepId);
      Taro.showLoading({ title: '上传中...' });
      const url = await uploadImage(tempPath);
      setStepImages((prev) => ({ ...prev, [stepId]: url }));
      Taro.showToast({ title: '上传成功', icon: 'success' });
    } catch {
      setStepImages((prev) => { const n = { ...prev }; delete n[stepId]; return n; });
      Taro.showToast({ title: '上传失败，请重试', icon: 'none' });
    } finally {
      setUploadingStepId(null);
      Taro.hideLoading();
    }
  };

  const isStepDone = (stepId: string) => detail.completedSteps?.includes(stepId);

  const isStepActive = (index: number) => {
    if (!detail.joined) return false;
    if (index === 0) return true;
    return isStepDone(detail.steps[index - 1]?.stepId);
  };

  const renderStepBody = (s: any, active: boolean) => {
    const t = s.stepType || 1;

    const needFill = s.isRequired === 1;

    // 文字步骤：展示示例文字 + 立即复制按钮
    if (t === 1) {
      return (
        <View className="step-body">
          {s.exampleText && (
            <>
              <Text className="step-body__label">示例参考</Text>
              <Text className="step-body__text">{s.exampleText}</Text>
            </>
          )}
          {active && s.exampleText && (
            <View className="step-body__actions">
              <View className="step-body__action step-body__action--primary" onClick={() => handleCopy(s.exampleText)}>
                立即复制
              </View>
            </View>
          )}
        </View>
      );
    }

    // 链接步骤：展示链接 + 立即前往按钮
    if (t === 2) {
      return (
        <View className="step-body">
          {s.exampleText && (
            <>
              <Text className="step-body__label">目标链接</Text>
              <Text className="step-body__text step-body__text--link">{s.exampleText}</Text>
            </>
          )}
          {active && s.exampleText && (
            <View className="step-body__actions">
              <View className="step-body__action" onClick={() => handleCopy(s.exampleText)}>
                复制链接
              </View>
              <View className="step-body__action step-body__action--primary" onClick={() => handleOpenLink(s.exampleText)}>
                立即前往
              </View>
            </View>
          )}
        </View>
      );
    }

    // 图片步骤：左边示例图片（右上角"示例"角标），右边用户上传区（仅 isRequired=1 时显示上传区）
    if (t === 3) {
      const userImg = stepImages[s.stepId];
      const isUploading = uploadingStepId === s.stepId;
      return (
        <View className="step-body">
          <View className={needFill ? 'step-body__img-row' : ''}>
            {/* 示例图片 */}
            <View className="step-body__img-wrap">
              {s.stepImage ? (
                <>
                  <Image className="step-body__img" src={s.stepImage} mode="aspectFill" />
                  <Text className="step-body__img-badge">示例</Text>
                </>
              ) : (
                <View className="step-body__img step-body__img--placeholder">
                  <Text className="step-body__img-placeholder-text">暂无示例</Text>
                </View>
              )}
            </View>
            {/* 右：用户上传区（仅需要填写时显示） */}
            {needFill && (
              <View
                className={`step-body__img-wrap step-body__img-wrap--upload ${!active ? 'step-body__img-wrap--disabled' : ''}`}
                onClick={() => active && handleChooseImage(s.stepId)}
              >
                {userImg ? (
                  <>
                    <Image className="step-body__img" src={userImg} mode="aspectFill" />
                    {active && !isUploading && (
                      <Text className="step-body__img-badge step-body__img-badge--reupload">重新上传</Text>
                    )}
                  </>
                ) : (
                  <View className="step-body__img step-body__img--upload-placeholder">
                    <Text className="step-body__upload-icon">📷</Text>
                    <Text className="step-body__upload-text">{active ? '点击上传' : '请先报名'}</Text>
                  </View>
                )}
              </View>
            )}
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
              return (
                <View key={s.stepId} className={`step-card ${done ? 'step-card--done' : ''} ${active && !done ? 'step-card--active' : ''}`}>
                  <View className="step-card__header">
                    <View className={`step-card__num ${done ? 'step-card__num--done' : ''}`}>
                      {done ? '✓' : (s.stepNo || i + 1)}
                    </View>
                    <View className="step-card__info">
                      <View className="step-card__title-row">
                        <Text className="step-card__title">{s.title}</Text>
                        <Text className="step-card__tag">{stepTypeLabel[s.stepType] || '步骤'}</Text>
                      </View>
                      {s.description && <Text className="step-card__desc">{s.description}</Text>}
                    </View>
                  </View>

                  {renderStepBody(s, active)}

                  {detail.joined && !done && active && (
                    <View className="step-card__complete" onClick={() => handleCompleteStep(s.stepId)}>
                      完成此步骤
                    </View>
                  )}
                  {!active && !done && (
                    <View className="step-card__locked">
                      {!detail.joined ? '请先报名' : '请先完成上一步骤'}
                    </View>
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
