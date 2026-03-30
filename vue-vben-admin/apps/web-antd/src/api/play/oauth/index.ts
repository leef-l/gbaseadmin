import { requestClient } from '#/api/request';

import type {
  OauthItem,
  OauthListParams,
  OauthCreateParams,
  OauthUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/oauth';

/** 获取ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨列表 */
export function getOauthList(params: OauthListParams) {
  return requestClient.get<{ list: OauthItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨详情 */
export function getOauthDetail(id: string) {
  return requestClient.get<OauthItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨ */
export function createOauth(data: OauthCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨ */
export function updateOauth(data: OauthUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨ */
export function deleteOauth(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
