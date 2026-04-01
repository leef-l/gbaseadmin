import { useState } from 'react';
import { WebView } from '@tarojs/components';
import Taro, { useLoad } from '@tarojs/taro';

export default function WebviewPage() {
  const [url, setUrl] = useState('');

  useLoad(() => {
    const params = Taro.getCurrentInstance().router?.params;
    if (params?.url) {
      setUrl(decodeURIComponent(params.url));
    }
  });

  if (!url) return null;

  return <WebView src={url} />;
}
