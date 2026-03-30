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

export function completeStep(activityId: string, stepId: string) {
  return post('/api/playapi/activity/complete_step', { activityId, stepId });
}

export function claimReward(activityId: string, rewardId: string) {
  return post('/api/playapi/activity/claim_reward', { activityId, rewardId });
}

export function getMyActivities(params?: { page?: number; pageSize?: number }) {
  return get('/api/playapi/activity/my_list', params);
}
