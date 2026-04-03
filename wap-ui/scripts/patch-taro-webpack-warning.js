const fs = require('fs');
const path = require('path');

const target = path.join(
  __dirname,
  '..',
  'node_modules',
  '@tarojs',
  'components',
  'dist',
  'components',
  'taro-video-core.js',
);

const marker = '/* webpackExports: ["default"] */';

if (!fs.existsSync(target)) {
  process.exit(0);
}

const source = fs.readFileSync(target, 'utf8');
if (!source.includes(marker)) {
  process.exit(0);
}

fs.writeFileSync(target, source.replace(marker, ''), 'utf8');
console.log('Patched taro-video-core.js to remove webpackExports warning.');
