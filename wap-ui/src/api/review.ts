import { get, post } from './request';

export function getReviewList(params: { coachId: string; page?: number; pageSize?: number }) {
  return get('/api/playapi/review/list', params);
}

export function createReview(data: {
  orderId: string;
  score: number;
  content: string;
  images?: string;
  isAnonymous?: number;
}) {
  return post('/api/playapi/review/create', data);
}

export function replyReview(reviewId: string, reply: string) {
  return post('/api/playapi/review/reply', { reviewId, reply });
}
