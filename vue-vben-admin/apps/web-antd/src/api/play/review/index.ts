import { requestClient } from '#/api/request';

import type {
  ReviewItem,
  ReviewListParams,
  ReviewCreateParams,
  ReviewUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/review';

/** 获取è¯„ä»·è¡¨列表 */
export function getReviewList(params: ReviewListParams) {
  return requestClient.get<{ list: ReviewItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取è¯„ä»·è¡¨详情 */
export function getReviewDetail(id: string) {
  return requestClient.get<ReviewItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建è¯„ä»·è¡¨ */
export function createReview(data: ReviewCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新è¯„ä»·è¡¨ */
export function updateReview(data: ReviewUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除è¯„ä»·è¡¨ */
export function deleteReview(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
