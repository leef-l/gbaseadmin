import { View, Text } from '@tarojs/components';
import './index.scss';

interface Props {
  id: string;
  amount: number;
  minAmount: number;
  name: string;
  expireTime: string;
  status?: 'unused' | 'used' | 'expired';
  onUse?: () => void;
}

export default function CouponCard({ amount, minAmount, name, expireTime, status = 'unused', onUse }: Props) {
  const leftCls = `coupon-card__left ${status !== 'unused' ? `coupon-card__left--${status}` : ''}`;
  return (
    <View className="coupon-card card">
      <View className={leftCls}>
        <Text className="coupon-card__amount">¥{(amount / 100).toFixed(0)}</Text>
        <Text className="coupon-card__condition">满{(minAmount / 100).toFixed(0)}可用</Text>
      </View>
      <View className="coupon-card__right">
        <View>
          <Text className="coupon-card__name">{name}</Text>
          <Text className="coupon-card__expire">{expireTime}</Text>
        </View>
        {status === 'unused' && onUse && (
          <View className="coupon-card__btn" onClick={onUse}>去使用</View>
        )}
      </View>
    </View>
  );
}
