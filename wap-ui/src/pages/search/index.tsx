import { useState, useRef, useCallback } from 'react';
import { View, Text, Input } from '@tarojs/components';
import Taro, { useReachBottom } from '@tarojs/taro';
import { search } from '../../api/goods';
import CoachCard from '../../components/CoachCard';
import GoodsCard from '../../components/GoodsCard';
import EmptyState from '../../components/EmptyState';
import LoadMore from '../../components/LoadMore';
import './index.scss';

const HISTORY_KEY = 'searchHistory';
const MAX_HISTORY = 10;
const HOT_WORDS = ['游戏陪玩', '看电影', '唱歌', '哄睡'];
const PAGE_SIZE = 20;

export default function SearchPage() {
  const [keyword, setKeyword] = useState('');
  const [searched, setSearched] = useState(false);
  const [resultTab, setResultTab] = useState(0);
  const [coaches, setCoaches] = useState<any[]>([]);
  const [goods, setGoods] = useState<any[]>([]);
  const [coachHasMore, setCoachHasMore] = useState(false);
  const [goodsHasMore, setGoodsHasMore] = useState(false);
  const [loading, setLoading] = useState(false);
  const coachPageRef = useRef(1);
  const goodsPageRef = useRef(1);
  const currentKeywordRef = useRef('');

  const [history, setHistory] = useState<string[]>(() => {
    try { return Taro.getStorageSync(HISTORY_KEY) || []; } catch { return []; }
  });
  const hotWords = HOT_WORDS;

  const saveHistory = useCallback((kw: string) => {
    setHistory((prev) => {
      const next = [kw, ...prev.filter((h) => h !== kw)].slice(0, MAX_HISTORY);
      try { Taro.setStorageSync(HISTORY_KEY, next); } catch { /* ignore */ }
      return next;
    });
  }, []);

  const clearHistory = useCallback(() => {
    setHistory([]);
    try { Taro.removeStorageSync(HISTORY_KEY); } catch { /* ignore */ }
  }, []);

  const doSearch = async (searchWord?: string) => {
    const kw = (searchWord ?? keyword).trim();
    if (!kw) return;
    setKeyword(kw);
    setSearched(true);
    saveHistory(kw);
    currentKeywordRef.current = kw;
    coachPageRef.current = 1;
    goodsPageRef.current = 1;
    setLoading(true);
    try {
      const [coachRes, goodsRes] = await Promise.all([
        search({ keyword: kw, type: 'coach', page: 1, pageSize: PAGE_SIZE }),
        search({ keyword: kw, type: 'goods', page: 1, pageSize: PAGE_SIZE }),
      ]);
      const coachRows = (coachRes as any)?.list || [];
      const goodsRows = (goodsRes as any)?.list || [];
      setCoaches(coachRows);
      setGoods(goodsRows);
      setCoachHasMore(coachRows.length >= PAGE_SIZE);
      setGoodsHasMore(goodsRows.length >= PAGE_SIZE);
      coachPageRef.current = 2;
      goodsPageRef.current = 2;
    } catch {
      Taro.showToast({ title: '搜索失败', icon: 'none' });
    } finally {
      setLoading(false);
    }
  };

  const loadMoreCoaches = useCallback(async () => {
    if (loading || !coachHasMore) return;
    setLoading(true);
    try {
      const res = await search({ keyword: currentKeywordRef.current, type: 'coach', page: coachPageRef.current, pageSize: PAGE_SIZE });
      const rows = (res as any)?.list || [];
      setCoaches((prev) => [...prev, ...rows]);
      setCoachHasMore(rows.length >= PAGE_SIZE);
      coachPageRef.current += 1;
    } finally {
      setLoading(false);
    }
  }, [loading, coachHasMore]);

  const loadMoreGoods = useCallback(async () => {
    if (loading || !goodsHasMore) return;
    setLoading(true);
    try {
      const res = await search({ keyword: currentKeywordRef.current, type: 'goods', page: goodsPageRef.current, pageSize: PAGE_SIZE });
      const rows = (res as any)?.list || [];
      setGoods((prev) => [...prev, ...rows]);
      setGoodsHasMore(rows.length >= PAGE_SIZE);
      goodsPageRef.current += 1;
    } finally {
      setLoading(false);
    }
  }, [loading, goodsHasMore]);

  useReachBottom(() => {
    if (!searched) return;
    if (resultTab === 0) {
      loadMoreCoaches();
    } else {
      loadMoreGoods();
    }
  });

  return (
    <View className="search">
      <View className="search__header">
        <Text className="search__back" onClick={() => Taro.navigateBack()}>←</Text>
        <Input className="search__input" placeholder="搜索陪玩师、游戏、服务..." focus value={keyword} onInput={(e) => setKeyword(e.detail.value)} onConfirm={() => doSearch()} />
        <Text className="search__btn" onClick={() => doSearch()}>搜索</Text>
      </View>

      {!searched ? (
        <View className="search__history">
          <View className="search__section-title">
            搜索历史 <Text className="search__clear" onClick={clearHistory}>清空</Text>
          </View>
          <View className="search__tags">
            {history.map((h, i) => <Text key={i} className="search__tag" onClick={() => doSearch(h)}>{h}</Text>)}
          </View>
          <View className="search__section-title" style={{ marginTop: '20px' }}>热门搜索</View>
          <View className="search__tags">
            {hotWords.map((w, i) => <Text key={i} className="search__hot-tag" onClick={() => doSearch(w)}>{w}</Text>)}
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
                <LoadMore hasMore={coachHasMore} />
              </View>
            )
          ) : (
            goods.length === 0 ? <EmptyState text="未找到相关内容" /> : (
              <View className="search__list">
                {goods.map((g) => <GoodsCard key={g.id} {...g} onClick={() => Taro.navigateTo({ url: `/pages/goods/detail?id=${g.id}` })} />)}
                <LoadMore hasMore={goodsHasMore} />
              </View>
            )
          )}
        </View>
      )}
    </View>
  );
}
