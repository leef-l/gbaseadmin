/** 活动步骤表类型定义 */

/** 活动步骤表项 */
export interface ActivityStepItem {
  id: string;
  activityID: string;
  activityTitle?: string;
  stepNum?: number;
  title: string;
  stepType?: number; // 1=文字 2=链接 3=图片
  exampleText?: string; // 示例文字或链接URL
  descContent?: string;
  stepImage?: string;
  isRequired?: number; // 0=不需要填写 1=需要填写
  sort?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** 活动步骤表列表查询参数 */
export interface ActivityStepListParams {
  pageNum: number;
  pageSize: number;
  activityID?: string;
}

/** 活动步骤表创建参数 */
export interface ActivityStepCreateParams {
  activityID: string;
  stepNum?: number;
  title: string;
  stepType?: number;
  exampleText?: string;
  descContent?: string;
  stepImage?: string;
  isRequired?: number;
  sort?: number;
}

/** 活动步骤表更新参数 */
export interface ActivityStepUpdateParams {
  id: string;
  activityID: string;
  stepNum?: number;
  title: string;
  stepType?: number;
  exampleText?: string;
  descContent?: string;
  stepImage?: string;
  isRequired?: number;
  sort?: number;
}
