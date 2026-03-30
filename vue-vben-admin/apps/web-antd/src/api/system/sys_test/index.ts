import { requestClient } from '#/api/request';

import type {
  SysTestItem,
  SysTestListParams,
  SysTestCreateParams,
  SysTestUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/system/sys_test';

/** 获取æµ‹è¯•è¡¨列表 */
export function getSysTestList(params: SysTestListParams) {
  return requestClient.get<{ list: SysTestItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取æµ‹è¯•è¡¨详情 */
export function getSysTestDetail(id: string) {
  return requestClient.get<SysTestItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建æµ‹è¯•è¡¨ */
export function createSysTest(data: SysTestCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新æµ‹è¯•è¡¨ */
export function updateSysTest(data: SysTestUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除æµ‹è¯•è¡¨ */
export function deleteSysTest(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}

/** 获取æµ‹è¯•è¡¨树形结构 */
export async function getSysTestTree(params?: Record<string, any>) {
  const res = await requestClient.get<{ list: SysTestItem[] }>(`${PREFIX}/tree`, { params });
  return res?.list ?? [];
}
