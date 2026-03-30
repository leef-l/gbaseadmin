/** 评价表类型定义 */

/** 评价表项 */
export interface ReviewItem {
  id: string;
  orderID: string;
  memberID: string;
  coachID: string;
  score?: number;
  reviewContent?: string;
  reviewImage?: string;
  replyContent?: string;
  replyAt?: string;
  isAnonymous?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** 评价表列表查询参数 */
export interface ReviewListParams {
  pageNum: number;
  pageSize: number;
  isAnonymous?: number;
  status?: number;
}

/** 评价表创建参数 */
export interface ReviewCreateParams {
  orderID: string;
  memberID: string;
  coachID: string;
  score?: number;
  reviewContent?: string;
  reviewImage?: string;
  replyContent?: string;
  replyAt?: string;
  isAnonymous?: number;
  status?: number;
}

/** 评价表更新参数 */
export interface ReviewUpdateParams {
  id: string;
  orderID: string;
  memberID: string;
  coachID: string;
  score?: number;
  reviewContent?: string;
  reviewImage?: string;
  replyContent?: string;
  replyAt?: string;
  isAnonymous?: number;
  status?: number;
}
