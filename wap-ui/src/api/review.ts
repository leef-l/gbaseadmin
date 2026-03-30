import { get, post } from './request';

export function getReviewList(params: any) {
  return get('/api/playapi/review/list', params);
}

export function createReview(data: any) {
  return post('/api/playapi/review/create', data);
}
