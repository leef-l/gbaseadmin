import { get, post } from './request';

export function createOrder(data: any) {
  return post('/api/playapi/order/create', data);
}

export function getOrderList(params: any) {
  return get('/api/playapi/order/list', params);
}

export function getOrderDetail(id: string) {
  return get('/api/playapi/order/detail', { id });
}

export function cancelOrder(id: string) {
  return post('/api/playapi/order/cancel', { id });
}

export function acceptOrder(id: string) {
  return post('/api/playapi/order/accept', { id });
}

export function finishOrder(id: string) {
  return post('/api/playapi/order/finish', { id });
}
