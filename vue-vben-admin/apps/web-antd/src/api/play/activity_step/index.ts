import { requestClient } from '#/api/request';

import type {
  ActivityStepItem,
  ActivityStepListParams,
  ActivityStepCreateParams,
  ActivityStepUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/activity_step';

/** 获取æ´»åŠ¨æ­¥éª¤è¡¨列表 */
export function getActivityStepList(params: ActivityStepListParams) {
  return requestClient.get<{ list: ActivityStepItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取æ´»åŠ¨æ­¥éª¤è¡¨详情 */
export function getActivityStepDetail(id: string) {
  return requestClient.get<ActivityStepItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建æ´»åŠ¨æ­¥éª¤è¡¨ */
export function createActivityStep(data: ActivityStepCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新æ´»åŠ¨æ­¥éª¤è¡¨ */
export function updateActivityStep(data: ActivityStepUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除æ´»åŠ¨æ­¥éª¤è¡¨ */
export function deleteActivityStep(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
