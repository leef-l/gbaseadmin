import { post } from './request';

export function sendCode(phone: string, scene: string = 'login') {
  return post('/api/playapi/auth/send_code', { phone, scene });
}

export function login(phone: string, code: string) {
  return post('/api/playapi/auth/login', { phone, code });
}

export function wxLogin(code: string) {
  return post('/api/playapi/auth/wx_login', { code });
}

export function alipayLogin(authCode: string) {
  return post('/api/playapi/auth/alipay_login', { authCode });
}

export function refreshToken(refreshToken: string) {
  return post('/api/playapi/auth/refresh_token', { refreshToken });
}
