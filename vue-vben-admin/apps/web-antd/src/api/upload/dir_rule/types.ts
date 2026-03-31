/** æ–‡ä»¶ç›®å½•è§„åˆ™类型定义 */

/** æ–‡ä»¶ç›®å½•è§„åˆ™项 */
export interface DirRuleItem {
  id: string;
  dirID: string;
  dirName?: string;
  category?: number;
  savePath?: string;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** æ–‡ä»¶ç›®å½•è§„åˆ™列表查询参数 */
export interface DirRuleListParams {
  pageNum: number;
  pageSize: number;
  category?: number;
  status?: number;
}

/** æ–‡ä»¶ç›®å½•è§„åˆ™创建参数 */
export interface DirRuleCreateParams {
  dirID: string;
  category?: number;
  savePath?: string;
  status?: number;
}

/** æ–‡ä»¶ç›®å½•è§„åˆ™更新参数 */
export interface DirRuleUpdateParams {
  id: string;
  dirID: string;
  category?: number;
  savePath?: string;
  status?: number;
}
