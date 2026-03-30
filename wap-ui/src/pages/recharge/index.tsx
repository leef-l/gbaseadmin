import { useState, useCallback } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad } from '@tarojs/taro';
import { getRechargePlans, createRecharge } from '../../api/recharge';
import { useAuthStore } from '../../store/auth';
import './index.scss';

export default function RechargePage() {
  const { userInfo } = useAuthStore();
  const [plans, setPlans] = useState<any[]>([]);
  const [selectedId, setSelectedId] = useState('');
  const [loading, setLoading] = useState(false);

  const fetchPlans = useCallback(async () => {
    try {
      const res = await getRechargePlans();
      const list = res?.data || [];
      setPlans(list);
      if (list.length > 0) setSelectedId(list[0].id);
    } catch (e) {
      console.error(e);
    }
  }, []);

  useLoad(() => { fetchPlans(); });

  const handleRecharge = async () => {
    if (!selectedId || loading) return;
    setLoading(true);
    try {
      await createRecharge({ planId: selectedId });
      Taro.showToast({ title: '充值成功', icon: 'success' });
      setTimeout(() => Taro.navigateBack(), 1500);
    } catch (e) {
      Taro.showToast({ title: '充值失败', icon: 'none' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <View className="recharge">
      <View className="recharge__header">
        <Text className="recharge__label">当前余额（元）</Text>
        <Text className="recharge__balance">
          {((userInfo?.balance || 0) / 100).toFixed(2)}
        </Text>
      </View>

      <View className="recharge__plans">
        <Text className="recharge__title">选择充值方案</Text>
        <View className="recharge__grid">
          {plans.map((p) => (
            <View
              key={p.id}
              className={`recharge__card ${selectedId === p.id ? 'recharge__card--active' : ''}`}
              onClick={() => setSelectedId(p.id)}
            >
              <Text className="recharge__amount">
                ¥{(p.amount / 100).toFixed(2)}
              </Text>
              {p.giftAmount > 0 && (
                <Text className="recharge__gift">
                  赠送¥{(p.giftAmount / 100).toFixed(2)}
                </Text>
              )}
            </View>
          ))}
        </View>
      </View>

      <View className="recharge__bottom safe-bottom">
        <View
          className={`recharge__btn ${loading ? 'recharge__btn--disabled' : ''}`}
          onClick={handleRecharge}
        >
          <Text>{loading ? '处理中...' : '确认充值'}</Text>
        </View>
      </View>
    </View>
  );
}
