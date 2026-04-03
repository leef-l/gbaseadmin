import { useCallback, useState } from 'react';
import { View, Text, Image } from '@tarojs/components';
import Taro, { useLoad } from '@tarojs/taro';
import { getCategoryList, getGoodsList } from '../../api/goods';
import EmptyState from '../../components/EmptyState';
import './index.scss';

interface CategoryItem {
  categoryId: string;
  name: string;
}

interface GoodsItem {
  goodsId: string;
  title: string;
  price: number;
  coverImage?: string;
}

export default function CategoryPage() {
  const [categories, setCategories] = useState<CategoryItem[]>([]);
  const [activeId, setActiveId] = useState('');
  const [goods, setGoods] = useState<GoodsItem[]>([]);

  const loadGoods = useCallback(async (categoryId: string) => {
    try {
      const res = await getGoodsList({ categoryId });
      setGoods(res?.list || []);
    } catch {
      setGoods([]);
      Taro.showToast({ title: '加载服务失败', icon: 'none' });
    }
  }, []);

  useLoad(() => {
    void (async () => {
      try {
        const res = await getCategoryList();
        const list = (res?.list || []).map((item: any) => ({
          categoryId: item.categoryId || item.id,
          name: item.name,
        }));
        setCategories(list);
        if (list.length === 0) {
          setActiveId('');
          setGoods([]);
          return;
        }
        setActiveId(list[0].categoryId);
        void loadGoods(list[0].categoryId);
      } catch {
        setCategories([]);
        setActiveId('');
        setGoods([]);
        Taro.showToast({ title: '加载分类失败', icon: 'none' });
      }
    })();
  });

  const handleCategoryClick = (id: string) => {
    setActiveId(id);
    void loadGoods(id);
  };

  return (
    <View className="category">
      <View className="category__left">
        {categories.length === 0 ? (
          <View className="category__left-empty">暂无分类</View>
        ) : (
          categories.map((c) => (
            <View
              key={c.categoryId}
              className={`category__item ${activeId === c.categoryId ? 'category__item--active' : ''}`}
              onClick={() => handleCategoryClick(c.categoryId)}
            >
              {c.name}
            </View>
          ))
        )}
      </View>
      <View className="category__right">
        {categories.length === 0 ? (
          <View className="category__empty">
            <EmptyState text="暂无分类" />
          </View>
        ) : (
          <>
            <View className="category__header">
              <Text className="category__name">{categories.find(c => c.categoryId === activeId)?.name}</Text>
              <Text className="category__count">{goods.length}个服务</Text>
            </View>
            {goods.length === 0 ? (
              <View className="category__empty">
                <EmptyState text="该分类暂无服务" />
              </View>
            ) : (
              <View className="category__grid">
                {goods.map((g) => (
                  <View key={g.goodsId} className="category__goods" onClick={() => Taro.navigateTo({ url: `/pages/goods/detail?id=${g.goodsId}` })}>
                    <View className="category__cover">
                      {g.coverImage && <Image src={g.coverImage} className="category__cover-img" mode="aspectFill" />}
                    </View>
                    <View className="category__info">
                      <Text className="category__title">{g.title}</Text>
                      <Text className="price">¥{(g.price / 100).toFixed(2)}</Text>
                    </View>
                  </View>
                ))}
              </View>
            )}
          </>
        )}
      </View>
    </View>
  );
}
