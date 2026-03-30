import { useState } from 'react';
import { View, Text, Input } from '@tarojs/components';
import Taro from '@tarojs/taro';
import { getCoachList } from '../../api/coach';
import { getGoodsList } from '../../api/goods';
import CoachCard from '../../components/CoachCard';
import GoodsCard from '../../components/GoodsCard';
import EmptyState from '../../components/EmptyState';
import './index.scss';

export default function SearchPage() {
  const [keyword, setKeyword] = useState('');
  const [searched, setSearched] = useState(false);
  const [resultTab, setResultTab] = useState(0);
  const [coaches, setCoaches] = useState<any[]>([]);
  const [goods, setGoods] = useState<any[]>([]);
  const [history] = useState(['王者荣耀', '语音聊天', '英雄联盟']);
  const [hotWords] = useState(['游戏陪玩', '看电影', '唱歌', '哄睡']);

  const doSearch = async () => {
    if (!keyword.trim()) return;
    setSearched(true);
    try {
      const [coachRes, goodsRes] = await Promise.all([
        getCoachList({ keyword: keyword.trim(), page: 1, pageSize: 20 }),
        getGoodsList({ keyword: keyword.trim(), page: 1, pageSize: 20 }),
      ]);
      setCoaches(coachRes?.list || []);
      setGoods(goodsRes?.list || []);
    } catch {
      Taro.showToast({ title: '搜索失败', icon: 'none' });
    }
  };

  return (
    <View className="search">
      <View className="search__header">
        <Text className="search__back" onClick={() => Taro.navigateBack()}>←</Text>
        <Input className="search__input" placeholder="搜索陪玩师、游戏、服务..." focus value={keyword} onInput={(e) => setKeyword(e.detail.value)} onConfirm={doSearch} />
        <Text className="search__btn" onClick={doSearch}>搜索</Text>
      </View>

      {!searched ? (
        <View className="search__history">
          <View className="search__section-title">
            搜索历史 <Text className="search__clear">清空</Text>
          </View>
          <View className="search__tags">
            {history.map((h, i) => <Text key={i} className="search__tag" onClick={() => { setKeyword(h); }}>{h}</Text>)}
          </View>
          <View className="search__section-title" style={{ marginTop: '20px' }}>热门搜索</View>
          <View className="search__tags">
            {hotWords.map((w, i) => <Text key={i} className="search__hot-tag" onClick={() => { setKeyword(w); }}>{w}</Text>)}
          </View>
        </View>
      ) : (
        <View className="search__results">
          <View className="search__result-tabs">
            <Text className={`search__result-tab ${resultTab === 0 ? 'search__result-tab--active' : ''}`} onClick={() => setResultTab(0)}>陪玩师</Text>
            <Text className={`search__result-tab ${resultTab === 1 ? 'search__result-tab--active' : ''}`} onClick={() => setResultTab(1)}>商品</Text>
          </View>
          {resultTab === 0 ? (
            coaches.length === 0 ? <EmptyState text="未找到相关内容" /> : (
              <View className="search__grid">
                {coaches.map((c) => <CoachCard key={c.id} {...c} onClick={() => Taro.navigateTo({ url: `/pages/coach/detail?id=${c.id}` })} />)}
              </View>
            )
          ) : (
            goods.length === 0 ? <EmptyState text="未找到相关内容" /> : (
              <View className="search__list">
                {goods.map((g) => <GoodsCard key={g.id} {...g} onClick={() => Taro.navigateTo({ url: `/pages/goods/detail?id=${g.id}` })} />)}
              </View>
            )
          )}
        </View>
      )}
    </View>
  );
}
