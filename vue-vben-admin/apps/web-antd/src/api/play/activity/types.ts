/** æ´»åŠ¨è¡¨类型定义 */

/** æ´»åŠ¨è¡¨项 */
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

/** æ´»åŠ¨è¡¨列表查询参数 */
export interface ActivityListParams {
  pageNum: number;
  pageSize: number;
  type?: number;
  conditionType?: number;
  isAutoReward?: number;
  status?: number;
}

/** æ´»åŠ¨è¡¨创建参数 */
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

/** æ´»åŠ¨è¡¨更新参数 */
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
