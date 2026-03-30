import { get, post } from './request';

export function getCoachList(params: any) {
  return get('/api/playapi/coach/list', params);
}

export function getCoachDetail(id: string) {
  return get(`/api/playapi/coach/detail`, { id });
}

export function applyCoach(data: any) {
  return post('/api/playapi/coach/apply', data);
}

export function getApplyStatus() {
  return get('/api/playapi/coach/apply_status');
}

export function getCoachOrders(params: any) {
  return get('/api/playapi/coach/orders', params);
}
