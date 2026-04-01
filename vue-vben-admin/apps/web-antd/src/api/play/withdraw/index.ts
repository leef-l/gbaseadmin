import { requestClient } from '#/api/request';

import type {
  WithdrawItem,
  WithdrawListParams,
  WithdrawCreateParams,
  WithdrawUpdateParams,
} from './types';

/** API 前缀 */
const PREFIX = '/play/withdraw';

/** 获取陪玩师提现记录列表 */
export function getWithdrawList(params: WithdrawListParams) {
  return requestClient.get<{ list: WithdrawItem[]; total: number }>(
    `${PREFIX}/list`,
    { params },
  );
}

/** 获取陪玩师提现记录详情 */
export function getWithdrawDetail(id: string) {
  return requestClient.get<WithdrawItem>(`${PREFIX}/detail`, {
    params: { id },
  });
}

/** 创建陪玩师提现记录 */
export function createWithdraw(data: WithdrawCreateParams) {
  return requestClient.post(`${PREFIX}/create`, data);
}

/** 更新陪玩师提现记录 */
export function updateWithdraw(data: WithdrawUpdateParams) {
  return requestClient.put(`${PREFIX}/update`, data);
}

/** 删除陪玩师提现记录 */
export function deleteWithdraw(id: string) {
  return requestClient.delete(`${PREFIX}/delete`, { data: { id } });
}

/** 批量删除陪玩师提现记录 */
export function batchDeleteWithdraw(ids: string[]) {
  return requestClient.delete(`${PREFIX}/batch-delete`, { data: { ids } });
}

/** 导出陪玩师提现记录 */
export function exportWithdraw(params?: Record<string, any>) {
  return requestClient.get(`${PREFIX}/export`, {
    params,
    responseType: 'blob',
  });
}
