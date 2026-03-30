import { useState } from 'react';
import { View, Text } from '@tarojs/components';
import Taro, { useLoad } from '@tarojs/taro';
import { getMessageList, markAllRead } from '../../api/message';
import EmptyState from '../../components/EmptyState';
import './index.scss';

const iconMap: Record<string, { bg: string; icon: string }> = {
  order: { bg: '#6C5CE7', icon: '🛍️' },
  system: { bg: '#0984e3', icon: '🔔' },
  activity: { bg: '#d63031', icon: '🎁' },
};

export default function MessagePage() {
  const [messages, setMessages] = useState<any[]>([]);

  const fetchMessages = async () => {
    const res = await getMessageList();
    setMessages(res?.list || []);
  };

  useLoad(() => {
    fetchMessages();
  });

  const handleMarkAllRead = async () => {
    await markAllRead();
    fetchMessages();
  };

  return (
    <View className="message">
      <View className="message__header">
        <Text className="message__title">消息</Text>
        <Text className="message__read-all" onClick={handleMarkAllRead}>全部已读</Text>
      </View>
      {messages.length === 0 ? <EmptyState text="暂无消息" /> : (
        <View className="message__list">
          {messages.map((m) => {
            const ic = iconMap[m.type] || iconMap.system;
            return (
              <View key={m.id} className="message__item card">
                <View className="message__icon" style={{ background: ic.bg }}>{ic.icon}</View>
                <View className="message__content">
                  <Text className="message__name">{m.title}</Text>
                  <Text className="message__desc">{m.desc}</Text>
                </View>
                <View className="message__right">
                  <Text className="message__time">{m.time}</Text>
                  {m.unread && <View className="message__dot" />}
                </View>
              </View>
            );
          })}
        </View>
      )}
    </View>
  );
}
