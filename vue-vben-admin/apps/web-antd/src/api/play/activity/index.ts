import { requestClient } from '#/api/request';

import type {
  ActivityItem,
  ActivityListParams,
  ActivityCreateParams,
  ActivityUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/activity';

/** 获取活动表列表 */
export function getActivityList(params: ActivityListParams) {
  return requestClient.get<{ list: ActivityItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取活动表详情 */
export function getActivityDetail(id: string) {
  return requestClient.get<ActivityItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建活动表 */
export function createActivity(data: ActivityCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新活动表 */
export function updateActivity(data: ActivityUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除活动表 */
export function deleteActivity(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
