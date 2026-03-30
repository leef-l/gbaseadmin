/** 店铺表类型定义 */

/** 店铺表项 */
export interface ShopItem {
  id: string;
  title: string;
  logoImage?: string;
  coverImage?: string;
  contactName?: string;
  contactPhone?: string;
  intro?: string;
  commissionRate?: number;
  coachNum?: number;
  sort?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** 店铺表列表查询参数 */
export interface ShopListParams {
  pageNum: number;
  pageSize: number;
  status?: number;
}

/** 店铺表创建参数 */
export interface ShopCreateParams {
  title: string;
  logoImage?: string;
  coverImage?: string;
  contactName?: string;
  contactPhone?: string;
  intro?: string;
  commissionRate?: number;
  coachNum?: number;
  sort?: number;
  status?: number;
}

/** 店铺表更新参数 */
export interface ShopUpdateParams {
  id: string;
  title: string;
  logoImage?: string;
  coverImage?: string;
  contactName?: string;
  contactPhone?: string;
  intro?: string;
  commissionRate?: number;
  coachNum?: number;
  sort?: number;
  status?: number;
}
