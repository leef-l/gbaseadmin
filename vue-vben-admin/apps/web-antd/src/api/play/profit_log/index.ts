import { requestClient } from '#/api/request';

import type {
  ProfitLogItem,
  ProfitLogListParams,
  ProfitLogCreateParams,
  ProfitLogUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/profit_log';

/** 获取åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨列表 */
export function getProfitLogList(params: ProfitLogListParams) {
  return requestClient.get<{ list: ProfitLogItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨详情 */
export function getProfitLogDetail(id: string) {
  return requestClient.get<ProfitLogItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨ */
export function createProfitLog(data: ProfitLogCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨ */
export function updateProfitLog(data: ProfitLogUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨ */
export function deleteProfitLog(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
