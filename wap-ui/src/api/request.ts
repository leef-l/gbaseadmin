import Taro from '@tarojs/taro';
import { useAuthStore } from '../store/auth';

const BASE_URL = process.env.TARO_APP_API || '';

interface RequestOptions {
  url: string;
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE';
  data?: any;
  header?: Record<string, string>;
}

interface ApiResponse<T = any> {
  code: number;
  message: string;
  data: T;
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
    useAuthStore.getState().logout();
    Taro.navigateTo({ url: '/pages/login/index' });
    return Promise.reject(new Error('未登录'));
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
