/** åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨类型定义 */

/** åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨项 */
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

/** åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨列表查询参数 */
export interface ProfitLogListParams {
  pageNum: number;
  pageSize: number;
  settleStatus?: number;
}

/** åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨创建参数 */
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

/** åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨更新参数 */
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
