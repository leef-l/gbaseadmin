/** æ–‡ä»¶ç›®å½•类型定义 */

/** æ–‡ä»¶ç›®å½•项 */
export interface DirItem {
  id: string;
  parentID?: string;
  dirName?: string;
  name: string;
  path: string;
  sort?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
  children?: DirItem[];
}

/** æ–‡ä»¶ç›®å½•列表查询参数 */
export interface DirListParams {
  pageNum: number;
  pageSize: number;
  status?: number;
}

/** æ–‡ä»¶ç›®å½•创建参数 */
export interface DirCreateParams {
  parentID?: string;
  name: string;
  path: string;
  sort?: number;
  status?: number;
}

/** æ–‡ä»¶ç›®å½•更新参数 */
export interface DirUpdateParams {
  id: string;
  parentID?: string;
  name: string;
  path: string;
  sort?: number;
  status?: number;
}
