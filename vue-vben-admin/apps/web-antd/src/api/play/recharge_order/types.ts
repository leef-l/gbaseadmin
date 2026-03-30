/** å……å€¼è®¢å•è¡¨类型定义 */

/** å……å€¼è®¢å•è¡¨项 */
export interface RechargeOrderItem {
  id: string;
  orderNo: string;
  memberID: string;
  rechargePlanID: string;
  rechargePlanTitle?: string;
  amount: string;
  giftAmount?: string;
  payType?: number;
  tradeNo?: string;
  payStatus?: number;
  payAt?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** å……å€¼è®¢å•è¡¨列表查询参数 */
export interface RechargeOrderListParams {
  pageNum: number;
  pageSize: number;
  payType?: number;
  payStatus?: number;
}

/** å……å€¼è®¢å•è¡¨创建参数 */
export interface RechargeOrderCreateParams {
  orderNo: string;
  memberID: string;
  rechargePlanID: string;
  amount: string;
  giftAmount?: string;
  payType?: number;
  tradeNo?: string;
  payStatus?: number;
  payAt?: string;
}

/** å……å€¼è®¢å•è¡¨更新参数 */
export interface RechargeOrderUpdateParams {
  id: string;
  orderNo: string;
  memberID: string;
  rechargePlanID: string;
  amount: string;
  giftAmount?: string;
  payType?: number;
  tradeNo?: string;
  payStatus?: number;
  payAt?: string;
}
