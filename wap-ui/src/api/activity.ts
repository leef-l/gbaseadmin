import { get, post } from './request';

export function getActivityList(params: any) {
  return get('/api/playapi/activity/list', params);
}

export function getActivityDetail(id: string) {
  return get('/api/playapi/activity/detail', { id });
}

export function joinActivity(id: string) {
  return post('/api/playapi/activity/join', { id });
}

export function completeStep(data: any) {
  return post('/api/playapi/activity/complete_step', data);
}

export function claimReward(id: string) {
  return post('/api/playapi/activity/claim_reward', { id });
}
