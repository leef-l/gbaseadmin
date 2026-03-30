import { requestClient } from '#/api/request';

import type {
  MemberLevelItem,
  MemberLevelListParams,
  MemberLevelCreateParams,
  MemberLevelUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/member_level';

/** 获取会员等级表列表 */
export function getMemberLevelList(params: MemberLevelListParams) {
  return requestClient.get<{ list: MemberLevelItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取会员等级表详情 */
export function getMemberLevelDetail(id: string) {
  return requestClient.get<MemberLevelItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建会员等级表 */
export function createMemberLevel(data: MemberLevelCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新会员等级表 */
export function updateMemberLevel(data: MemberLevelUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除会员等级表 */
export function deleteMemberLevel(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
