import { requestClient } from '#/api/request';

import type {
  RechargePlanItem,
  RechargePlanListParams,
  RechargePlanCreateParams,
  RechargePlanUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/recharge_plan';

/** 获取å……å€¼æ–¹æ¡ˆè¡¨列表 */
export function getRechargePlanList(params: RechargePlanListParams) {
  return requestClient.get<{ list: RechargePlanItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取å……å€¼æ–¹æ¡ˆè¡¨详情 */
export function getRechargePlanDetail(id: string) {
  return requestClient.get<RechargePlanItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建å……å€¼æ–¹æ¡ˆè¡¨ */
export function createRechargePlan(data: RechargePlanCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新å……å€¼æ–¹æ¡ˆè¡¨ */
export function updateRechargePlan(data: RechargePlanUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除å……å€¼æ–¹æ¡ˆè¡¨ */
export function deleteRechargePlan(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
