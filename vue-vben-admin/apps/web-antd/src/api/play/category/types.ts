/** 商品分类表类型定义 */

/** 商品分类表项 */
export interface CategoryItem {
  id: string;
  parentID?: string;
  categoryTitle?: string;
  title: string;
  icon?: string;
  coverImage?: string;
  sort?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
  children?: CategoryItem[];
}

/** 商品分类表列表查询参数 */
export interface CategoryListParams {
  pageNum: number;
  pageSize: number;
  status?: number;
}

/** 商品分类表创建参数 */
export interface CategoryCreateParams {
  parentID?: string;
  title: string;
  icon?: string;
  coverImage?: string;
  sort?: number;
  status?: number;
}

/** 商品分类表更新参数 */
export interface CategoryUpdateParams {
  id: string;
  parentID?: string;
  title: string;
  icon?: string;
  coverImage?: string;
  sort?: number;
  status?: number;
}
