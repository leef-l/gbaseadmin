import { get, post } from './request';

export function getMessageList(params?: { type?: number; page?: number; pageSize?: number }) {
  return get('/api/playapi/member/messages', params);
}

export function markRead(id: string) {
  return post('/api/playapi/member/message/read', { id });
}

export function markAllRead() {
  return post('/api/playapi/member/message/read_all');
}
