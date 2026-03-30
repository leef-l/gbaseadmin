import { requestClient } from '#/api/request';

import type {
  RechargeOrderItem,
  RechargeOrderListParams,
  RechargeOrderCreateParams,
  RechargeOrderUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/recharge_order';

/** 获取充值订单表列表 */
export function getRechargeOrderList(params: RechargeOrderListParams) {
  return requestClient.get<{ list: RechargeOrderItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取充值订单表详情 */
export function getRechargeOrderDetail(id: string) {
  return requestClient.get<RechargeOrderItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建充值订单表 */
export function createRechargeOrder(data: RechargeOrderCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新充值订单表 */
export function updateRechargeOrder(data: RechargeOrderUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除充值订单表 */
export function deleteRechargeOrder(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
