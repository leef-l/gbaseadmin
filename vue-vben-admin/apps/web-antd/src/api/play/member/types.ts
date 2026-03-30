/** 会员表类型定义 */

/** 会员表项 */
export interface MemberItem {
  id: string;
  phone: string;
  nickname?: string;
  avatar?: string;
  gender?: number;
  memberLevelID?: string;
  memberLevelTitle?: string;
  exp?: number;
  balance?: string;
  isCoach?: number;
  status?: number;
  lastLoginAt?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** 会员表列表查询参数 */
export interface MemberListParams {
  pageNum: number;
  pageSize: number;
  gender?: number;
  isCoach?: number;
  status?: number;
}

/** 会员表创建参数 */
export interface MemberCreateParams {
  phone: string;
  password: string;
  nickname?: string;
  avatar?: string;
  gender?: number;
  memberLevelID?: string;
  exp?: number;
  balance?: string;
  isCoach?: number;
  status?: number;
  lastLoginAt?: string;
}

/** 会员表更新参数 */
export interface MemberUpdateParams {
  id: string;
  phone: string;
  password: string;
  nickname?: string;
  avatar?: string;
  gender?: number;
  memberLevelID?: string;
  exp?: number;
  balance?: string;
  isCoach?: number;
  status?: number;
  lastLoginAt?: string;
}
