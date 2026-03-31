import { get, post } from './request';

export function getActivityList(params?: { page?: number; pageSize?: number }) {
  return get('/api/playapi/activity/list', params);
}

export function getActivityDetail(id: string) {
  return get('/api/playapi/activity/detail', { id });
}

export function joinActivity(id: string) {
  return post('/api/playapi/activity/join', { id });
}

export function completeStep(data: { id: string; stepId: string }) {
  return post('/api/playapi/activity/complete_step', data);
}

export function claimReward(id: string, rewardId: string) {
  return post('/api/playapi/activity/claim_reward', { id, rewardId });
}

export function getMyActivities(params?: { page?: number; pageSize?: number }) {
  return get('/api/playapi/activity/my_list', params);
}
