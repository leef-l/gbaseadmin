/** 第三方登录绑定表类型定义 */

/** 第三方登录绑定表项 */
export interface OauthItem {
  id: string;
  memberID: string;
  provider: number;
  openID: string;
  unionID?: string;
  nickname?: string;
  avatar?: string;
  accessToken?: string;
  refreshToken?: string;
  expireAt?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** 第三方登录绑定表列表查询参数 */
export interface OauthListParams {
  pageNum: number;
  pageSize: number;
  provider?: number;
}

/** 第三方登录绑定表创建参数 */
export interface OauthCreateParams {
  memberID: string;
  provider: number;
  openID: string;
  unionID?: string;
  nickname?: string;
  avatar?: string;
  accessToken?: string;
  refreshToken?: string;
  expireAt?: string;
}

/** 第三方登录绑定表更新参数 */
export interface OauthUpdateParams {
  id: string;
  memberID: string;
  provider: number;
  openID: string;
  unionID?: string;
  nickname?: string;
  avatar?: string;
  accessToken?: string;
  refreshToken?: string;
  expireAt?: string;
}
