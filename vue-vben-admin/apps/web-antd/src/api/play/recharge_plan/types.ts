/** 充值方案表类型定义 */

/** 充值方案表项 */
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

/** 充值方案表列表查询参数 */
export interface RechargePlanListParams {
  pageNum: number;
  pageSize: number;
  status?: number;
}

/** 充值方案表创建参数 */
export interface RechargePlanCreateParams {
  title: string;
  amount: string;
  giftAmount?: string;
  coverImage?: string;
  sort?: number;
  status?: number;
}

/** 充值方案表更新参数 */
export interface RechargePlanUpdateParams {
  id: string;
  title: string;
  amount: string;
  giftAmount?: string;
  coverImage?: string;
  sort?: number;
  status?: number;
}
