import { requestClient } from '#/api/request';

import type {
  CouponMemberItem,
  CouponMemberListParams,
  CouponMemberCreateParams,
  CouponMemberUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/coupon_member';

/** 获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨列表 */
export function getCouponMemberList(params: CouponMemberListParams) {
  return requestClient.get<{ list: CouponMemberItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨详情 */
export function getCouponMemberDetail(id: string) {
  return requestClient.get<CouponMemberItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨ */
export function createCouponMember(data: CouponMemberCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨ */
export function updateCouponMember(data: CouponMemberUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨ */
export function deleteCouponMember(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
