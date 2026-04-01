/** 活动步骤提交记录类型定义 */

/** 活动步骤提交记录项 */
export interface ActivityStepLogItem {
  id: string;
  activityID?: string;
  activityTitle?: string;
  stepID?: string;
  joinID?: string;
  memberID?: string;
  stepType?: number;
  submitText?: string;
  submitImage?: string;
  auditStatus?: number;
  auditRemark?: string;
  auditBy?: string;
  auditAt?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** 活动步骤提交记录列表查询参数 */
export interface ActivityStepLogListParams {
  pageNum: number;
  pageSize: number;
  stepType?: number;
  auditStatus?: number;
}

/** 活动步骤提交记录创建参数 */
export interface ActivityStepLogCreateParams {
  activityID?: string;
  stepID?: string;
  joinID?: string;
  memberID?: string;
  stepType?: number;
  submitText?: string;
  submitImage?: string;
  auditStatus?: number;
  auditRemark?: string;
  auditBy?: string;
  auditAt?: string;
}

/** 活动步骤提交记录更新参数 */
export interface ActivityStepLogUpdateParams {
  id: string;
  activityID?: string;
  stepID?: string;
  joinID?: string;
  memberID?: string;
  stepType?: number;
  submitText?: string;
  submitImage?: string;
  auditStatus?: number;
  auditRemark?: string;
  auditBy?: string;
  auditAt?: string;
}
