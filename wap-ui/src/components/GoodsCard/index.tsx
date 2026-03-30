import { View, Text, Image } from '@tarojs/components';
import './index.scss';

interface Props {
  id: string;
  cover?: string;
  title: string;
  coachName?: string;
  coachAvatar?: string;
  price: number;
  desc?: string;
  onClick?: () => void;
}

export default function GoodsCard({ cover, title, coachName, coachAvatar, price, desc, onClick }: Props) {
  return (
    <View className="goods-card card" onClick={onClick}>
      <View className="goods-card__cover">
        {cover ? <Image className="goods-card__img" src={cover} mode="aspectFill" /> : <View className="goods-card__placeholder" />}
        <View className="goods-card__mask" />
        <Text className="goods-card__title">{title}</Text>
      </View>
      <View className="goods-card__info">
        <View className="goods-card__left">
          {coachAvatar && <Image className="goods-card__avatar" src={coachAvatar} />}
          <View>
            {coachName && <Text className="goods-card__coach">{coachName}</Text>}
            {desc && <Text className="goods-card__desc">{desc}</Text>}
          </View>
        </View>
        <View className="goods-card__right">
          <Text className="price">¥{(price / 100).toFixed(2)}</Text>
          <View className="goods-card__btn">下单</View>
        </View>
      </View>
    </View>
  );
}
