import { useState } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad } from '@tarojs/taro';
import { getCategoryList, getGoodsList } from '../../api/goods';
import './index.scss';

export default function CategoryPage() {
  const [categories, setCategories] = useState<any[]>([]);
  const [activeId, setActiveId] = useState('');
  const [goods, setGoods] = useState<any[]>([]);

  const loadGoods = async (categoryId: string) => {
    const res = await getGoodsList({ categoryId });
    setGoods(res?.list || []);
  };

  useLoad(async () => {
    const res = await getCategoryList();
    const list = res?.list || [];
    setCategories(list);
    if (list.length > 0) {
      setActiveId(list[0].id);
      loadGoods(list[0].id);
    }
  });

  const handleCategoryClick = (id: string) => {
    setActiveId(id);
    loadGoods(id);
  };

  return (
    <View className="category">
      <View className="category__left">
        {categories.map((c) => (
          <View
            key={c.id}
            className={`category__item ${activeId === c.id ? 'category__item--active' : ''}`}
            onClick={() => handleCategoryClick(c.id)}
          >
            {c.name}
          </View>
        ))}
      </View>
      <View className="category__right">
        <View className="category__header">
          <Text className="category__name">{categories.find(c => c.id === activeId)?.name}</Text>
          <Text className="category__count">{goods.length}个服务</Text>
        </View>
        <View className="category__grid">
          {goods.map((g) => (
            <View key={g.id} className="category__goods" onClick={() => Taro.navigateTo({ url: `/pages/goods/detail?id=${g.id}` })}>
              <View className="category__cover" />
              <View className="category__info">
                <Text className="category__title">{g.title}</Text>
                <Text className="price">¥{(g.price / 100).toFixed(2)}</Text>
              </View>
            </View>
          ))}
        </View>
      </View>
    </View>
  );
}
