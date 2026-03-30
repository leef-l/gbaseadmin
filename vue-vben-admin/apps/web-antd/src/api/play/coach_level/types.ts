/** 陪玩师等级表类型定义 */

/** 陪玩师等级表项 */
export interface CoachLevelItem {
  id: string;
  title: string;
  level?: number;
  icon?: string;
  minOrders?: number;
  minScore?: number;
  commissionRate?: number;
  sort?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** 陪玩师等级表列表查询参数 */
export interface CoachLevelListParams {
  pageNum: number;
  pageSize: number;
  level?: number;
  status?: number;
}

/** 陪玩师等级表创建参数 */
export interface CoachLevelCreateParams {
  title: string;
  level?: number;
  icon?: string;
  minOrders?: number;
  minScore?: number;
  commissionRate?: number;
  sort?: number;
  status?: number;
}

/** 陪玩师等级表更新参数 */
export interface CoachLevelUpdateParams {
  id: string;
  title: string;
  level?: number;
  icon?: string;
  minOrders?: number;
  minScore?: number;
  commissionRate?: number;
  sort?: number;
  status?: number;
}
