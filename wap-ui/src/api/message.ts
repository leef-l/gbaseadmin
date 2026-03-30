import { get, post } from './request';

export function getMessageList(params?: { page?: number; pageSize?: number }) {
  return get('/api/playapi/member/messages', params);
}

export function markRead(messageId: string) {
  return post('/api/playapi/member/message/read', { messageId });
}

export function markAllRead() {
  return post('/api/playapi/member/message/read_all');
}
