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
  userInfo: UserInfo | null;
  setToken: (token: string) => void;
  setUserInfo: (info: UserInfo) => void;
  logout: () => void;
  loadFromStorage: () => void;
}

export const useAuthStore = create<AuthState>((set) => ({
  token: '',
  userInfo: null,
  setToken: (token) => {
    Taro.setStorageSync('token', token);
    set({ token });
  },
  setUserInfo: (userInfo) => {
    Taro.setStorageSync('userInfo', JSON.stringify(userInfo));
    set({ userInfo });
  },
  logout: () => {
    Taro.removeStorageSync('token');
    Taro.removeStorageSync('userInfo');
    set({ token: '', userInfo: null });
  },
  loadFromStorage: () => {
    const token = Taro.getStorageSync('token') || '';
    const raw = Taro.getStorageSync('userInfo');
    const userInfo = raw ? JSON.parse(raw) : null;
    set({ token, userInfo });
  },
}));
