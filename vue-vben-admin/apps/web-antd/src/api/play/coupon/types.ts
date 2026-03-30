/** ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨类型定义 */

/** ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨项 */
export interface CouponItem {
  id: string;
  title: string;
  type?: number;
  isNewMember?: number;
  faceValue?: string;
  minAmount?: string;
  totalNum?: number;
  usedNum?: number;
  claimNum?: number;
  perLimit?: number;
  validStartAt: string;
  validEndAt: string;
  sort?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨列表查询参数 */
export interface CouponListParams {
  pageNum: number;
  pageSize: number;
  type?: number;
  isNewMember?: number;
  status?: number;
}

/** ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨创建参数 */
export interface CouponCreateParams {
  title: string;
  type?: number;
  isNewMember?: number;
  faceValue?: string;
  minAmount?: string;
  totalNum?: number;
  usedNum?: number;
  claimNum?: number;
  perLimit?: number;
  validStartAt: string;
  validEndAt: string;
  sort?: number;
  status?: number;
}

/** ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨更新参数 */
export interface CouponUpdateParams {
  id: string;
  title: string;
  type?: number;
  isNewMember?: number;
  faceValue?: string;
  minAmount?: string;
  totalNum?: number;
  usedNum?: number;
  claimNum?: number;
  perLimit?: number;
  validStartAt: string;
  validEndAt: string;
  sort?: number;
  status?: number;
}
