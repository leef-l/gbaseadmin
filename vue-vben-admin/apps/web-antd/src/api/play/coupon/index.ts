import { requestClient } from '#/api/request';

import type {
  CouponItem,
  CouponListParams,
  CouponCreateParams,
  CouponUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/coupon';

/** 获取ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨列表 */
export function getCouponList(params: CouponListParams) {
  return requestClient.get<{ list: CouponItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨详情 */
export function getCouponDetail(id: string) {
  return requestClient.get<CouponItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨ */
export function createCoupon(data: CouponCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨ */
export function updateCoupon(data: CouponUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨ */
export function deleteCoupon(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
