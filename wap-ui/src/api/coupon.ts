import { get, post } from './request';

export function getAvailableCoupons(params?: { page?: number; pageSize?: number }) {
  return get('/api/playapi/coupon/available', params);
}

export function receiveCoupon(couponId: string) {
  return post('/api/playapi/coupon/receive', { couponId });
}

export function getMyCoupons(params?: { status?: number; page?: number; pageSize?: number }) {
  return get('/api/playapi/coupon/mine', params);
}

export function getUsableCoupons(orderAmount: number) {
  return get('/api/playapi/coupon/usable', { orderAmount });
}
