import { requestClient } from '#/api/request';

import type {
  ShopItem,
  ShopListParams,
  ShopCreateParams,
  ShopUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/shop';

/** 获取店铺表列表 */
export function getShopList(params: ShopListParams) {
  return requestClient.get<{ list: ShopItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取店铺表详情 */
export function getShopDetail(id: string) {
  return requestClient.get<ShopItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建店铺表 */
export function createShop(data: ShopCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新店铺表 */
export function updateShop(data: ShopUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除店铺表 */
export function deleteShop(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
