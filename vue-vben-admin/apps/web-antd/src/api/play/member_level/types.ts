/** 会员等级表类型定义 */

/** 会员等级表项 */
export interface MemberLevelItem {
  id: string;
  title: string;
  level?: number;
  icon?: string;
  minExp?: number;
  discount?: number;
  sort?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** 会员等级表列表查询参数 */
export interface MemberLevelListParams {
  pageNum: number;
  pageSize: number;
  level?: number;
  status?: number;
}

/** 会员等级表创建参数 */
export interface MemberLevelCreateParams {
  title: string;
  level?: number;
  icon?: string;
  minExp?: number;
  discount?: number;
  sort?: number;
  status?: number;
}

/** 会员等级表更新参数 */
export interface MemberLevelUpdateParams {
  id: string;
  title: string;
  level?: number;
  icon?: string;
  minExp?: number;
  discount?: number;
  sort?: number;
  status?: number;
}
