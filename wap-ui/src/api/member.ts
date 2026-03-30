import { get, put, post } from './request';

export function getMemberInfo() {
  return get('/api/playapi/member/info');
}

export function updateMember(data: any) {
  return put('/api/playapi/member/update', data);
}

export function switchRole() {
  return post('/api/playapi/member/switch_role');
}

export function getBalanceLog(params: any) {
  return get('/api/playapi/member/balance_log', params);
}
