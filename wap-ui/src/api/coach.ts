import { get, post, put } from './request';

export function getCoachList(params: any) {
  return get('/api/playapi/coach/list', params);
}

export function getCoachDetail(coachId: string) {
  return get('/api/playapi/coach/detail', { coachId });
}

export function applyCoach(data: any) {
  return post('/api/playapi/coach/apply', data);
}

export function getApplyStatus() {
  return get('/api/playapi/coach/apply_status');
}

export function setOnline(isOnline: number) {
  return put('/api/playapi/coach/online', { isOnline });
}

export function getMyGoodsList(params?: { status?: number; page?: number; pageSize?: number }) {
  return get('/api/playapi/coach/my_goods', params);
}

export function createGoods(data: any) {
  return post('/api/playapi/coach/goods/create', data);
}

export function updateGoods(data: any) {
  return put('/api/playapi/coach/goods/update', data);
}

export function updateGoodsStatus(data: { goodsId: string; status: number }) {
  return put('/api/playapi/coach/goods/status', data);
}

export function getIncome() {
  return get('/api/playapi/coach/income');
}

export function getCoachOrders(params?: { status?: number; page?: number; pageSize?: number }) {
  return get('/api/playapi/coach/orders', params);
}
