import { useState, useCallback, useRef } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import { getCoachList } from '../../api/coach';
import CoachCard from '../../components/CoachCard';
import LoadMore from '../../components/LoadMore';
import EmptyState from '../../components/EmptyState';
import './list.scss';

const categories = ['全部', '游戏陪玩', '语音聊天', '看电影', '唱歌', '叫醒哄睡'];
const sortKeys = ['', 'score', 'orderCount', 'price_asc'];
const sorts = ['综合', '评分最高', '接单最多', '价格最低'];

export default function CoachListPage() {
  const [activeCat, setActiveCat] = useState(0);
  const [activeSort, setActiveSort] = useState(0);
  const [list, setList] = useState<any[]>([]);
  const [hasMore, setHasMore] = useState(true);
  const [loading, setLoading] = useState(false);
  const pageRef = useRef(1);

  const fetchList = useCallback(async (reset = false) => {
    if (loading) return;
    const p = reset ? 1 : pageRef.current;
    setLoading(true);
    try {
      const res = await getCoachList({
        page: p,
        pageSize: 10,
        categoryId: activeCat > 0 ? String(activeCat) : undefined,
        sortBy: sortKeys[activeSort] || undefined,
      });
      const items = res?.list || [];
      if (reset) {
        setList(items);
      } else {
        setList((prev) => [...prev, ...items]);
      }
      pageRef.current = p + 1;
      setHasMore(items.length >= 10);
    } catch {
      // request.ts 已统一 toast
    } finally {
      setLoading(false);
    }
  }, [activeCat, activeSort, loading]);

  useLoad(() => { fetchList(true); });

  usePullDownRefresh(() => {
    fetchList(true).then(() => Taro.stopPullDownRefresh());
  });

  useReachBottom(() => {
    if (!hasMore || loading) return;
    fetchList(false);
  });

  const handleCatChange = (i: number) => {
    setActiveCat(i);
    pageRef.current = 1;
    setTimeout(() => fetchList(true), 0);
  };

  const handleSortChange = (i: number) => {
    setActiveSort(i);
    pageRef.current = 1;
    setTimeout(() => fetchList(true), 0);
  };

  return (
    <View className="coach-list">
      <View className="coach-list__filter">
        <View className="coach-list__categories">
          {categories.map((c, i) => (
            <Text key={i} className={`coach-list__cat ${activeCat === i ? 'coach-list__cat--active' : ''}`} onClick={() => handleCatChange(i)}>{c}</Text>
          ))}
        </View>
        <View className="coach-list__sorts">
          {sorts.map((s, i) => (
            <Text key={i} className={`coach-list__sort ${activeSort === i ? 'coach-list__sort--active' : ''}`} onClick={() => handleSortChange(i)}>{s}</Text>
          ))}
        </View>
      </View>
      {list.length === 0 ? <EmptyState text="暂无陪玩师" /> : (
        <View className="coach-list__grid">
          {list.map((c) => (
            <CoachCard key={c.coachId} {...c} onClick={() => Taro.navigateTo({ url: `/pages/coach/detail?id=${c.coachId}` })} />
          ))}
        </View>
      )}
      <LoadMore hasMore={hasMore} loading={loading} />
    </View>
  );
}
