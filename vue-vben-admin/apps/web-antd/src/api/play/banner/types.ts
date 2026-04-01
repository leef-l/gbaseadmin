/** 首页Banner轮播类型定义 */

/** 首页Banner轮播项 */
export interface BannerItem {
  id: string;
  title?: string;
  image?: string;
  linkType?: number;
  linkValue?: string;
  sort?: number;
  status?: number;
  startTime?: string;
  endTime?: string;
  remark?: string;
  createdAt?: string;
  updatedAt?: string;
}

/** 首页Banner轮播列表查询参数 */
export interface BannerListParams {
  pageNum: number;
  pageSize: number;
}

/** 首页Banner轮播创建参数 */
export interface BannerCreateParams {
  title?: string;
  image?: string;
  linkType?: number;
  linkValue?: string;
  sort?: number;
  status?: number;
  startTime?: string;
  endTime?: string;
  remark?: string;
}

/** 首页Banner轮播更新参数 */
export interface BannerUpdateParams {
  id: string;
  title?: string;
  image?: string;
  linkType?: number;
  linkValue?: string;
  sort?: number;
  status?: number;
  startTime?: string;
  endTime?: string;
  remark?: string;
}
