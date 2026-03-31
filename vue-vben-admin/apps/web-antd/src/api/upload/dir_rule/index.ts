import { requestClient } from '#/api/request';

import type {
  DirRuleItem,
  DirRuleListParams,
  DirRuleCreateParams,
  DirRuleUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/upload/dir_rule';

/** 获取æ–‡ä»¶ç›®å½•è§„åˆ™列表 */
export function getDirRuleList(params: DirRuleListParams) {
  return requestClient.get<{ list: DirRuleItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取æ–‡ä»¶ç›®å½•è§„åˆ™详情 */
export function getDirRuleDetail(id: string) {
  return requestClient.get<DirRuleItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建æ–‡ä»¶ç›®å½•è§„åˆ™ */
export function createDirRule(data: DirRuleCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新æ–‡ä»¶ç›®å½•è§„åˆ™ */
export function updateDirRule(data: DirRuleUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除æ–‡ä»¶ç›®å½•è§„åˆ™ */
export function deleteDirRule(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
