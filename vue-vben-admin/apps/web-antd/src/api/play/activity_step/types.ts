/** æ´»åŠ¨æ­¥éª¤è¡¨类型定义 */

/** æ´»åŠ¨æ­¥éª¤è¡¨项 */
export interface ActivityStepItem {
  id: string;
  activityID: string;
  activityTitle?: string;
  stepNum?: number;
  title: string;
  descContent?: string;
  stepImage?: string;
  sort?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** æ´»åŠ¨æ­¥éª¤è¡¨列表查询参数 */
export interface ActivityStepListParams {
  pageNum: number;
  pageSize: number;
}

/** æ´»åŠ¨æ­¥éª¤è¡¨创建参数 */
export interface ActivityStepCreateParams {
  activityID: string;
  stepNum?: number;
  title: string;
  descContent?: string;
  stepImage?: string;
  sort?: number;
}

/** æ´»åŠ¨æ­¥éª¤è¡¨更新参数 */
export interface ActivityStepUpdateParams {
  id: string;
  activityID: string;
  stepNum?: number;
  title: string;
  descContent?: string;
  stepImage?: string;
  sort?: number;
}
