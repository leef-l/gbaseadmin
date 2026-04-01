import { useState, useCallback } from 'react';
import { View, Text, Input } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import { useAuthStore } from '../../store/auth';
import { getReviewList, replyReview } from '../../api/review';
import EmptyState from '../../components/EmptyState';
import LoadMore from '../../components/LoadMore';
import './reviews.scss';

export default function WorkspaceReviewsPage() {
  const { userInfo } = useAuthStore();
  const [list, setList] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(false);
  const [loading, setLoading] = useState(false);

  // 回复弹窗状态
  const [replyTarget, setReplyTarget] = useState<{ reviewId: string; content: string } | null>(null);
  const [replyText, setReplyText] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const fetchList = useCallback(async (p = 1) => {
    if (!userInfo?.id) return;
    setLoading(true);
    try {
      const res = await getReviewList({ coachId: userInfo.id, page: p, pageSize: 10 });
      const items = res.list || [];
      setList(p === 1 ? items : (prev) => [...prev, ...items]);
      setHasMore(items.length >= 10);
      setPage(p);
    } catch { /* ignore */ }
    setLoading(false);
  }, [userInfo?.id]);

  useLoad(() => { fetchList(1); });
  usePullDownRefresh(() => { fetchList(1).then(() => Taro.stopPullDownRefresh()); });
  useReachBottom(() => { if (hasMore && !loading) fetchList(page + 1); });

  const openReply = (item: any) => {
    setReplyTarget({ reviewId: item.reviewId, content: item.content });
    setReplyText('');
  };

  const handleReply = async () => {
    if (!replyText.trim()) {
      Taro.showToast({ title: '请输入回复内容', icon: 'none' });
      return;
    }
    if (!replyTarget) return;
    setSubmitting(true);
    try {
      await replyReview(replyTarget.reviewId, replyText.trim());
      Taro.showToast({ title: '回复成功', icon: 'success' });
      setReplyTarget(null);
      // 本地更新列表中该条评价的回复
      setList((prev) =>
        prev.map((r) =>
          r.reviewId === replyTarget.reviewId ? { ...r, reply: replyText.trim() } : r
        )
      );
    } catch { /* ignore */ }
    setSubmitting(false);
  };

  const renderStars = (score: number) => {
    return Array.from({ length: 5 }, (_, i) => (
      <Text key={i} className={`ws-reviews__star ${i < score ? 'ws-reviews__star--on' : ''}`}>★</Text>
    ));
  };

  return (
    <View className="ws-reviews">
      <View className="ws-reviews__content">
        {list.length === 0 && !loading ? (
          <EmptyState text="暂无评价" />
        ) : (
          list.map((r) => (
            <View key={r.reviewId} className="ws-reviews__card card">
              <View className="ws-reviews__card-top">
                <Text className="ws-reviews__card-user">{r.isAnonymous ? '匿名用户' : (r.nickName || '用户')}</Text>
                <View className="ws-reviews__stars">{renderStars(r.score || 5)}</View>
              </View>
              <Text className="ws-reviews__card-content">{r.content}</Text>
              <Text className="ws-reviews__card-time">{r.createdAt || ''}</Text>

              {r.reply ? (
                <View className="ws-reviews__reply">
                  <Text className="ws-reviews__reply-label">我的回复：</Text>
                  <Text className="ws-reviews__reply-text">{r.reply}</Text>
                </View>
              ) : (
                <View className="ws-reviews__reply-btn" onClick={() => openReply(r)}>
                  回复评价
                </View>
              )}
            </View>
          ))
        )}
        {list.length > 0 && <LoadMore hasMore={hasMore} loading={loading} />}
      </View>

      {/* 回复弹窗 */}
      {replyTarget && (
        <View className="ws-reviews__modal-mask" onClick={() => setReplyTarget(null)}>
          <View className="ws-reviews__modal" onClick={(e) => e.stopPropagation()}>
            <View className="ws-reviews__modal-header">
              <Text className="ws-reviews__modal-title">回复评价</Text>
              <Text className="ws-reviews__modal-close" onClick={() => setReplyTarget(null)}>✕</Text>
            </View>
            <View className="ws-reviews__modal-body">
              <View className="ws-reviews__origin">
                <Text className="ws-reviews__origin-label">用户评价：</Text>
                <Text className="ws-reviews__origin-text">{replyTarget.content}</Text>
              </View>
              <Input
                className="ws-reviews__reply-input"
                placeholder="请输入回复内容..."
                value={replyText}
                onInput={(e) => setReplyText(e.detail.value)}
              />
            </View>
            <View
              className={`ws-reviews__modal-submit ${submitting ? 'ws-reviews__modal-submit--loading' : ''}`}
              onClick={submitting ? undefined : handleReply}
            >
              {submitting ? '提交中...' : '确认回复'}
            </View>
          </View>
        </View>
      )}
    </View>
  );
}
