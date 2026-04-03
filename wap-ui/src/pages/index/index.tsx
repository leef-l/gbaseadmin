import { useState, useCallback } from 'react';
import { View, Text, Swiper, SwiperItem, Image } from '@tarojs/components';
import Taro, { useLoad, usePullDownRefresh, useReachBottom } from '@tarojs/taro';
import { getActivityList } from '../../api/activity';
import { getCoachList } from '../../api/coach';
import { getGoodsList } from '../../api/goods';
import { getBannerList, BannerItem } from '../../api/banner';
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
  const [banners, setBanners] = useState<BannerItem[]>([]);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);
  const [loading, setLoading] = useState(false);

  const handleBannerClick = (item: BannerItem) => {
    switch (item.linkType) {
      case 1: // 内页
        Taro.navigateTo({ url: item.linkValue });
        break;
      case 2: // 外链
        if (Taro.getEnv() === Taro.ENV_TYPE.WEB) {
          window.location.href = item.linkValue;
        } else {
          // 小程序环境需要 webview 页面承载，暂用路由跳转（webview 页面需单独开发）
          Taro.navigateTo({ url: `/pages/webview/index?url=${encodeURIComponent(item.linkValue)}` });
        }
        break;
      case 3: // 活动页
        Taro.navigateTo({ url: `/pages/activity/detail?activityId=${item.linkValue}` });
        break;
      case 4: // 商品页
        Taro.navigateTo({ url: `/pages/goods/detail?id=${item.linkValue}` });
        break;
      case 5: // 陪玩师页
        Taro.navigateTo({ url: `/pages/coach/detail?id=${item.linkValue}` });
        break;
      case 6: // 唤醒App
        if (Taro.getEnv() === Taro.ENV_TYPE.WEB) {
          window.location.href = item.linkValue;
        } else {
          // 小程序环境不支持唤醒外部 App
          Taro.showToast({ title: '请在浏览器中打开', icon: 'none' });
        }
        break;
    }
  };

  const fetchData = useCallback(async (reset = false) => {
    const p = reset ? 1 : page;
    if (!reset && loading) return;
    setLoading(true);
    try {
      const [actRes, coachRes, goodsRes, bannerRes] = await Promise.all([
        getActivityList({ page: 1, pageSize: 4 }),
        getCoachList({ page: 1, pageSize: 4 }),
        getGoodsList({ page: p, pageSize: 10 }),
        getBannerList(),
      ]);
      setActivities(actRes?.list || []);
      setCoaches(coachRes?.list || []);
      setBanners(bannerRes?.list || []);
      if (reset) {
        setGoods(goodsRes?.list || []);
        setPage(2);
      } else {
        setGoods((prev) => [...prev, ...(goodsRes?.list || [])]);
        setPage(p + 1);
      }
      setHasMore((goodsRes?.list || []).length >= 10);
    } catch {
      // request.ts 已统一 toast
    } finally {
      setLoading(false);
    }
  }, [page, loading]);

  useLoad(() => { fetchData(true); });

  usePullDownRefresh(() => {
    fetchData(true).then(() => Taro.stopPullDownRefresh());
  });

  useReachBottom(() => {
    if (!hasMore || loading) return;
    fetchData(false);
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
        {banners.length > 0 ? (
          <Swiper className="swiper" autoplay circular indicatorDots indicatorActiveColor="#fff">
            {banners.map((item) => (
              <SwiperItem key={item.bannerId} onClick={() => handleBannerClick(item)}>
                <Image className="banner-img" src={item.image} mode="aspectFill" />
              </SwiperItem>
            ))}
          </Swiper>
        ) : (
          <Swiper className="swiper" autoplay circular indicatorDots indicatorActiveColor="#fff">
            <SwiperItem><View className="banner-item banner-1"><View><Text className="banner-text">新人专享</Text><Text className="banner-desc">首单立减20元</Text></View></View></SwiperItem>
            <SwiperItem><View className="banner-item banner-2"><View><Text className="banner-text">邀请有礼</Text><Text className="banner-desc">邀请好友各得50元</Text></View></View></SwiperItem>
            <SwiperItem><View className="banner-item banner-3"><View><Text className="banner-text">限时活动</Text><Text className="banner-desc">参与赢大奖</Text></View></View></SwiperItem>
          </Swiper>
        )}
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
            <ActivityCard key={a.activityId} {...a} onClick={() => Taro.navigateTo({ url: `/pages/activity/detail?activityId=${a.activityId}` })} />
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
            <CoachCard key={c.coachId} {...c} onClick={() => Taro.navigateTo({ url: `/pages/coach/detail?id=${c.coachId}` })} />
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
          <GoodsCard key={g.goodsId} {...g} onClick={() => Taro.navigateTo({ url: `/pages/goods/detail?id=${g.goodsId}` })} />
        ))}
        <LoadMore hasMore={hasMore} loading={loading} />
      </View>
    </View>
  );
}
