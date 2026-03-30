import { requestClient } from '#/api/request';

import type {
  PaymentItem,
  PaymentListParams,
  PaymentCreateParams,
  PaymentUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/payment';

/** 获取æ”¯ä»˜è®°å½•è¡¨列表 */
export function getPaymentList(params: PaymentListParams) {
  return requestClient.get<{ list: PaymentItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取æ”¯ä»˜è®°å½•è¡¨详情 */
export function getPaymentDetail(id: string) {
  return requestClient.get<PaymentItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建æ”¯ä»˜è®°å½•è¡¨ */
export function createPayment(data: PaymentCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新æ”¯ä»˜è®°å½•è¡¨ */
export function updatePayment(data: PaymentUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除æ”¯ä»˜è®°å½•è¡¨ */
export function deletePayment(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
