import { View, Text, Image } from '@tarojs/components';
import './index.scss';

interface Props {
  coachId?: string;
  avatar?: string;
  nickname: string;
  intro?: string;
  minPrice?: number;
  score?: number;
  isOnline?: boolean;
  orderCount?: number;
  onClick?: () => void;
}

export default function CoachCard({ avatar, nickname, intro, minPrice, score, isOnline, onClick }: Props) {
  return (
    <View className="coach-card card" onClick={onClick}>
      <View className="coach-card__cover">
        {avatar ? <Image className="coach-card__avatar" src={avatar} mode="aspectFill" /> : <View className="coach-card__avatar-placeholder" />}
        {isOnline && <View className="coach-card__online" />}
      </View>
      <View className="coach-card__info">
        <Text className="coach-card__name">{nickname}</Text>
        {intro && <Text className="coach-card__intro">{intro}</Text>}
        <View className="coach-card__bottom">
          {minPrice !== undefined && <Text className="price">{(minPrice / 100).toFixed(0)}<Text className="coach-card__unit">元起</Text></Text>}
          {score !== undefined && <Text className="coach-card__score">★ {score.toFixed(1)}</Text>}
        </View>
      </View>
    </View>
  );
}
