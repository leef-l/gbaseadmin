/** 优惠券模板表类型定义 */

/** 优惠券模板表项 */
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

/** 优惠券模板表列表查询参数 */
export interface CouponListParams {
  pageNum: number;
  pageSize: number;
  type?: number;
  isNewMember?: number;
  status?: number;
}

/** 优惠券模板表创建参数 */
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

/** 优惠券模板表更新参数 */
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
