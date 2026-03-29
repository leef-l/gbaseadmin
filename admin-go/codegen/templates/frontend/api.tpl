import { requestClient } from '#/api/request';

import type {
  {{.ModelName}}Item,
  {{.ModelName}}ListParams,
  {{.ModelName}}CreateParams,
  {{.ModelName}}UpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/system/{{.ModuleName}}';

/** 获取{{.Comment}}列表 */
export function get{{.ModelName}}List(params: {{.ModelName}}ListParams) {
  return requestClient.get<{ list: {{.ModelName}}Item[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取{{.Comment}}详情 */
export function get{{.ModelName}}Detail(id: string) {
  return requestClient.get<{{.ModelName}}Item>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建{{.Comment}} */
export function create{{.ModelName}}(data: {{.ModelName}}CreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新{{.Comment}} */
export function update{{.ModelName}}(data: {{.ModelName}}UpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除{{.Comment}} */
export function delete{{.ModelName}}(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { params: { id } });
}
{{- if .HasParentID}}

/** 获取{{.Comment}}树形结构 */
export function get{{.ModelName}}Tree(params?: Record<string, any>) {
  return requestClient.get<{{.ModelName}}Item[]>(`${PREFIX}/tree`, { params });
}
{{- end}}
