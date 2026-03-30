/** 余额流水表类型定义 */

/** 余额流水表项 */
export interface BalanceLogItem {
  id: string;
  memberID: string;
  bizType: number;
  bizID?: string;
  changeAmount: string;
  beforeBalance: string;
  afterBalance: string;
  remark?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** 余额流水表列表查询参数 */
export interface BalanceLogListParams {
  pageNum: number;
  pageSize: number;
  bizType?: number;
}

/** 余额流水表创建参数 */
export interface BalanceLogCreateParams {
  memberID: string;
  bizType: number;
  bizID?: string;
  changeAmount: string;
  beforeBalance: string;
  afterBalance: string;
  remark?: string;
}

/** 余额流水表更新参数 */
export interface BalanceLogUpdateParams {
  id: string;
  memberID: string;
  bizType: number;
  bizID?: string;
  changeAmount: string;
  beforeBalance: string;
  afterBalance: string;
  remark?: string;
}
