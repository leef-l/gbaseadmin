import { requestClient } from '#/api/request';

import type {
  PaymentItem,
  PaymentListParams,
  PaymentCreateParams,
  PaymentUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/payment';

/** 获取支付记录表列表 */
export function getPaymentList(params: PaymentListParams) {
  return requestClient.get<{ list: PaymentItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取支付记录表详情 */
export function getPaymentDetail(id: string) {
  return requestClient.get<PaymentItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建支付记录表 */
export function createPayment(data: PaymentCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新支付记录表 */
export function updatePayment(data: PaymentUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除支付记录表 */
export function deletePayment(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
