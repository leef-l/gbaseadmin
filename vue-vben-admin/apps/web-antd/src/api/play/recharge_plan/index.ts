import { requestClient } from '#/api/request';

import type {
  RechargePlanItem,
  RechargePlanListParams,
  RechargePlanCreateParams,
  RechargePlanUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/recharge_plan';

/** 获取充值方案表列表 */
export function getRechargePlanList(params: RechargePlanListParams) {
  return requestClient.get<{ list: RechargePlanItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取充值方案表详情 */
export function getRechargePlanDetail(id: string) {
  return requestClient.get<RechargePlanItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建充值方案表 */
export function createRechargePlan(data: RechargePlanCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新充值方案表 */
export function updateRechargePlan(data: RechargePlanUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除充值方案表 */
export function deleteRechargePlan(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
