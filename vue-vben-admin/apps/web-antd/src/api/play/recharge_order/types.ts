/** 充值订单表类型定义 */

/** 充值订单表项 */
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

/** 充值订单表列表查询参数 */
export interface RechargeOrderListParams {
  pageNum: number;
  pageSize: number;
  payType?: number;
  payStatus?: number;
}

/** 充值订单表创建参数 */
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

/** 充值订单表更新参数 */
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
