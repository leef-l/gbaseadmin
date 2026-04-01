import { create } from 'zustand';
import Taro from '@tarojs/taro';

interface UserInfo {
  id: string;
  nickname: string;
  avatar: string;
  phone: string;
  levelId: string;
  levelName: string;
  levelTitle?: string;
  balance: number;
  isCoach: number;
  currentRole: 'member' | 'coach';
  exp?: number;
  couponCount?: number;
}

interface AuthState {
  token: string;
  refreshToken: string;
  userInfo: UserInfo | null;
  setToken: (token: string) => void;
  setRefreshToken: (refreshToken: string) => void;
  setUserInfo: (info: UserInfo) => void;
  logout: () => void;
  loadFromStorage: () => void;
}

export const useAuthStore = create<AuthState>((set) => ({
  token: '',
  refreshToken: '',
  userInfo: null,
  setToken: (token) => {
    Taro.setStorageSync('token', token);
    set({ token });
  },
  setRefreshToken: (refreshToken) => {
    Taro.setStorageSync('refreshToken', refreshToken);
    set({ refreshToken });
  },
  setUserInfo: (userInfo) => {
    Taro.setStorageSync('userInfo', JSON.stringify(userInfo));
    set({ userInfo });
  },
  logout: () => {
    Taro.removeStorageSync('token');
    Taro.removeStorageSync('refreshToken');
    Taro.removeStorageSync('userInfo');
    set({ token: '', refreshToken: '', userInfo: null });
  },
  loadFromStorage: () => {
    const token = Taro.getStorageSync('token') || '';
    const refreshToken = Taro.getStorageSync('refreshToken') || '';
    const raw = Taro.getStorageSync('userInfo');
    const userInfo = raw ? JSON.parse(raw) : null;
    set({ token, refreshToken, userInfo });
  },
}));
