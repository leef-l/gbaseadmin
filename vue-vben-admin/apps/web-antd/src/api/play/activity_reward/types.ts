/** 活动奖励表类型定义 */

/** 活动奖励表项 */
export interface ActivityRewardItem {
  id: string;
  activityID: string;
  activityTitle?: string;
  rewardType?: number;
  rewardValue?: string;
  rewardName: string;
  sort?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** 活动奖励表列表查询参数 */
export interface ActivityRewardListParams {
  pageNum: number;
  pageSize: number;
  rewardType?: number;
}

/** 活动奖励表创建参数 */
export interface ActivityRewardCreateParams {
  activityID: string;
  rewardType?: number;
  rewardValue?: string;
  rewardName: string;
  sort?: number;
}

/** 活动奖励表更新参数 */
export interface ActivityRewardUpdateParams {
  id: string;
  activityID: string;
  rewardType?: number;
  rewardValue?: string;
  rewardName: string;
  sort?: number;
}
