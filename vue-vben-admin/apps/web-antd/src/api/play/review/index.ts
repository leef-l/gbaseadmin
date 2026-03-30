import { requestClient } from '#/api/request';

import type {
  ReviewItem,
  ReviewListParams,
  ReviewCreateParams,
  ReviewUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/review';

/** 获取评价表列表 */
export function getReviewList(params: ReviewListParams) {
  return requestClient.get<{ list: ReviewItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取评价表详情 */
export function getReviewDetail(id: string) {
  return requestClient.get<ReviewItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建评价表 */
export function createReview(data: ReviewCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新评价表 */
export function updateReview(data: ReviewUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除评价表 */
export function deleteReview(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
