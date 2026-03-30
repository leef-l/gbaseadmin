import { useState } from 'react';
import { View, Text, Textarea } from '@tarojs/components';
import Taro, { useRouter } from '@tarojs/taro';
import { createReview } from '../../api/review';
import './review.scss';

export default function ReviewPage() {
  const { params } = useRouter();
  const [score, setScore] = useState(5);
  const [content, setContent] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    if (!content.trim()) {
      Taro.showToast({ title: '请输入评价内容', icon: 'none' });
      return;
    }
    try {
      setSubmitting(true);
      await createReview({ orderId: params.orderId, score, content: content.trim() });
      Taro.showToast({ title: '评价成功', icon: 'success' });
      setTimeout(() => Taro.navigateBack(), 1500);
    } catch {
      setSubmitting(false);
    }
  };

  return (
    <View className="review">
      <View className="review__score card">
        <Text className="review__label">服务评分</Text>
        <View className="review__stars">
          {[1, 2, 3, 4, 5].map((s) => (
            <Text key={s} className={`review__star ${s <= score ? 'review__star--active' : ''}`} onClick={() => setScore(s)}>★</Text>
          ))}
        </View>
        <Text className="review__score-text">{['', '非常差', '较差', '一般', '满意', '非常满意'][score]}</Text>
      </View>

      <View className="review__content card">
        <Text className="review__label">评价内容</Text>
        <Textarea
          className="review__textarea"
          placeholder="分享你的服务体验，帮助其他用户做出选择~"
          maxlength={500}
          value={content}
          onInput={(e) => setContent(e.detail.value)}
        />
        <Text className="review__count">{content.length}/500</Text>
      </View>

      <View className="bottom-bar">
        <View className={`btn-primary ${submitting ? 'btn-disabled' : ''}`} style={{ width: '100%', textAlign: 'center' }} onClick={submitting ? undefined : handleSubmit}>
          {submitting ? '提交中...' : '提交评价'}
        </View>
      </View>
    </View>
  );
}
