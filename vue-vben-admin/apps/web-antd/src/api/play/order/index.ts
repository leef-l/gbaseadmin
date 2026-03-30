import { requestClient } from '#/api/request';

import type {
  OrderItem,
  OrderListParams,
  OrderCreateParams,
  OrderUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/order';

/** 获取订单表列表 */
export function getOrderList(params: OrderListParams) {
  return requestClient.get<{ list: OrderItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取订单表详情 */
export function getOrderDetail(id: string) {
  return requestClient.get<OrderItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建订单表 */
export function createOrder(data: OrderCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新订单表 */
export function updateOrder(data: OrderUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除订单表 */
export function deleteOrder(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}

/** 变更订单状态 */
export function changeOrderStatus(data: { id: string; orderStatus: number; cancelReason?: string }) {
  return requestClient.post(`${PREFIX}/change-status`, data);
}
