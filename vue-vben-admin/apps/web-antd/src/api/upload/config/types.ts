/** ä¸Šä¼ é…ç½®类型定义 */

/** ä¸Šä¼ é…ç½®项 */
export interface ConfigItem {
  id: string;
  name: string;
  storage?: number;
  isDefault?: number;
  localPath?: string;
  ossEndpoint?: string;
  ossBucket?: string;
  ossAccessKey?: string;
  ossSecretKey?: string;
  cosRegion?: string;
  cosBucket?: string;
  cosSecretID?: string;
  cosSecretKey?: string;
  maxSize?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

/** ä¸Šä¼ é…ç½®列表查询参数 */
export interface ConfigListParams {
  pageNum: number;
  pageSize: number;
  storage?: number;
  isDefault?: number;
  status?: number;
}

/** ä¸Šä¼ é…ç½®创建参数 */
export interface ConfigCreateParams {
  name: string;
  storage?: number;
  isDefault?: number;
  localPath?: string;
  ossEndpoint?: string;
  ossBucket?: string;
  ossAccessKey?: string;
  ossSecretKey?: string;
  cosRegion?: string;
  cosBucket?: string;
  cosSecretID?: string;
  cosSecretKey?: string;
  maxSize?: number;
  status?: number;
}

/** ä¸Šä¼ é…ç½®更新参数 */
export interface ConfigUpdateParams {
  id: string;
  name: string;
  storage?: number;
  isDefault?: number;
  localPath?: string;
  ossEndpoint?: string;
  ossBucket?: string;
  ossAccessKey?: string;
  ossSecretKey?: string;
  cosRegion?: string;
  cosBucket?: string;
  cosSecretID?: string;
  cosSecretKey?: string;
  maxSize?: number;
  status?: number;
}
