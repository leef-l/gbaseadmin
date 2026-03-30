/** 订单表类型定义 */

/** 订单表项 */
export interface OrderItem {
  id: string;
  orderNo: string;
  memberID: string;
  coachID: string;
  shopID?: string;
  shopTitle?: string;
  goodsID: string;
  goodsTitle: string;
  goodsPrice: string;
  quantity?: number;
  totalAmount?: string;
  discountAmount?: string;
  couponAmount?: string;
  payAmount?: string;
  couponMemberID?: string;
  payType?: number;
  orderStatus?: number;
  payAt?: string;
  startAt?: string;
  finishAt?: string;
  cancelAt?: string;
  cancelReason?: string;
  remark?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** 订单表列表查询参数 */
export interface OrderListParams {
  pageNum: number;
  pageSize: number;
  payType?: number;
  orderStatus?: number;
}

/** 订单表创建参数 */
export interface OrderCreateParams {
  orderNo: string;
  memberID: string;
  coachID: string;
  shopID?: string;
  goodsID: string;
  goodsTitle: string;
  goodsPrice: string;
  quantity?: number;
  totalAmount?: string;
  discountAmount?: string;
  couponAmount?: string;
  payAmount?: string;
  couponMemberID?: string;
  payType?: number;
  orderStatus?: number;
  payAt?: string;
  startAt?: string;
  finishAt?: string;
  cancelAt?: string;
  cancelReason?: string;
  remark?: string;
}

/** 订单表更新参数 */
export interface OrderUpdateParams {
  id: string;
  orderNo: string;
  memberID: string;
  coachID: string;
  shopID?: string;
  goodsID: string;
  goodsTitle: string;
  goodsPrice: string;
  quantity?: number;
  totalAmount?: string;
  discountAmount?: string;
  couponAmount?: string;
  payAmount?: string;
  couponMemberID?: string;
  payType?: number;
  orderStatus?: number;
  payAt?: string;
  startAt?: string;
  finishAt?: string;
  cancelAt?: string;
  cancelReason?: string;
  remark?: string;
}
