/** 活动表类型定义 */

/** 活动表项 */
export interface ActivityItem {
  id: string;
  title: string;
  coverImage?: string;
  descContent?: string;
  type?: number;
  conditionType?: number;
  conditionValue?: string;
  isAutoReward?: number;
  startAt: string;
  endAt: string;
  maxNum?: number;
  joinNum?: number;
  sort?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** 活动表列表查询参数 */
export interface ActivityListParams {
  pageNum: number;
  pageSize: number;
  type?: number;
  conditionType?: number;
  isAutoReward?: number;
  status?: number;
}

/** 活动表创建参数 */
export interface ActivityCreateParams {
  title: string;
  coverImage?: string;
  descContent?: string;
  type?: number;
  conditionType?: number;
  conditionValue?: string;
  isAutoReward?: number;
  startAt: string;
  endAt: string;
  maxNum?: number;
  joinNum?: number;
  sort?: number;
  status?: number;
}

/** 活动表更新参数 */
export interface ActivityUpdateParams {
  id: string;
  title: string;
  coverImage?: string;
  descContent?: string;
  type?: number;
  conditionType?: number;
  conditionValue?: string;
  isAutoReward?: number;
  startAt: string;
  endAt: string;
  maxNum?: number;
  joinNum?: number;
  sort?: number;
  status?: number;
}
