import { requestClient } from '#/api/request';

import type {
  GoodsItem,
  GoodsListParams,
  GoodsCreateParams,
  GoodsUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/goods';

/** 获取å•†å“è¡¨列表 */
export function getGoodsList(params: GoodsListParams) {
  return requestClient.get<{ list: GoodsItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取å•†å“è¡¨详情 */
export function getGoodsDetail(id: string) {
  return requestClient.get<GoodsItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建å•†å“è¡¨ */
export function createGoods(data: GoodsCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新å•†å“è¡¨ */
export function updateGoods(data: GoodsUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除å•†å“è¡¨ */
export function deleteGoods(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
