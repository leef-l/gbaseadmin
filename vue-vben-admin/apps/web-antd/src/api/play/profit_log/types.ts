/** 利润分成流水表类型定义 */

/** 利润分成流水表项 */
export interface ProfitLogItem {
  id: string;
  orderID: string;
  orderNo: string;
  payAmount?: string;
  coachID: string;
  shopID?: string;
  shopTitle?: string;
  platformRate?: number;
  platformAmount?: string;
  shopRate?: number;
  shopAmount?: string;
  coachAmount?: string;
  settleStatus?: number;
  settleAt?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** 利润分成流水表列表查询参数 */
export interface ProfitLogListParams {
  pageNum: number;
  pageSize: number;
  settleStatus?: number;
}

/** 利润分成流水表创建参数 */
export interface ProfitLogCreateParams {
  orderID: string;
  orderNo: string;
  payAmount?: string;
  coachID: string;
  shopID?: string;
  platformRate?: number;
  platformAmount?: string;
  shopRate?: number;
  shopAmount?: string;
  coachAmount?: string;
  settleStatus?: number;
  settleAt?: string;
}

/** 利润分成流水表更新参数 */
export interface ProfitLogUpdateParams {
  id: string;
  orderID: string;
  orderNo: string;
  payAmount?: string;
  coachID: string;
  shopID?: string;
  platformRate?: number;
  platformAmount?: string;
  shopRate?: number;
  shopAmount?: string;
  coachAmount?: string;
  settleStatus?: number;
  settleAt?: string;
}
