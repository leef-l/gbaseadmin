/** å•†å“è¡¨类型定义 */

/** å•†å“è¡¨项 */
export interface GoodsItem {
  id: string;
  categoryID: string;
  categoryTitle?: string;
  coachID: string;
  title: string;
  coverImage?: string;
  descContent?: string;
  price?: string;
  unit?: string;
  salesNum?: number;
  sort?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** å•†å“è¡¨列表查询参数 */
export interface GoodsListParams {
  pageNum: number;
  pageSize: number;
  status?: number;
}

/** å•†å“è¡¨创建参数 */
export interface GoodsCreateParams {
  categoryID: string;
  coachID: string;
  title: string;
  coverImage?: string;
  descContent?: string;
  price?: string;
  unit?: string;
  salesNum?: number;
  sort?: number;
  status?: number;
}

/** å•†å“è¡¨更新参数 */
export interface GoodsUpdateParams {
  id: string;
  categoryID: string;
  coachID: string;
  title: string;
  coverImage?: string;
  descContent?: string;
  price?: string;
  unit?: string;
  salesNum?: number;
  sort?: number;
  status?: number;
}
