import type { UserConfigExport } from '@tarojs/cli';

export default {
  env: {
    NODE_ENV: '"development"',
  },
  defineConstants: {
    'process.env.TARO_APP_API': '""',
  },
  mini: {},
  h5: {
    publicPath: '/',
    devServer: {
      client: {
        overlay: {
          warnings: false,
        },
      },
    },
  },
  webpackChain(chain) {
    chain.merge({
      ignoreWarnings: [
        {
          module: /@tarojs[\\/]components[\\/]dist[\\/]components[\\/]taro-video-core\.js$/,
          message: /webpackExports/,
        },
      ],
    });
  },
} satisfies UserConfigExport;
