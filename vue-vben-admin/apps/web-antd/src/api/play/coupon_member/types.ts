/** ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨类型定义 */

/** ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨项 */
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

/** ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨列表查询参数 */
export interface CouponMemberListParams {
  pageNum: number;
  pageSize: number;
  useStatus?: number;
}

/** ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨创建参数 */
export interface CouponMemberCreateParams {
  couponID: string;
  memberID: string;
  orderID?: string;
  useStatus?: number;
  claimAt?: string;
  useAt?: string;
  expireAt?: string;
}

/** ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨更新参数 */
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
