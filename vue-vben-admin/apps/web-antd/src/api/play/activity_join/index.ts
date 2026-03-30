import { requestClient } from '#/api/request';

import type {
  ActivityJoinItem,
  ActivityJoinListParams,
  ActivityJoinCreateParams,
  ActivityJoinUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/activity_join';

/** 获取活动参与记录表列表 */
export function getActivityJoinList(params: ActivityJoinListParams) {
  return requestClient.get<{ list: ActivityJoinItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取活动参与记录表详情 */
export function getActivityJoinDetail(id: string) {
  return requestClient.get<ActivityJoinItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建活动参与记录表 */
export function createActivityJoin(data: ActivityJoinCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新活动参与记录表 */
export function updateActivityJoin(data: ActivityJoinUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除活动参与记录表 */
export function deleteActivityJoin(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
