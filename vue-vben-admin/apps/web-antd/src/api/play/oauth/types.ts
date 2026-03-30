/** ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨类型定义 */

/** ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨项 */
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

/** ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨列表查询参数 */
export interface OauthListParams {
  pageNum: number;
  pageSize: number;
  provider?: number;
}

/** ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨创建参数 */
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

/** ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨更新参数 */
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
