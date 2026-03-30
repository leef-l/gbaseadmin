/** 陪玩师申请表类型定义 */

/** 陪玩师申请表项 */
export interface CoachApplyItem {
  id: string;
  memberID: string;
  realName: string;
  idCard: string;
  idCardFrontImage: string;
  idCardBackImage: string;
  skillDesc?: string;
  auditStatus?: number;
  auditRemark?: string;
  auditAt?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** 陪玩师申请表列表查询参数 */
export interface CoachApplyListParams {
  pageNum: number;
  pageSize: number;
  auditStatus?: number;
}

/** 陪玩师申请表创建参数 */
export interface CoachApplyCreateParams {
  memberID: string;
  realName: string;
  idCard: string;
  idCardFrontImage: string;
  idCardBackImage: string;
  skillDesc?: string;
  auditStatus?: number;
  auditRemark?: string;
  auditAt?: string;
}

/** 陪玩师申请表更新参数 */
export interface CoachApplyUpdateParams {
  id: string;
  memberID: string;
  realName: string;
  idCard: string;
  idCardFrontImage: string;
  idCardBackImage: string;
  skillDesc?: string;
  auditStatus?: number;
  auditRemark?: string;
  auditAt?: string;
}
