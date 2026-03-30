import { get, put, post } from './request';

export function getMemberInfo() {
  return get('/api/playapi/member/info');
}

export function updateMember(data: { nickname?: string; avatar?: string; gender?: number }) {
  return put('/api/playapi/member/update', data);
}

export function switchRole(role: 'member' | 'coach') {
  return post('/api/playapi/member/switch_role', { role });
}

export function getBalanceLog(params?: { type?: string; page?: number; pageSize?: number }) {
  return get('/api/playapi/member/balance_log', params);
}
