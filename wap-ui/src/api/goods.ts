import { get } from './request';

export function getGoodsList(params: any) {
  return get('/api/playapi/goods/list', params);
}

export function getGoodsDetail(id: string) {
  return get('/api/playapi/goods/detail', { id });
}

export function getCategoryList() {
  return get('/api/playapi/category/list');
}

export function search(params: any) {
  return get('/api/playapi/search', params);
}
