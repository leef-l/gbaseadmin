/** ä½™é¢æµæ°´è¡¨类型定义 */

/** ä½™é¢æµæ°´è¡¨项 */
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

/** ä½™é¢æµæ°´è¡¨列表查询参数 */
export interface BalanceLogListParams {
  pageNum: number;
  pageSize: number;
  bizType?: number;
}

/** ä½™é¢æµæ°´è¡¨创建参数 */
export interface BalanceLogCreateParams {
  memberID: string;
  bizType: number;
  bizID?: string;
  changeAmount: string;
  beforeBalance: string;
  afterBalance: string;
  remark?: string;
}

/** ä½™é¢æµæ°´è¡¨更新参数 */
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
