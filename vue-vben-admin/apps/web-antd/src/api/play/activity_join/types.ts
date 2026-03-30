/** 活动参与记录表类型定义 */

/** 活动参与记录表项 */
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

/** 活动参与记录表列表查询参数 */
export interface ActivityJoinListParams {
  pageNum: number;
  pageSize: number;
  joinStatus?: number;
}

/** 活动参与记录表创建参数 */
export interface ActivityJoinCreateParams {
  activityID: string;
  memberID: string;
  joinStatus?: number;
  currentStep?: number;
  finishAt?: string;
  rewardAt?: string;
  remark?: string;
}

/** 活动参与记录表更新参数 */
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
