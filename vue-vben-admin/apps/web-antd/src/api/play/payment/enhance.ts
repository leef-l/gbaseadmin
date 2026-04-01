import { requestClient } from '#/api/request';

const PREFIX = '/play/payment';

/** 管理员代客余额支付 */
export function payByBalance(data: { orderID: string; memberID: string }) {
  return requestClient.post(`${PREFIX}/pay_by_balance`, data);
}
