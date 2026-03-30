import { useState } from 'react';
import { View, Text, Image } from '@tarojs/components';
import Taro, { useLoad } from '@tarojs/taro';
import './index.scss';

const mockCategories = [
  { id: '1', name: '游戏陪玩' },
  { id: '2', name: '语音聊天' },
  { id: '3', name: '看电影' },
  { id: '4', name: '唱歌' },
  { id: '5', name: '叫醒哄睡' },
];

export default function CategoryPage() {
  const [categories, setCategories] = useState<any[]>([]);
  const [activeId, setActiveId] = useState('');
  const [goods, setGoods] = useState<any[]>([]);

  useLoad(() => {
    // TODO: 接入真实API
    setCategories(mockCategories);
    setActiveId(mockCategories[0]?.id || '');
    setGoods([
      { id: '1', title: '王者荣耀上分', price: 3000 },
      { id: '2', title: '英雄联盟双排', price: 2500 },
      { id: '3', title: '和平精英吃鸡', price: 2000 },
      { id: '4', title: '原神深渊代打', price: 5000 },
    ]);
  });

  const handleCategoryClick = (id: string) => {
    setActiveId(id);
    // TODO: 根据分类加载商品
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
