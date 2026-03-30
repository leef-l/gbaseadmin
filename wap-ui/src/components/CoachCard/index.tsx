import { View, Text, Image } from '@tarojs/components';
import './index.scss';

interface Props {
  id: string;
  avatar?: string;
  nickname: string;
  tags?: string[];
  price: number;
  score?: number;
  online?: boolean;
  level?: string;
  onClick?: () => void;
}

export default function CoachCard({ avatar, nickname, tags = [], price, score, online, level, onClick }: Props) {
  return (
    <View className="coach-card card" onClick={onClick}>
      <View className="coach-card__cover">
        {avatar ? <Image className="coach-card__avatar" src={avatar} mode="aspectFill" /> : <View className="coach-card__avatar-placeholder" />}
        {online && <View className="coach-card__online" />}
        {level && <Text className="coach-card__level">{level}</Text>}
      </View>
      <View className="coach-card__info">
        <Text className="coach-card__name">{nickname}</Text>
        <View className="coach-card__tags">
          {tags.slice(0, 2).map((t, i) => <Text key={i} className="tag">{t}</Text>)}
        </View>
        <View className="coach-card__bottom">
          <Text className="price">{(price / 100).toFixed(0)}<Text className="coach-card__unit">元/局</Text></Text>
          {score !== undefined && <Text className="coach-card__score">★ {score.toFixed(1)}</Text>}
        </View>
      </View>
    </View>
  );
}
