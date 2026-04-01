import { View, Text, Image } from '@tarojs/components';
import Taro from '@tarojs/taro';
import { useAuthStore } from '../../store/auth';
import './index.scss';

const orderShortcuts = [
  { icon: '💰', text: '待支付', status: 0 },
  { icon: '⏳', text: '进行中', status: 2 },
  { icon: '✅', text: '已完成', status: 3 },
  { icon: '↩️', text: '退款', status: 5 },
];

const menuItems = [
  { text: '充值中心', url: '/pages/recharge/index' },
  { text: '领券中心', url: '/pages/coupon/center' },
  { text: '我的优惠券', url: '/pages/coupon/list' },
  { text: '余额明细', url: '/pages/mine/balance' },
  { text: '我的评价', url: '' },
];

export default function MinePage() {
  const { userInfo, token } = useAuthStore();
  const isLoggedIn = !!token;

  const goLogin = () => {
    if (!isLoggedIn) Taro.navigateTo({ url: '/pages/login/index' });
  };

  return (
    <View className="mine">
      <View className="mine__header">
        <View className="mine__user" onClick={goLogin}>
          <View className="mine__avatar">
            {userInfo?.avatar
              ? <Image src={userInfo.avatar} className="mine__avatar-img" mode="aspectFill" />
              : <Text>👤</Text>}
          </View>
          <View className="mine__info">
            <Text className="mine__name">
              {isLoggedIn ? userInfo?.nickname || '用户' : '点击登录'}
              {isLoggedIn && userInfo?.levelTitle && <Text className="mine__level">{userInfo.levelTitle}</Text>}
            </Text>
            {isLoggedIn && userInfo?.phone && <Text className="mine__phone">{userInfo.phone}</Text>}
          </View>
          <Text className="mine__settings" onClick={(e) => { e.stopPropagation(); Taro.navigateTo({ url: '/pages/mine/settings' }); }}>⚙️</Text>
        </View>
        {isLoggedIn && userInfo?.isCoach === 1 && (
          <Text className="mine__role">{userInfo.currentRole === 'coach' ? '陪玩师模式' : '会员模式'}</Text>
        )}
      </View>

      <View className="mine__assets card">
        <View onClick={() => isLoggedIn && Taro.navigateTo({ url: '/pages/mine/balance' })}>
          <Text className="mine__asset-value">¥{isLoggedIn ? ((userInfo?.balance || 0) / 100).toFixed(2) : '0.00'}</Text>
          <Text className="mine__asset-label">余额</Text>
        </View>
        <View onClick={() => isLoggedIn && Taro.navigateTo({ url: '/pages/coupon/list' })}>
          {/* couponCount 字段由后端 getMemberInfo 返回 */}
          <Text className="mine__asset-value">{isLoggedIn ? (userInfo?.couponCount ?? 0) : 0}</Text>
          <Text className="mine__asset-label">优惠券</Text>
        </View>
        <View>
          <Text className="mine__asset-value">{isLoggedIn ? (userInfo?.exp ?? 0) : 0}</Text>
          <Text className="mine__asset-label">经验值</Text>
        </View>
      </View>

      {/* 订单快捷入口 */}
      <View className="mine__orders card">
        <View className="mine__orders-header">
          <Text className="mine__orders-title">我的订单</Text>
          <Text className="mine__orders-more" onClick={() => Taro.navigateTo({ url: '/pages/order/list' })}>查看全部 &gt;</Text>
        </View>
        <View className="mine__orders-grid">
          {orderShortcuts.map((o, i) => (
            <View key={i} onClick={() => Taro.navigateTo({ url: `/pages/order/list?status=${o.status}` })}>
              <Text className="mine__order-icon">{o.icon}</Text>
              <Text className="mine__order-text">{o.text}</Text>
            </View>
          ))}
        </View>
      </View>

      {/* 功能菜单 */}
      <View className="mine__menu card">
        {menuItems.map((m, i) => (
          <View key={i} className="mine__menu-item" onClick={() => m.url && Taro.navigateTo({ url: m.url })}>
            <Text>{m.text}</Text>
            <Text className="mine__menu-arrow">&gt;</Text>
          </View>
        ))}
        {isLoggedIn && userInfo?.isCoach === 0 && (
          <View className="mine__menu-item" onClick={() => Taro.navigateTo({ url: '/pages/coach/apply' })}>
            <Text>申请成为陪玩师</Text>
            <Text className="mine__menu-arrow">&gt;</Text>
          </View>
        )}
        {isLoggedIn && userInfo?.isCoach === 1 && userInfo.currentRole === 'coach' && (
          <View className="mine__menu-item" onClick={() => Taro.navigateTo({ url: '/pages/workspace/index' })}>
            <Text>陪玩师工作台</Text>
            <Text className="mine__menu-arrow">&gt;</Text>
          </View>
        )}
        <View className="mine__menu-item" onClick={() => Taro.navigateTo({ url: '/pages/mine/settings' })}>
          <Text>设置</Text>
          <Text className="mine__menu-arrow">&gt;</Text>
        </View>
      </View>
    </View>
  );
}