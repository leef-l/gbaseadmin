/** 支付记录表类型定义 */

/** 支付记录表项 */
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

/** 支付记录表列表查询参数 */
export interface PaymentListParams {
  pageNum: number;
  pageSize: number;
  payType?: number;
  payStatus?: number;
}

/** 支付记录表创建参数 */
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

/** 支付记录表更新参数 */
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
