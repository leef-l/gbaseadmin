import { View, Text } from '@tarojs/components';
import './index.scss';

interface Props {
  id: string;
  title: string;
  type?: string;
  time?: string;
  participants?: number;
  reward?: string;
  onClick?: () => void;
}

export default function ActivityCard({ title, type, time, participants, reward, onClick }: Props) {
  return (
    <View className="activity-card card" onClick={onClick}>
      <View className="activity-card__top">
        {type && <Text className="activity-card__type">{type}</Text>}
        <Text className="activity-card__slogan">参与赢好礼</Text>
      </View>
      <View className="activity-card__bottom">
        <Text className="activity-card__title">{title}</Text>
        <Text className="activity-card__meta">
          {time}{participants !== undefined && ` · ${participants}人参与`}
        </Text>
        {reward && <Text className="tag">{reward}</Text>}
      </View>
    </View>
  );
}
