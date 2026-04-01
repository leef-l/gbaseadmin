import { requestClient } from '#/api/request';

import type {
  MessageItem,
  MessageListParams,
  MessageCreateParams,
  MessageUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/message';

/** 获取会员消息列表 */
export function getMessageList(params: MessageListParams) {
  return requestClient.get<{ list: MessageItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取会员消息详情 */
export function getMessageDetail(id: string) {
  return requestClient.get<MessageItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建会员消息 */
export function createMessage(data: MessageCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新会员消息 */
export function updateMessage(data: MessageUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除会员消息 */
export function deleteMessage(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}

/** 批量删除会员消息 */
export function batchDeleteMessage(ids: string[]) {
  return requestClient.delete(`${PREFIX}/batch-delete`, { data: { ids } });
}

/** 导出会员消息 */
export function exportMessage(params?: Record<string, any>) {
  return requestClient.get(`${PREFIX}/export`, {
    params,
    responseType: 'blob',
  });
}
