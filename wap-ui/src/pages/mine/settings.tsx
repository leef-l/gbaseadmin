import { View, Text } from '@tarojs/components';
import Taro from '@tarojs/taro';
import { useAuthStore } from '../../store/auth';
import './settings.scss';

export default function SettingsPage() {
  const { token, logout } = useAuthStore();
  const isLoggedIn = !!token;

  const handleClearCache = () => {
    Taro.showModal({
      title: '提示',
      content: '确定清除缓存？',
      success: (res) => {
        if (res.confirm) {
          Taro.clearStorageSync();
          Taro.showToast({ title: '缓存已清除', icon: 'success' });
        }
      },
    });
  };

  const handleLogout = () => {
    Taro.showModal({
      title: '提示',
      content: '确定退出登录？',
      success: (res) => {
        if (res.confirm) {
          logout();
          Taro.redirectTo({ url: '/pages/login/index' });
        }
      },
    });
  };

  const menuItems = [
    { text: '修改密码', onClick: () => Taro.showToast({ title: '暂未开放', icon: 'none' }) },
    { text: '清除缓存', onClick: handleClearCache },
    { text: '关于我们', onClick: () => Taro.showToast({ title: '陪玩平台 v1.0.0', icon: 'none' }) },
  ];

  return (
    <View className="settings">
      <View className="settings__menu card">
        {menuItems.map((m, i) => (
          <View key={i} className="settings__item" onClick={m.onClick}>
            <Text>{m.text}</Text>
            <Text className="settings__arrow">&gt;</Text>
          </View>
        ))}
      </View>

      {isLoggedIn && (
        <View className="settings__logout" onClick={handleLogout}>
          <Text>退出登录</Text>
        </View>
      )}
    </View>
  );
}
