/** æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨类型定义 */

/** æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨项 */
export interface ActivityJoinItem {
  id: string;
  activityID: string;
  activityTitle?: string;
  memberID: string;
  joinStatus?: number;
  currentStep?: number;
  finishAt?: string;
  rewardAt?: string;
  remark?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨列表查询参数 */
export interface ActivityJoinListParams {
  pageNum: number;
  pageSize: number;
  joinStatus?: number;
}

/** æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨创建参数 */
export interface ActivityJoinCreateParams {
  activityID: string;
  memberID: string;
  joinStatus?: number;
  currentStep?: number;
  finishAt?: string;
  rewardAt?: string;
  remark?: string;
}

/** æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨更新参数 */
export interface ActivityJoinUpdateParams {
  id: string;
  activityID: string;
  memberID: string;
  joinStatus?: number;
  currentStep?: number;
  finishAt?: string;
  rewardAt?: string;
  remark?: string;
}
