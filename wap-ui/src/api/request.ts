import Taro from '@tarojs/taro';
import { useAuthStore } from '../store/auth';

const BASE_URL = process.env.TARO_APP_API || '';

interface RequestOptions {
  url: string;
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE';
  data?: any;
  header?: Record<string, string>;
  /** 内部标记：此次请求是刷新 token 本身，不再递归刷新 */
  _isRefresh?: boolean;
}

interface ApiResponse<T = any> {
  code: number;
  message: string;
  data: T;
}

interface RefreshTokenResult {
  token: string;
  refreshToken: string;
}

/** 正在进行中的刷新 Promise，避免并发时多次刷新 */
let refreshingPromise: Promise<string> | null = null;

async function doRefreshToken(): Promise<string> {
  const { refreshToken, setToken, setRefreshToken, logout } = useAuthStore.getState();
  if (!refreshToken) {
    logout();
    Taro.navigateTo({ url: '/pages/login/index' });
    return Promise.reject(new Error('无刷新凭证'));
  }

  const res = await Taro.request({
    url: `${BASE_URL}/api/playapi/auth/refresh_token`,
    method: 'POST',
    data: { refreshToken },
    header: { 'Content-Type': 'application/json' },
  });

  const body = res.data as ApiResponse<RefreshTokenResult>;
  if (body.code !== 0 || !body.data?.token) {
    logout();
    Taro.navigateTo({ url: '/pages/login/index' });
    return Promise.reject(new Error('刷新 token 失败'));
  }

  setToken(body.data.token);
  if (body.data.refreshToken) {
    setRefreshToken(body.data.refreshToken);
  }
  return body.data.token;
}

export async function request<T = any>(options: RequestOptions): Promise<T> {
  const token = useAuthStore.getState().token;
  const header: Record<string, string> = {
    'Content-Type': 'application/json',
    ...options.header,
  };
  if (token) {
    header['Authorization'] = `Bearer ${token}`;
  }

  const res = await Taro.request({
    url: `${BASE_URL}${options.url}`,
    method: options.method || 'GET',
    data: options.data,
    header,
  });

  const body = res.data as ApiResponse<T>;

  if (body.code === 401) {
    // 刷新 token 请求本身 401，直接退出，不再递归
    if (options._isRefresh) {
      useAuthStore.getState().logout();
      Taro.navigateTo({ url: '/pages/login/index' });
      return Promise.reject(new Error('登录已过期'));
    }

    // 多个并发请求同时 401 时，复用同一个刷新 Promise
    if (!refreshingPromise) {
      refreshingPromise = doRefreshToken().finally(() => {
        refreshingPromise = null;
      });
    }

    let newToken: string;
    try {
      newToken = await refreshingPromise;
    } catch {
      return Promise.reject(new Error('登录已过期'));
    }

    // 用新 token 重试原请求（标记 _isRefresh 防止再次触发刷新）
    return request<T>({
      ...options,
      header: {
        ...options.header,
        Authorization: `Bearer ${newToken}`,
      },
      _isRefresh: true,
    });
  }

  if (body.code !== 0) {
    Taro.showToast({ title: body.message || '请求失败', icon: 'none' });
    return Promise.reject(new Error(body.message));
  }

  return body.data;
}

export function get<T = any>(url: string, data?: any) {
  return request<T>({ url, method: 'GET', data });
}

export function post<T = any>(url: string, data?: any) {
  return request<T>({ url, method: 'POST', data });
}

export function put<T = any>(url: string, data?: any) {
  return request<T>({ url, method: 'PUT', data });
}

export function del<T = any>(url: string, data?: any) {
  return request<T>({ url, method: 'DELETE', data });
}
