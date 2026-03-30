import { View, Text, Image } from '@tarojs/components';
import StatusTag from '../StatusTag';
import './index.scss';

interface Props {
  orderId?: string;
  orderNo: string;
  orderStatus: number;
  goodsTitle: string;
  goodsImage?: string;
  coachName?: string;
  quantity: number;
  goodsPrice?: number;
  payAmount: number;
  actions?: { label: string; type?: 'primary' | 'default'; onClick: () => void }[];
}

const statusMap: Record<number, { text: string; color: string }> = {
  0: { text: '待支付', color: '#e17055' },
  1: { text: '已支付', color: '#0984e3' },
  2: { text: '进行中', color: '#6C5CE7' },
  3: { text: '已完成', color: '#00b894' },
  4: { text: '已取消', color: '#b2bec3' },
  5: { text: '退款中', color: '#d63031' },
  6: { text: '已退款', color: '#b2bec3' },
};

export default function OrderCard({ orderNo, orderStatus, goodsTitle, goodsImage, coachName, quantity, goodsPrice, payAmount, actions }: Props) {
  const s = statusMap[orderStatus] || { text: '未知', color: '#b2bec3' };
  return (
    <View className="order-card card">
      <View className="order-card__header">
        <Text className="order-card__no">订单号: {orderNo}</Text>
        <StatusTag text={s.text} color={s.color} />
      </View>
      <View className="order-card__body">
        <View className="order-card__cover">
          {goodsImage ? <Image className="order-card__img" src={goodsImage} mode="aspectFill" /> : <View className="order-card__img-placeholder" />}
        </View>
        <View className="order-card__detail">
          <Text className="order-card__goods">{goodsTitle}</Text>
          {coachName && <Text className="order-card__coach">{coachName}</Text>}
          {goodsPrice && <Text className="order-card__spec">{quantity} x ¥{(goodsPrice / 100).toFixed(2)}</Text>}
        </View>
      </View>
      <View className="order-card__footer">
        <Text className="order-card__amount">实付 <Text className="price">¥{(payAmount / 100).toFixed(2)}</Text></Text>
        <View className="order-card__actions">
          {actions?.map((a, i) => (
            <View key={i} className={`order-card__btn ${a.type === 'primary' ? 'order-card__btn--primary' : ''}`} onClick={a.onClick}>{a.label}</View>
          ))}
        </View>
      </View>
    </View>
  );
}
