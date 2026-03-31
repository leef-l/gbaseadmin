import { requestClient } from '#/api/request';

import type {
  FileItem,
  FileListParams,
  FileCreateParams,
  FileUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/upload/file';

/** 获取æ–‡ä»¶è®°å½•列表 */
export function getFileList(params: FileListParams) {
  return requestClient.get<{ list: FileItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取æ–‡ä»¶è®°å½•详情 */
export function getFileDetail(id: string) {
  return requestClient.get<FileItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建æ–‡ä»¶è®°å½• */
export function createFile(data: FileCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新æ–‡ä»¶è®°å½• */
export function updateFile(data: FileUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除æ–‡ä»¶è®°å½• */
export function deleteFile(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
