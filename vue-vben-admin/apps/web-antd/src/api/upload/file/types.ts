/** æ–‡ä»¶è®°å½•类型定义 */

/** æ–‡ä»¶è®°å½•项 */
export interface FileItem {
  id: string;
  dirID?: string;
  dirName?: string;
  name: string;
  url: string;
  ext?: string;
  size?: string;
  mime?: string;
  storage?: number;
  isImage?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** æ–‡ä»¶è®°å½•列表查询参数 */
export interface FileListParams {
  pageNum: number;
  pageSize: number;
  dirID?: string;
  name?: string;
  storage?: number;
  isImage?: number;
}

/** æ–‡ä»¶è®°å½•创建参数 */
export interface FileCreateParams {
  dirID?: string;
  name: string;
  url: string;
  ext?: string;
  size?: string;
  mime?: string;
  storage?: number;
  isImage?: number;
}

/** æ–‡ä»¶è®°å½•更新参数 */
export interface FileUpdateParams {
  id: string;
  dirID?: string;
  name: string;
  url: string;
  ext?: string;
  size?: string;
  mime?: string;
  storage?: number;
  isImage?: number;
}
