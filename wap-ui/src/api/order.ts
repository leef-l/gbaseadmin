import { get, post } from './request';

export function createOrder(data: {
  goodsId: string;
  quantity: number;
  couponMemberId?: string;
  remark?: string;
}) {
  return post('/api/playapi/order/create', data);
}

export function getOrderList(params?: { status?: number; page?: number; pageSize?: number }) {
  return get('/api/playapi/order/list', params);
}

export function getOrderDetail(orderId: string) {
  return get('/api/playapi/order/detail', { orderId });
}

export function cancelOrder(orderId: string, cancelReason?: string) {
  return post('/api/playapi/order/cancel', { orderId, cancelReason });
}

export function refundOrder(orderId: string, refundReason: string) {
  return post('/api/playapi/order/refund', { orderId, refundReason });
}

export function acceptOrder(orderId: string) {
  return post('/api/playapi/order/accept', { orderId });
}

export function finishOrder(orderId: string) {
  return post('/api/playapi/order/finish', { orderId });
}
