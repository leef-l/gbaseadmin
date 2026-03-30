import { requestClient } from '#/api/request';

import type {
  ActivityStepItem,
  ActivityStepListParams,
  ActivityStepCreateParams,
  ActivityStepUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/activity_step';

/** 获取活动步骤表列表 */
export function getActivityStepList(params: ActivityStepListParams) {
  return requestClient.get<{ list: ActivityStepItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取活动步骤表详情 */
export function getActivityStepDetail(id: string) {
  return requestClient.get<ActivityStepItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建活动步骤表 */
export function createActivityStep(data: ActivityStepCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新活动步骤表 */
export function updateActivityStep(data: ActivityStepUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除活动步骤表 */
export function deleteActivityStep(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
