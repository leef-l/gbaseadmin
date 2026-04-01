import type { UserConfigExport } from '@tarojs/cli';

export default {
  env: {
    NODE_ENV: '"development"',
  },
  defineConstants: {
    'process.env.TARO_APP_API': '""',
  },
  mini: {},
  h5: {},
} satisfies UserConfigExport;
