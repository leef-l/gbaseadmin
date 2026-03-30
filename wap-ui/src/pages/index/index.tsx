import { useState, useCallback } from 'react';
import { View, Text, Swiper, SwiperItem } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import CoachCard from '../../components/CoachCard';
import GoodsCard from '../../components/GoodsCard';
import ActivityCard from '../../components/ActivityCard';
import LoadMore from '../../components/LoadMore';
import './index.scss';

const navItems = [
  { text: '游戏陪玩', color: '#6C5CE7', icon: '🎮' },
  { text: '语音聊天', color: '#FF6B6B', icon: '🎤' },
  { text: '看电影', color: '#00D2D3', icon: '🎬' },
  { text: '唱歌', color: '#FECA57', icon: '🎵' },
  { text: '更多', color: '#fd79a8', icon: '📋' },
];

export default function IndexPage() {
  const [activities, setActivities] = useState<any[]>([]);
  const [coaches, setCoaches] = useState<any[]>([]);
  const [goods, setGoods] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);
  const [loading, setLoading] = useState(false);

  const fetchData = useCallback(async (reset = false) => {
    // TODO: 接入真实API
    setActivities([
      { id: '1', title: '新人专享活动', type: '限时', time: '3月30日', participants: 128, reward: '赠送50元券' },
      { id: '2', title: '邀请好友赢奖励', type: '长期', time: '长期有效', participants: 356, reward: '双倍积分' },
    ]);
    setCoaches([
      { id: '1', nickname: '小甜', tags: ['王者荣耀', '英雄联盟'], price: 3000, score: 4.9, online: true, level: 'Lv.5' },
      { id: '2', nickname: '阿杰', tags: ['和平精英', '语音聊天'], price: 2500, score: 4.8, online: true, level: 'Lv.4' },
      { id: '3', nickname: '小鱼', tags: ['原神', '看电影'], price: 2000, score: 4.7, online: false, level: 'Lv.3' },
      { id: '4', nickname: '大白', tags: ['LOL', '唱歌'], price: 3500, score: 4.9, online: true, level: 'Lv.5' },
    ]);
    setGoods([
      { id: '1', title: '王者荣耀陪玩', coachName: '小甜', price: 3000, desc: '国服百星，带飞上分' },
      { id: '2', title: '和平精英吃鸡', coachName: '阿杰', price: 2500, desc: '职业选手，稳定吃鸡' },
    ]);
  }, []);

  useLoad(() => { fetchData(true); });

  usePullDownRefresh(() => {
    fetchData(true).then(() => Taro.stopPullDownRefresh());
  });

  useReachBottom(() => {
    if (!hasMore || loading) return;
    // TODO: 加载更多商品
  });

  return (
    <View className="home">
      <View className="home__header">
        <View className="home__top-bar">
          <View className="home__location">📍 全国</View>
          <View className="home__icons">🔔 📷</View>
        </View>
        <View className="home__search" onClick={() => Taro.navigateTo({ url: '/pages/search/index' })}>
          搜索陪玩师、游戏、服务...
        </View>
      </View>

      {/* Banner */}
      <View className="home__banner">
        <Swiper className="swiper" autoplay circular indicatorDots indicatorActiveColor="#fff">
          <SwiperItem><View className="banner-item banner-1"><View><Text className="banner-text">新人专享</Text><Text className="banner-desc">首单立减20元</Text></View></View></SwiperItem>
          <SwiperItem><View className="banner-item banner-2"><View><Text className="banner-text">邀请有礼</Text><Text className="banner-desc">邀请好友各得50元</Text></View></View></SwiperItem>
          <SwiperItem><View className="banner-item banner-3"><View><Text className="banner-text">限时活动</Text><Text className="banner-desc">参与赢大奖</Text></View></View></SwiperItem>
        </Swiper>
      </View>

      {/* 快捷导航 */}
      <View className="home__nav">
        {navItems.map((item, i) => (
          <View key={i} className="home__nav-item" onClick={() => Taro.navigateTo({ url: `/pages/coach/list?categoryId=${i + 1}` })}>
            <View className="home__nav-icon" style={{ background: item.color }}>{item.icon}</View>
            <Text className="home__nav-text">{item.text}</Text>
          </View>
        ))}
      </View>

      {/* 热门活动 */}
      <View className="home__activities">
        <View className="section-header">
          <Text className="title">热门活动</Text>
          <Text className="more" onClick={() => Taro.switchTab({ url: '/pages/activity/list' })}>查看全部 &gt;</Text>
        </View>
        <View className="scroll-wrap">
          {activities.map((a) => (
            <ActivityCard key={a.id} {...a} onClick={() => Taro.navigateTo({ url: `/pages/activity/detail?id=${a.id}` })} />
          ))}
        </View>
      </View>

      {/* 推荐陪玩师 */}
      <View className="home__coaches">
        <View className="section-header">
          <Text className="title">推荐陪玩师</Text>
          <Text className="more" onClick={() => Taro.navigateTo({ url: '/pages/coach/list' })}>查看全部 &gt;</Text>
        </View>
        <View className="grid">
          {coaches.map((c) => (
            <CoachCard key={c.id} {...c} onClick={() => Taro.navigateTo({ url: `/pages/coach/detail?id=${c.id}` })} />
          ))}
        </View>
      </View>

      {/* 热门商品 */}
      <View className="section-header">
        <Text className="title">热门服务</Text>
        <Text className="more">查看全部 &gt;</Text>
      </View>
      <View className="home__goods">
        {goods.map((g) => (
          <GoodsCard key={g.id} {...g} onClick={() => Taro.navigateTo({ url: `/pages/goods/detail?id=${g.id}` })} />
        ))}
        <LoadMore hasMore={hasMore} loading={loading} />
      </View>
    </View>
  );
}