/** 会员优惠券表类型定义 */

/** 会员优惠券表项 */
export interface CouponMemberItem {
  id: string;
  couponID: string;
  couponTitle?: string;
  memberID: string;
  orderID?: string;
  useStatus?: number;
  claimAt?: string;
  useAt?: string;
  expireAt?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** 会员优惠券表列表查询参数 */
export interface CouponMemberListParams {
  pageNum: number;
  pageSize: number;
  useStatus?: number;
}

/** 会员优惠券表创建参数 */
export interface CouponMemberCreateParams {
  couponID: string;
  memberID: string;
  orderID?: string;
  useStatus?: number;
  claimAt?: string;
  useAt?: string;
  expireAt?: string;
}

/** 会员优惠券表更新参数 */
export interface CouponMemberUpdateParams {
  id: string;
  couponID: string;
  memberID: string;
  orderID?: string;
  useStatus?: number;
  claimAt?: string;
  useAt?: string;
  expireAt?: string;
}
