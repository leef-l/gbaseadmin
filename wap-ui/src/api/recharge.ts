import { get, post } from './request';

export function getRechargePlans() {
  return get('/api/playapi/recharge/plans');
}

export function createRecharge(data: any) {
  return post('/api/playapi/recharge/create', data);
}
