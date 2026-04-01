import { requestClient } from '#/api/request';

import type {
  BannerItem,
  BannerListParams,
  BannerCreateParams,
  BannerUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/banner';

/** 获取首页Banner轮播列表 */
export function getBannerList(params: BannerListParams) {
  return requestClient.get<{ list: BannerItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取首页Banner轮播详情 */
export function getBannerDetail(id: string) {
  return requestClient.get<BannerItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建首页Banner轮播 */
export function createBanner(data: BannerCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新首页Banner轮播 */
export function updateBanner(data: BannerUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除首页Banner轮播 */
export function deleteBanner(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}

/** 批量删除首页Banner轮播 */
export function batchDeleteBanner(ids: string[]) {
  return requestClient.delete(`${PREFIX}/batch-delete`, { data: { ids } });
}

/** 导出首页Banner轮播 */
export function exportBanner(params?: Record<string, any>) {
  return requestClient.get(`${PREFIX}/export`, {
    params,
    responseType: 'blob',
  });
}
