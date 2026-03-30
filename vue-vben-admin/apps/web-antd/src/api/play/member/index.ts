import { requestClient } from '#/api/request';

import type {
  MemberItem,
  MemberListParams,
  MemberCreateParams,
  MemberUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/member';

/** 获取会员表列表 */
export function getMemberList(params: MemberListParams) {
  return requestClient.get<{ list: MemberItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取会员表详情 */
export function getMemberDetail(id: string) {
  return requestClient.get<MemberItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建会员表 */
export function createMember(data: MemberCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新会员表 */
export function updateMember(data: MemberUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除会员表 */
export function deleteMember(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
