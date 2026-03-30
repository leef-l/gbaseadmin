import { requestClient } from '#/api/request';

import type {
  CategoryItem,
  CategoryListParams,
  CategoryCreateParams,
  CategoryUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/category';

/** 获取商品分类表列表 */
export function getCategoryList(params: CategoryListParams) {
  return requestClient.get<{ list: CategoryItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取商品分类表详情 */
export function getCategoryDetail(id: string) {
  return requestClient.get<CategoryItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建商品分类表 */
export function createCategory(data: CategoryCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新商品分类表 */
export function updateCategory(data: CategoryUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除商品分类表 */
export function deleteCategory(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}

/** 获取商品分类表树形结构 */
export async function getCategoryTree(params?: Record<string, any>) {
  const res = await requestClient.get<{ list: CategoryItem[] }>(`${PREFIX}/tree`, { params });
  return res?.list ?? [];
}
