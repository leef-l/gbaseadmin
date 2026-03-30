/** 商品表类型定义 */

/** 商品表项 */
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

/** 商品表列表查询参数 */
export interface GoodsListParams {
  pageNum: number;
  pageSize: number;
  status?: number;
}

/** 商品表创建参数 */
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

/** 商品表更新参数 */
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
