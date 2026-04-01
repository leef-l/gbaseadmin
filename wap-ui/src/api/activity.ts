import { get, post } from './request';

export function getActivityList(params?: { page?: number; pageSize?: number }) {
  return get('/api/playapi/activity/list', params);
}

export function getActivityDetail(activityId: string) {
  return get('/api/playapi/activity/detail', { activityId });
}

export function joinActivity(activityId: string) {
  return post('/api/playapi/activity/join', { activityId });
}

export function completeStep(data: { activityId: string; stepId: string; imageUrl?: string; submitText?: string }) {
  return post('/api/playapi/activity/complete_step', data);
}

export function quitActivity(activityId: string) {
  return post('/api/playapi/activity/quit', { activityId });
}

export function claimReward(activityId: string) {
  return post('/api/playapi/activity/claim_reward', { activityId });
}

export function getMyActivities(params?: { page?: number; pageSize?: number }) {
  return get('/api/playapi/activity/my_list', params);
}
