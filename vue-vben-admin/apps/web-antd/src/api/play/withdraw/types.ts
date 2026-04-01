/** 陪玩师提现记录类型定义 */

/** 陪玩师提现记录项 */
export interface WithdrawItem {
  id: string;
  coachID: string;
  coachRealName?: string;
  memberID: string;
  memberNickname?: string;
  amount?: number;
  status?: number;
  reason?: string;
  auditedAt?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** 陪玩师提现记录列表查询参数 */
export interface WithdrawListParams {
  pageNum: number;
  pageSize: number;
  orderBy?: string;
  orderDir?: string;
  startTime?: string;
  endTime?: string;
}

/** 陪玩师提现记录创建参数 */
export interface WithdrawCreateParams {
  coachID: string;
  memberID: string;
  amount?: number;
  status?: number;
  reason?: string;
  auditedAt?: string;
}

/** 陪玩师提现记录更新参数 */
export interface WithdrawUpdateParams {
  id: string;
  coachID: string;
  memberID: string;
  amount?: number;
  status?: number;
  reason?: string;
  auditedAt?: string;
}
