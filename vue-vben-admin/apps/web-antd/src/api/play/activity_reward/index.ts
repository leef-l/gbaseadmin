import { requestClient } from '#/api/request';

import type {
  ActivityRewardItem,
  ActivityRewardListParams,
  ActivityRewardCreateParams,
  ActivityRewardUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/activity_reward';

/** 获取活动奖励表列表 */
export function getActivityRewardList(params: ActivityRewardListParams) {
  return requestClient.get<{ list: ActivityRewardItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取活动奖励表详情 */
export function getActivityRewardDetail(id: string) {
  return requestClient.get<ActivityRewardItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建活动奖励表 */
export function createActivityReward(data: ActivityRewardCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新活动奖励表 */
export function updateActivityReward(data: ActivityRewardUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除活动奖励表 */
export function deleteActivityReward(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
