import type { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    meta: {
      icon: 'game-icons:joystick',
      order: 20,
      title: '陪玩管理',
    },
    name: 'Play',
    path: '/play',
    children: [
      {
        path: '/play/member-level',
        name: 'PlayMemberLevel',
        component: () => import('#/views/play/member_level/index.vue'),
        meta: { title: '会员等级' },
      },
      {
        path: '/play/member',
        name: 'PlayMember',
        component: () => import('#/views/play/member/index.vue'),
        meta: { title: '会员列表' },
      },
      {
        path: '/play/coach-level',
        name: 'PlayCoachLevel',
        component: () => import('#/views/play/coach_level/index.vue'),
        meta: { title: '陪玩师等级' },
      },
      {
        path: '/play/coach-apply',
        name: 'PlayCoachApply',
        component: () => import('#/views/play/coach_apply/index.vue'),
        meta: { title: '陪玩师申请' },
      },
      {
        path: '/play/coach',
        name: 'PlayCoach',
        component: () => import('#/views/play/coach/index.vue'),
        meta: { title: '陪玩师列表' },
      },
      {
        path: '/play/shop',
        name: 'PlayShop',
        component: () => import('#/views/play/shop/index.vue'),
        meta: { title: '店铺管理' },
      },
      {
        path: '/play/category',
        name: 'PlayCategory',
        component: () => import('#/views/play/category/index.vue'),
        meta: { title: '商品分类' },
      },
      {
        path: '/play/goods',
        name: 'PlayGoods',
        component: () => import('#/views/play/goods/index.vue'),
        meta: { title: '商品列表' },
      },
      {
        path: '/play/order',
        name: 'PlayOrder',
        component: () => import('#/views/play/order/index.vue'),
        meta: { title: '订单列表' },
      },
      {
        path: '/play/payment',
        name: 'PlayPayment',
        component: () => import('#/views/play/payment/index.vue'),
        meta: { title: '支付记录' },
      },
      {
        path: '/play/recharge-plan',
        name: 'PlayRechargePlan',
        component: () => import('#/views/play/recharge_plan/index.vue'),
        meta: { title: '充值方案' },
      },
      {
        path: '/play/recharge-order',
        name: 'PlayRechargeOrder',
        component: () => import('#/views/play/recharge_order/index.vue'),
        meta: { title: '充值订单' },
      },
      {
        path: '/play/balance-log',
        name: 'PlayBalanceLog',
        component: () => import('#/views/play/balance_log/index.vue'),
        meta: { title: '余额流水' },
      },
      {
        path: '/play/activity',
        name: 'PlayActivity',
        component: () => import('#/views/play/activity/index.vue'),
        meta: { title: '活动列表' },
      },
      {
        path: '/play/activity-reward',
        name: 'PlayActivityReward',
        component: () => import('#/views/play/activity_reward/index.vue'),
        meta: { title: '活动奖励' },
      },
      {
        path: '/play/activity-step',
        name: 'PlayActivityStep',
        component: () => import('#/views/play/activity_step/index.vue'),
        meta: { title: '活动步骤' },
      },
      {
        path: '/play/activity-join',
        name: 'PlayActivityJoin',
        component: () => import('#/views/play/activity_join/index.vue'),
        meta: { title: '参与记录' },
      },
      {
        path: '/play/coupon',
        name: 'PlayCoupon',
        component: () => import('#/views/play/coupon/index.vue'),
        meta: { title: '优惠券管理' },
      },
      {
        path: '/play/coupon-member',
        name: 'PlayCouponMember',
        component: () => import('#/views/play/coupon_member/index.vue'),
        meta: { title: '领取记录' },
      },
      {
        path: '/play/review',
        name: 'PlayReview',
        component: () => import('#/views/play/review/index.vue'),
        meta: { title: '评价管理' },
      },
      {
        path: '/play/profit-log',
        name: 'PlayProfitLog',
        component: () => import('#/views/play/profit_log/index.vue'),
        meta: { title: '利润流水' },
      },
      {
        path: '/play/statistics',
        name: 'PlayStatistics',
        component: () => import('#/views/play/statistics/index.vue'),
        meta: { title: '数据统计' },
      },
    ],
  },
];

export default routes;