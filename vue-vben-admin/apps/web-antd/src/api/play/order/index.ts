import { requestClient } from '#/api/request';

import type {
  OrderItem,
  OrderListParams,
  OrderCreateParams,
  OrderUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/order';

/** 获取è®¢å•è¡¨列表 */
export function getOrderList(params: OrderListParams) {
  return requestClient.get<{ list: OrderItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取è®¢å•è¡¨详情 */
export function getOrderDetail(id: string) {
  return requestClient.get<OrderItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建è®¢å•è¡¨ */
export function createOrder(data: OrderCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新è®¢å•è¡¨ */
export function updateOrder(data: OrderUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除è®¢å•è¡¨ */
export function deleteOrder(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
