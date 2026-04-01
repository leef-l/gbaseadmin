import { requestClient } from '#/api/request';

import type {
  ActivityStepLogItem,
  ActivityStepLogListParams,
  ActivityStepLogCreateParams,
  ActivityStepLogUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/activity_step_log';

/** 获取活动步骤提交记录列表 */
export function getActivityStepLogList(params: ActivityStepLogListParams) {
  return requestClient.get<{ list: ActivityStepLogItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取活动步骤提交记录详情 */
export function getActivityStepLogDetail(id: string) {
  return requestClient.get<ActivityStepLogItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建活动步骤提交记录 */
export function createActivityStepLog(data: ActivityStepLogCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新活动步骤提交记录 */
export function updateActivityStepLog(data: ActivityStepLogUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除活动步骤提交记录 */
export function deleteActivityStepLog(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
