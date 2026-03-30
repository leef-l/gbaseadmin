import { get } from './request';

export function getGoodsList(params?: any) {
  return get('/api/playapi/goods/list', params);
}

export function getGoodsDetail(goodsId: string) {
  return get('/api/playapi/goods/detail', { goodsId });
}

export function getCategoryList() {
  return get('/api/playapi/category/list');
}

export function search(params: { keyword: string; type?: string; page?: number; pageSize?: number }) {
  return get('/api/playapi/search', params);
}
