import { requestClient } from '#/api/request';

const PREFIX = '/play/order';

/** 变更订单状态 */
export function changeOrderStatus(data: {
  id: string;
  orderStatus: number;
  cancelReason?: string;
}) {
  return requestClient.post(`${PREFIX}/change-status`, data);
}
