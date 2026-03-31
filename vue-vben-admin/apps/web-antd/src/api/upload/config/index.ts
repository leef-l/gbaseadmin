import { requestClient } from '#/api/request';

import type {
  ConfigItem,
  ConfigListParams,
  ConfigCreateParams,
  ConfigUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/upload/config';

/** 获取ä¸Šä¼ é…ç½®列表 */
export function getConfigList(params: ConfigListParams) {
  return requestClient.get<{ list: ConfigItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取ä¸Šä¼ é…ç½®详情 */
export function getConfigDetail(id: string) {
  return requestClient.get<ConfigItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建ä¸Šä¼ é…ç½® */
export function createConfig(data: ConfigCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新ä¸Šä¼ é…ç½® */
export function updateConfig(data: ConfigUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除ä¸Šä¼ é…ç½® */
export function deleteConfig(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
