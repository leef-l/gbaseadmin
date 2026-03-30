/** è®¢å•è¡¨类型定义 */

/** è®¢å•è¡¨项 */
export interface OrderItem {
  id: string;
  orderNo: string;
  memberID: string;
  coachID: string;
  shopID?: string;
  shopTitle?: string;
  goodsID: string;
  goodsTitle?: string;
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

/** è®¢å•è¡¨列表查询参数 */
export interface OrderListParams {
  pageNum: number;
  pageSize: number;
  payType?: number;
  orderStatus?: number;
}

/** è®¢å•è¡¨创建参数 */
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

/** è®¢å•è¡¨更新参数 */
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
