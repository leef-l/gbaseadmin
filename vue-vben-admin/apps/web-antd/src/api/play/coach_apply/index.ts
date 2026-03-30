import { requestClient } from '#/api/request';

import type {
  CoachApplyItem,
  CoachApplyListParams,
  CoachApplyCreateParams,
  CoachApplyUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/coach_apply';

/** 获取陪玩师申请表列表 */
export function getCoachApplyList(params: CoachApplyListParams) {
  return requestClient.get<{ list: CoachApplyItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取陪玩师申请表详情 */
export function getCoachApplyDetail(id: string) {
  return requestClient.get<CoachApplyItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建陪玩师申请表 */
export function createCoachApply(data: CoachApplyCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新陪玩师申请表 */
export function updateCoachApply(data: CoachApplyUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除陪玩师申请表 */
export function deleteCoachApply(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}

/** 审核陪玩师申请 */
export function auditCoachApply(data: { id: string; auditStatus: number; auditRemark?: string }) {
  return requestClient.put(`${PREFIX}/audit`, data);
}
