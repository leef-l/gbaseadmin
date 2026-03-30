/** å……å€¼æ–¹æ¡ˆè¡¨类型定义 */

/** å……å€¼æ–¹æ¡ˆè¡¨项 */
export interface RechargePlanItem {
  id: string;
  title: string;
  amount: string;
  giftAmount?: string;
  coverImage?: string;
  sort?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** å……å€¼æ–¹æ¡ˆè¡¨列表查询参数 */
export interface RechargePlanListParams {
  pageNum: number;
  pageSize: number;
  status?: number;
}

/** å……å€¼æ–¹æ¡ˆè¡¨创建参数 */
export interface RechargePlanCreateParams {
  title: string;
  amount: string;
  giftAmount?: string;
  coverImage?: string;
  sort?: number;
  status?: number;
}

/** å……å€¼æ–¹æ¡ˆè¡¨更新参数 */
export interface RechargePlanUpdateParams {
  id: string;
  title: string;
  amount: string;
  giftAmount?: string;
  coverImage?: string;
  sort?: number;
  status?: number;
}
