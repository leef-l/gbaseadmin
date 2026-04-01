import { useState, useRef, useCallback, useEffect } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import { getActivityList } from '../../api/activity';
import ActivityCard from '../../components/ActivityCard';
import EmptyState from '../../components/EmptyState';
import LoadMore from '../../components/LoadMore';
import './list.scss';

const PAGE_SIZE = 10;
const tabs = ['全部', '进行中', '即将开始', '已结束'];

export default function ActivityListPage() {
  const [activeTab, setActiveTab] = useState(0);
  const [list, setList] = useState<any[]>([]);
  const [hasMore, setHasMore] = useState(false);
  const [loading, setLoading] = useState(false);
  const pageRef = useRef(1);

  const fetchList = useCallback(async (reset = false) => {
    if (loading) return;
    if (reset) pageRef.current = 1;
    setLoading(true);
    try {
      const res = await getActivityList({ page: pageRef.current, pageSize: PAGE_SIZE, status: activeTab });
      const rows = res?.list || [];
      if (reset) {
        setList(rows);
      } else {
        setList((prev) => [...prev, ...rows]);
      }
      setHasMore(rows.length >= PAGE_SIZE);
      pageRef.current += 1;
    } finally {
      setLoading(false);
    }
  }, [activeTab, loading]);

  useLoad(() => {
    // 首次加载由下方 useEffect([activeTab]) 统一触发，此处不重复请求
  });

  usePullDownRefresh(async () => {
    await fetchList(true);
    Taro.stopPullDownRefresh();
  });

  useReachBottom(() => {
    if (hasMore && !loading) fetchList();
  });

  useEffect(() => {
    fetchList(true);
  }, [activeTab]);

  return (
    <View className="activity-list">
      <View className="activity-list__tabs">
        {tabs.map((t, i) => (
          <Text key={i} className={`activity-list__tab ${activeTab === i ? 'activity-list__tab--active' : ''}`} onClick={() => { if (activeTab !== i) { setActiveTab(i); } }}>{t}</Text>
        ))}
      </View>
      <View className="activity-list__list">
        {list.length === 0 ? <EmptyState text="暂无活动" /> : list.map((a) => (
          <ActivityCard key={a.activityId} {...a} onClick={() => Taro.navigateTo({ url: `/pages/activity/detail?activityId=${a.activityId}` })} />
        ))}
      </View>
      {list.length > 0 && <LoadMore hasMore={hasMore} />}
    </View>
  );
}
