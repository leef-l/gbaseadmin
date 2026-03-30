import { requestClient } from '#/api/request';

import type {
  CoachItem,
  CoachListParams,
  CoachCreateParams,
  CoachUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/coach';

/** 获取陪玩师表列表 */
export function getCoachList(params: CoachListParams) {
  return requestClient.get<{ list: CoachItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取陪玩师表详情 */
export function getCoachDetail(id: string) {
  return requestClient.get<CoachItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建陪玩师表 */
export function createCoach(data: CoachCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新陪玩师表 */
export function updateCoach(data: CoachUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除陪玩师表 */
export function deleteCoach(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
