import { requestClient } from '#/api/request';

import type {
  BalanceLogItem,
  BalanceLogListParams,
  BalanceLogCreateParams,
  BalanceLogUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/balance_log';

/** 获取ä½™é¢æµæ°´è¡¨列表 */
export function getBalanceLogList(params: BalanceLogListParams) {
  return requestClient.get<{ list: BalanceLogItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取ä½™é¢æµæ°´è¡¨详情 */
export function getBalanceLogDetail(id: string) {
  return requestClient.get<BalanceLogItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建ä½™é¢æµæ°´è¡¨ */
export function createBalanceLog(data: BalanceLogCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新ä½™é¢æµæ°´è¡¨ */
export function updateBalanceLog(data: BalanceLogUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除ä½™é¢æµæ°´è¡¨ */
export function deleteBalanceLog(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}
