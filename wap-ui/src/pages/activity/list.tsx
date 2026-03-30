import { useState } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh } from '@tarojs/taro';
import ActivityCard from '../../components/ActivityCard';
import EmptyState from '../../components/EmptyState';
import LoadMore from '../../components/LoadMore';
import './list.scss';

const tabs = ['全部', '进行中', '即将开始', '已结束'];

export default function ActivityListPage() {
  const [activeTab, setActiveTab] = useState(0);
  const [list, setList] = useState<any[]>([]);
  const [hasMore, setHasMore] = useState(false);

  useLoad(() => {
    // TODO: 接入真实API
    setList([
      { id: '1', title: '新人专享活动', type: '限时', time: '2026-03-30', participants: 128, reward: '赠送50元券' },
      { id: '2', title: '邀请好友赢奖励', type: '长期', time: '长期有效', participants: 356, reward: '双倍积分' },
      { id: '3', title: '春季特惠', type: '限时', time: '2026-04-15', participants: 89, reward: '满100减30' },
    ]);
  });

  usePullDownRefresh(() => {
    Taro.stopPullDownRefresh();
  });

  return (
    <View className="activity-list">
      <View className="activity-list__tabs">
        {tabs.map((t, i) => (
          <Text key={i} className={`activity-list__tab ${activeTab === i ? 'activity-list__tab--active' : ''}`} onClick={() => setActiveTab(i)}>{t}</Text>
        ))}
      </View>
      <View className="activity-list__list">
        {list.length === 0 ? <EmptyState text="暂无活动" /> : list.map((a) => (
          <ActivityCard key={a.id} {...a} onClick={() => Taro.navigateTo({ url: `/pages/activity/detail?id=${a.id}` })} />
        ))}
      </View>
      {list.length > 0 && <LoadMore hasMore={hasMore} />}
    </View>
  );
}
