import Taro from '@tarojs/taro';
import { useAuthStore } from '../store/auth';

export function isLoggedIn(): boolean {
  return !!useAuthStore.getState().token;
}

export function requireAuth(): boolean {
  if (!isLoggedIn()) {
    Taro.navigateTo({ url: '/pages/login/index' });
    return false;
  }
  return true;
}
