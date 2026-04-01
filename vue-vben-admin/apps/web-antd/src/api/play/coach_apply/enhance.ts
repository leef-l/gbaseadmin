import { requestClient } from '#/api/request';

const PREFIX = '/play/coach_apply';

/** 审核陪玩师申请 */
export function auditCoachApply(data: {
  id: string;
  auditStatus: number;
  auditRemark?: string;
}) {
  return requestClient.put(`${PREFIX}/audit`, data);
}
