import { requestClient } from '#/api/request';

import type {
  DirItem,
  DirListParams,
  DirCreateParams,
  DirUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/upload/dir';

/** 获取æ–‡ä»¶ç›®å½•列表 */
export function getDirList(params: DirListParams) {
  return requestClient.get<{ list: DirItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取æ–‡ä»¶ç›®å½•详情 */
export function getDirDetail(id: string) {
  return requestClient.get<DirItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建æ–‡ä»¶ç›®å½• */
export function createDir(data: DirCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新æ–‡ä»¶ç›®å½• */
export function updateDir(data: DirUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除æ–‡ä»¶ç›®å½• */
export function deleteDir(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}

/** 获取æ–‡ä»¶ç›®å½•树形结构 */
export async function getDirTree(params?: Record<string, any>) {
  const res = await requestClient.get<{ list: DirItem[] }>(`${PREFIX}/tree`, { params });
  return res?.list ?? [];
}
