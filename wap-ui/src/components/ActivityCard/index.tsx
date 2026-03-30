import { View, Text } from '@tarojs/components';
import './index.scss';

const typeMap: Record<number, string> = { 1: '充值活动', 2: '下单活动', 3: '注册活动', 4: '图文步骤', 5: '自定义' };

interface Props {
  activityId?: string;
  title: string;
  type?: number;
  startTime?: string;
  endTime?: string;
  joinCount?: number;
  description?: string;
  onClick?: () => void;
}

export default function ActivityCard({ title, type, startTime, endTime, joinCount, onClick }: Props) {
  return (
    <View className="activity-card card" onClick={onClick}>
      <View className="activity-card__top">
        {type && <Text className="activity-card__type">{typeMap[type] || '活动'}</Text>}
        <Text className="activity-card__slogan">参与赢好礼</Text>
      </View>
      <View className="activity-card__bottom">
        <Text className="activity-card__title">{title}</Text>
        <Text className="activity-card__meta">
          {startTime && endTime ? `${startTime} ~ ${endTime}` : ''}{joinCount !== undefined && ` · ${joinCount}人参与`}
        </Text>
      </View>
    </View>
  );
}
