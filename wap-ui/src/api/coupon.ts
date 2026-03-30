import { get } from './request';

export function getCouponList(params: any) {
  return get('/api/playapi/coupon/list', params);
}

export function getUsableCoupons(amount: number) {
  return get('/api/playapi/coupon/usable', { amount });
}

export function receiveCoupon(id: string) {
  return get('/api/playapi/coupon/receive', { id });
}

export function getCouponCenter() {
  return get('/api/playapi/coupon/center');
}
