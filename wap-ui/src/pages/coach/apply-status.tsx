import { useState } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad } from '@tarojs/taro';
import { getApplyStatus } from '../../api/coach';
import './apply-status.scss';

export default function ApplyStatusPage() {
  const [status, setStatus] = useState<number>(-1); // 0审核中 1通过 2拒绝
  const [remark, setRemark] = useState('');

  useLoad(async () => {
    try {
      const res = await getApplyStatus();
      setStatus(res.status ?? 0);
      setRemark(res.remark || '');
    } catch {
      setStatus(0);
    }
  });

  const iconMap: Record<number, { icon: string; color: string; title: string; desc: string }> = {
    0: { icon: '⏳', color: '#e17055', title: '审核中', desc: '预计 1-3 个工作日内完成审核' },
    1: { icon: '✅', color: '#00b894', title: '审核通过', desc: '恭喜你成为陪玩师！' },
    2: { icon: '❌', color: '#d63031', title: '审核未通过', desc: remark || '请重新提交申请' },
  };

  if (status === -1) return <View className="apply-status"><Text className="apply-status__loading">加载中...</Text></View>;

  const info = iconMap[status];

  return (
    <View className="apply-status">
      <View className="apply-status__icon">{info.icon}</View>
      <Text className="apply-status__title" style={{ color: info.color }}>{info.title}</Text>
      <Text className="apply-status__desc">{info.desc}</Text>
      {status === 1 && (
        <View className="btn-primary apply-status__btn" onClick={() => Taro.navigateTo({ url: '/pages/workspace/index' })}>前往工作台</View>
      )}
      {status === 2 && (
        <View className="btn-primary apply-status__btn" onClick={() => Taro.navigateTo({ url: '/pages/coach/apply' })}>重新申请</View>
      )}
    </View>
  );
}
