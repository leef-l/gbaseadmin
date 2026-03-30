/** 陪玩师表类型定义 */

/** 陪玩师表项 */
export interface CoachItem {
  id: string;
  memberID: string;
  coachLevelID?: string;
  coachLevelTitle?: string;
  shopID?: string;
  shopTitle?: string;
  realName: string;
  intro?: string;
  coverImage?: string;
  totalOrders?: number;
  totalScore?: number;
  scoreNum?: number;
  incomeTotal?: string;
  incomeBalance?: string;
  isOnline?: number;
  sort?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** 陪玩师表列表查询参数 */
export interface CoachListParams {
  pageNum: number;
  pageSize: number;
  isOnline?: number;
  status?: number;
}

/** 陪玩师表创建参数 */
export interface CoachCreateParams {
  memberID: string;
  coachLevelID?: string;
  shopID?: string;
  realName: string;
  intro?: string;
  coverImage?: string;
  totalOrders?: number;
  totalScore?: number;
  scoreNum?: number;
  incomeTotal?: string;
  incomeBalance?: string;
  isOnline?: number;
  sort?: number;
  status?: number;
}

/** 陪玩师表更新参数 */
export interface CoachUpdateParams {
  id: string;
  memberID: string;
  coachLevelID?: string;
  shopID?: string;
  realName: string;
  intro?: string;
  coverImage?: string;
  totalOrders?: number;
  totalScore?: number;
  scoreNum?: number;
  incomeTotal?: string;
  incomeBalance?: string;
  isOnline?: number;
  sort?: number;
  status?: number;
}
