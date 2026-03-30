import { post } from './request';

export function pay(data: any) {
  return post('/api/playapi/payment/pay', data);
}
