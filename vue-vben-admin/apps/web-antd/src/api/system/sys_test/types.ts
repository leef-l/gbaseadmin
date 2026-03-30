/** æµ‹è¯•è¡¨类型定义 */

/** æµ‹è¯•è¡¨项 */
export interface SysTestItem {
  id: string;
  parentID?: string;
  title?: string;
  code?: string;
  type?: number;
  status?: number;
  sort?: number;
  remark?: string;
  createdAt?: string;
  updatedAt?: string;
  children?: SysTestItem[];
}

/** æµ‹è¯•è¡¨列表查询参数 */
export interface SysTestListParams {
  pageNum: number;
  pageSize: number;
  type?: number;
  status?: number;
}

/** æµ‹è¯•è¡¨创建参数 */
export interface SysTestCreateParams {
  parentID?: string;
  title?: string;
  code?: string;
  type?: number;
  status?: number;
  sort?: number;
  remark?: string;
}

/** æµ‹è¯•è¡¨更新参数 */
export interface SysTestUpdateParams {
  id: string;
  parentID?: string;
  title?: string;
  code?: string;
  type?: number;
  status?: number;
  sort?: number;
  remark?: string;
}
