/** æ´»åŠ¨å¥–åŠ±è¡¨类型定义 */

/** æ´»åŠ¨å¥–åŠ±è¡¨项 */
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

/** æ´»åŠ¨å¥–åŠ±è¡¨列表查询参数 */
export interface ActivityRewardListParams {
  pageNum: number;
  pageSize: number;
  rewardType?: number;
}

/** æ´»åŠ¨å¥–åŠ±è¡¨创建参数 */
export interface ActivityRewardCreateParams {
  activityID: string;
  rewardType?: number;
  rewardValue?: string;
  rewardName: string;
  sort?: number;
}

/** æ´»åŠ¨å¥–åŠ±è¡¨更新参数 */
export interface ActivityRewardUpdateParams {
  id: string;
  activityID: string;
  rewardType?: number;
  rewardValue?: string;
  rewardName: string;
  sort?: number;
}
