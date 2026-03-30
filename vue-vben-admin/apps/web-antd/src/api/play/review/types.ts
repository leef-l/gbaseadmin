/** è¯„ä»·è¡¨类型定义 */

/** è¯„ä»·è¡¨项 */
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

/** è¯„ä»·è¡¨列表查询参数 */
export interface ReviewListParams {
  pageNum: number;
  pageSize: number;
  isAnonymous?: number;
  status?: number;
}

/** è¯„ä»·è¡¨创建参数 */
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

/** è¯„ä»·è¡¨更新参数 */
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
