/** æ”¯ä»˜è®°å½•è¡¨类型定义 */

/** æ”¯ä»˜è®°å½•è¡¨项 */
export interface PaymentItem {
  id: string;
  orderID: string;
  memberID: string;
  paymentNo: string;
  tradeNo?: string;
  payType?: number;
  payAmount?: string;
  payStatus?: number;
  payAt?: string;
  refundAt?: string;
  refundAmount?: string;
  callbackContent?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** æ”¯ä»˜è®°å½•è¡¨列表查询参数 */
export interface PaymentListParams {
  pageNum: number;
  pageSize: number;
  payType?: number;
  payStatus?: number;
}

/** æ”¯ä»˜è®°å½•è¡¨创建参数 */
export interface PaymentCreateParams {
  orderID: string;
  memberID: string;
  paymentNo: string;
  tradeNo?: string;
  payType?: number;
  payAmount?: string;
  payStatus?: number;
  payAt?: string;
  refundAt?: string;
  refundAmount?: string;
  callbackContent?: string;
}

/** æ”¯ä»˜è®°å½•è¡¨更新参数 */
export interface PaymentUpdateParams {
  id: string;
  orderID: string;
  memberID: string;
  paymentNo: string;
  tradeNo?: string;
  payType?: number;
  payAmount?: string;
  payStatus?: number;
  payAt?: string;
  refundAt?: string;
  refundAmount?: string;
  callbackContent?: string;
}
