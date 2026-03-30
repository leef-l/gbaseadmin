import { requestClient } from '#/api/request';

import type {
  OauthItem,
  OauthListParams,
  OauthCreateParams,
  OauthUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/oauth';

/** 获取第三方登录绑定表列表 */
export function getOauthList(params: OauthListParams) {
  return requestClient.get<{ list: OauthItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取第三方登录绑定表详情 */
export function getOauthDetail(id: string) {
  return requestClient.get<OauthItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建第三方登录绑定表 */
export function createOauth(data: OauthCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新第三方登录绑定表 */
export function updateOauth(data: OauthUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除第三方登录绑定表 */
export function deleteOauth(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
