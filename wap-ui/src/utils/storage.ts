import Taro from '@tarojs/taro';

export function getStorage<T = any>(key: string): T | null {
  try {
    const val = Taro.getStorageSync(key);
    return val ? JSON.parse(val) : null;
  } catch {
    return null;
  }
}

export function setStorage(key: string, value: any): void {
  Taro.setStorageSync(key, JSON.stringify(value));
}

export function removeStorage(key: string): void {
  Taro.removeStorageSync(key);
}
