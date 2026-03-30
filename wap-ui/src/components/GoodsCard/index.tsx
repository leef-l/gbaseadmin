import { View, Text, Image } from '@tarojs/components';
import './index.scss';

interface Props {
  goodsId?: string;
  coverImage?: string;
  title: string;
  coachName?: string;
  coachAvatar?: string;
  price: number;
  unit?: string;
  salesNum?: number;
  onClick?: () => void;
}

export default function GoodsCard({ coverImage, title, coachName, coachAvatar, price, onClick }: Props) {
  return (
    <View className="goods-card card" onClick={onClick}>
      <View className="goods-card__cover">
        {coverImage ? <Image className="goods-card__img" src={coverImage} mode="aspectFill" /> : <View className="goods-card__placeholder" />}
        <View className="goods-card__mask" />
        <Text className="goods-card__title">{title}</Text>
      </View>
      <View className="goods-card__info">
        <View className="goods-card__left">
          {coachAvatar && <Image className="goods-card__avatar" src={coachAvatar} />}
          {coachName && <Text className="goods-card__coach">{coachName}</Text>}
        </View>
        <View className="goods-card__right">
          <Text className="price">¥{(price / 100).toFixed(2)}</Text>
          <View className="goods-card__btn">下单</View>
        </View>
      </View>
    </View>
  );
}
