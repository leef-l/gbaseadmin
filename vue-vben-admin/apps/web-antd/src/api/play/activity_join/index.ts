import { requestClient } from '#/api/request';

import type {
  ActivityJoinItem,
  ActivityJoinListParams,
  ActivityJoinCreateParams,
  ActivityJoinUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/activity_join';

/** 获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨列表 */
export function getActivityJoinList(params: ActivityJoinListParams) {
  return requestClient.get<{ list: ActivityJoinItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨详情 */
export function getActivityJoinDetail(id: string) {
  return requestClient.get<ActivityJoinItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨ */
export function createActivityJoin(data: ActivityJoinCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨ */
export function updateActivityJoin(data: ActivityJoinUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除æ´»åŠ¨å‚ä¸Žè®°å½•è¡¨ */
export function deleteActivityJoin(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
