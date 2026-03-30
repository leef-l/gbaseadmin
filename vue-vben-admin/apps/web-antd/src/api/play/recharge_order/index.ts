import { requestClient } from '#/api/request';

import type {
  RechargeOrderItem,
  RechargeOrderListParams,
  RechargeOrderCreateParams,
  RechargeOrderUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/recharge_order';

/** 获取å……å€¼è®¢å•è¡¨列表 */
export function getRechargeOrderList(params: RechargeOrderListParams) {
  return requestClient.get<{ list: RechargeOrderItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取å……å€¼è®¢å•è¡¨详情 */
export function getRechargeOrderDetail(id: string) {
  return requestClient.get<RechargeOrderItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建å……å€¼è®¢å•è¡¨ */
export function createRechargeOrder(data: RechargeOrderCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新å……å€¼è®¢å•è¡¨ */
export function updateRechargeOrder(data: RechargeOrderUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除å……å€¼è®¢å•è¡¨ */
export function deleteRechargeOrder(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
