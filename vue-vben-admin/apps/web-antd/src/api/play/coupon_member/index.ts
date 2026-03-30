import { requestClient } from '#/api/request';

import type {
  CouponMemberItem,
  CouponMemberListParams,
  CouponMemberCreateParams,
  CouponMemberUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/coupon_member';

/** 获取会员优惠券表列表 */
export function getCouponMemberList(params: CouponMemberListParams) {
  return requestClient.get<{ list: CouponMemberItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取会员优惠券表详情 */
export function getCouponMemberDetail(id: string) {
  return requestClient.get<CouponMemberItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建会员优惠券表 */
export function createCouponMember(data: CouponMemberCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新会员优惠券表 */
export function updateCouponMember(data: CouponMemberUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除会员优惠券表 */
export function deleteCouponMember(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
