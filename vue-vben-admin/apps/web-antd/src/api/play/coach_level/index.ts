import { requestClient } from '#/api/request';

import type {
  CoachLevelItem,
  CoachLevelListParams,
  CoachLevelCreateParams,
  CoachLevelUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/coach_level';

/** 获取陪玩师等级表列表 */
export function getCoachLevelList(params: CoachLevelListParams) {
  return requestClient.get<{ list: CoachLevelItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取陪玩师等级表详情 */
export function getCoachLevelDetail(id: string) {
  return requestClient.get<CoachLevelItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建陪玩师等级表 */
export function createCoachLevel(data: CoachLevelCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新陪玩师等级表 */
export function updateCoachLevel(data: CoachLevelUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除陪玩师等级表 */
export function deleteCoachLevel(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
