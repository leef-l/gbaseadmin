import { useState } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import CoachCard from '../../components/CoachCard';
import LoadMore from '../../components/LoadMore';
import EmptyState from '../../components/EmptyState';
import './list.scss';

const categories = ['全部', '游戏陪玩', '语音聊天', '看电影', '唱歌', '叫醒哄睡'];
const sorts = ['综合', '评分最高', '接单最多', '价格最低'];

export default function CoachListPage() {
  const [activeCat, setActiveCat] = useState(0);
  const [activeSort, setActiveSort] = useState(0);
  const [list, setList] = useState<any[]>([]);
  const [hasMore, setHasMore] = useState(true);
  const [loading, setLoading] = useState(false);

  useLoad(() => {
    setList([
      { id: '1', nickname: '小甜', tags: ['王者荣耀', '英雄联盟'], price: 3000, score: 4.9, online: true, level: 'Lv.5' },
      { id: '2', nickname: '阿杰', tags: ['和平精英', '语音聊天'], price: 2500, score: 4.8, online: true, level: 'Lv.4' },
      { id: '3', nickname: '小鱼', tags: ['原神', '看电影'], price: 2000, score: 4.7, online: false, level: 'Lv.3' },
      { id: '4', nickname: '大白', tags: ['LOL', '唱歌'], price: 3500, score: 4.9, online: true, level: 'Lv.5' },
      { id: '5', nickname: '小美', tags: ['王者荣耀', '语音'], price: 2800, score: 4.6, online: true, level: 'Lv.4' },
      { id: '6', nickname: '阿飞', tags: ['和平精英', '吃鸡'], price: 2200, score: 4.5, online: false, level: 'Lv.3' },
    ]);
  });

  usePullDownRefresh(() => Taro.stopPullDownRefresh());
  useReachBottom(() => {});

  return (
    <View className="coach-list">
      <View className="coach-list__filter">
        <View className="coach-list__categories">
          {categories.map((c, i) => (
            <Text key={i} className={`coach-list__cat ${activeCat === i ? 'coach-list__cat--active' : ''}`} onClick={() => setActiveCat(i)}>{c}</Text>
          ))}
        </View>
        <View className="coach-list__sorts">
          {sorts.map((s, i) => (
            <Text key={i} className={`coach-list__sort ${activeSort === i ? 'coach-list__sort--active' : ''}`} onClick={() => setActiveSort(i)}>{s}</Text>
          ))}
        </View>
      </View>
      {list.length === 0 ? <EmptyState text="暂无陪玩师" /> : (
        <View className="coach-list__grid">
          {list.map((c) => (
            <CoachCard key={c.id} {...c} onClick={() => Taro.navigateTo({ url: `/pages/coach/detail?id=${c.id}` })} />
          ))}
        </View>
      )}
      <LoadMore hasMore={hasMore} loading={loading} />
    </View>
  );
}
