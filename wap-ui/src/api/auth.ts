import { post } from './request';

export function sendCode(phone: string) {
  return post('/api/playapi/auth/send_code', { phone });
}

export function login(phone: string, code: string) {
  return post('/api/playapi/auth/login', { phone, code });
}

export function wxLogin(code: string) {
  return post('/api/playapi/auth/wx_login', { code });
}
