import { get } from './request';

export interface BannerItem {
  bannerId: string;
  title: string;
  image: string;
  linkType: number;
  linkValue: string;
}

export function getBannerList() {
  return get<{ list: BannerItem[] }>('/api/playapi/banner/list');
}
