/** 活动步骤表类型定义 */

/** 活动步骤表项 */
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

/** 活动步骤表列表查询参数 */
export interface ActivityStepListParams {
  pageNum: number;
  pageSize: number;
}

/** 活动步骤表创建参数 */
export interface ActivityStepCreateParams {
  activityID: string;
  stepNum?: number;
  title: string;
  descContent?: string;
  stepImage?: string;
  sort?: number;
}

/** 活动步骤表更新参数 */
export interface ActivityStepUpdateParams {
  id: string;
  activityID: string;
  stepNum?: number;
  title: string;
  descContent?: string;
  stepImage?: string;
  sort?: number;
}
