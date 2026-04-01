/** 会员消息类型定义 */

/** 会员消息项 */
export interface MessageItem {
  id: string;
  memberID: string;
  memberNickname?: string;
  title?: string;
  content?: string;
  msgType?: number;
  bizID?: string;
  isRead?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** 会员消息列表查询参数 */
export interface MessageListParams {
  pageNum: number;
  pageSize: number;
  orderBy?: string;
  orderDir?: string;
  startTime?: string;
  endTime?: string;
  title?: string;
}

/** 会员消息创建参数 */
export interface MessageCreateParams {
  memberID: string;
  title?: string;
  content?: string;
  msgType?: number;
  bizID?: string;
  isRead?: number;
  status?: number;
}

/** 会员消息更新参数 */
export interface MessageUpdateParams {
  id: string;
  memberID: string;
  title?: string;
  content?: string;
  msgType?: number;
  bizID?: string;
  isRead?: number;
  status?: number;
}
