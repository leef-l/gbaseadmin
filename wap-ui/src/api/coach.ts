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

export function getCoachInfo() {
  return get('/api/playapi/coach/info');
}

export function getProfitLogList(params: any) {
  return get('/api/playapi/coach/profit_log', params);
}

export function getMyGoodsList(params: any) {
  return get('/api/playapi/coach/goods', params);
}

export function updateGoodsStatus(data: { id: string; status: number }) {
  return post('/api/playapi/coach/goods_status', data);
}
