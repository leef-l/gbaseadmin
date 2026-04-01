import { useState, useEffect, useRef } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, useRouter } from '@tarojs/taro';
import { getOrderDetail } from '../../api/order';
import { pay } from '../../api/payment';
import { getMemberInfo } from '../../api/member';
import './pay.scss';

const payMethods = [
  { key: 'wechat', icon: '💚', name: '微信支付', color: '#00b894' },
  { key: 'alipay', icon: '💙', name: '支付宝支付', color: '#0984e3' },
  { key: 'balance', icon: '💜', name: '余额支付', color: '#7c3aed' },
];

// 30 分钟倒计时总秒数
const COUNTDOWN_SECONDS = 30 * 60;

function formatCountdown(seconds: number): string {
  const m = Math.floor(seconds / 60);
  const s = seconds % 60;
  return `${String(m).padStart(2, '0')}:${String(s).padStart(2, '0')}`;
}

export default function PayPage() {
  const { params } = useRouter();
  const [selected, setSelected] = useState('wechat');
  const [amount, setAmount] = useState(0);
  const [orderNo, setOrderNo] = useState('');
  const [balance, setBalance] = useState(0);
  const [countdown, setCountdown] = useState(COUNTDOWN_SECONDS);
  const [expired, setExpired] = useState(false);
  const timerRef = useRef<ReturnType<typeof setInterval> | null>(null);

  useLoad(async () => {
    try {
      const [orderData, memberData] = await Promise.all([
        getOrderDetail(params.orderId || ''),
        getMemberInfo(),
      ]);
      setAmount(orderData?.payAmount || 0);
      setOrderNo(orderData?.orderNo || params.orderId || '');
      setBalance(memberData?.balance || 0);

      // 若接口返回订单创建时间，则基于该时间计算剩余秒数
      if (orderData?.createdAt) {
        const createdMs = new Date(orderData.createdAt).getTime();
        const elapsed = Math.floor((Date.now() - createdMs) / 1000);
        const remaining = COUNTDOWN_SECONDS - elapsed;
        if (remaining <= 0) {
          setExpired(true);
          setCountdown(0);
          return;
        }
        setCountdown(remaining);
      }
    } catch {}
  });

  useEffect(() => {
    if (expired) return;
    timerRef.current = setInterval(() => {
      setCountdown((prev) => {
        if (prev <= 1) {
          clearInterval(timerRef.current!);
          setExpired(true);
          return 0;
        }
        return prev - 1;
      });
    }, 1000);
    return () => {
      if (timerRef.current) clearInterval(timerRef.current);
    };
  }, [expired]);

  const handlePay = async () => {
    if (selected === 'balance' && balance < amount) {
      Taro.showToast({ title: '余额不足', icon: 'none' });
      return;
    }
    try {
      Taro.showLoading({ title: '支付中...' });
      await pay({ orderId: params.orderId, payMethod: selected });
      Taro.hideLoading();
      Taro.showToast({ title: '支付成功', icon: 'success' });
      setTimeout(() => Taro.redirectTo({ url: `/pages/order/detail?id=${params.orderId}` }), 1500);
    } catch {
      Taro.hideLoading();
    }
  };

  return (
    <View className="pay">
      <View className="pay__countdown">
        {expired ? '支付超时，订单已关闭' : `请在 ${formatCountdown(countdown)} 内完成支付`}
      </View>
      <View className="pay__amount">
        <Text className="pay__price">¥{(amount / 100).toFixed(2)}</Text>
        <Text className="pay__order-no">订单编号: {orderNo}</Text>
      </View>
      <View className="pay__methods">
        {payMethods.map((m) => (
          <View key={m.key} className={`pay__method card ${selected === m.key ? 'pay__method--active' : ''}`} onClick={() => setSelected(m.key)}>
            <Text style={{ fontSize: '24px' }}>{m.icon}</Text>
            <View style={{ flex: 1 }}>
              <Text className="pay__method-name">{m.name}</Text>
              {m.key === 'balance' && <Text className="pay__balance">余额 ¥{(balance / 100).toFixed(2)}</Text>}
            </View>
            <View className={`pay__radio ${selected === m.key ? 'pay__radio--checked' : ''}`} />
          </View>
        ))}
      </View>
      <View className="bottom-bar">
        <View className="btn-primary" style={{ width: '100%', textAlign: 'center' }} onClick={handlePay}>确认支付 ¥{(amount / 100).toFixed(2)}</View>
      </View>
    </View>
  );
}
