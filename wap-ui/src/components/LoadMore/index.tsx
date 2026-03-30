import { View, Text } from '@tarojs/components';

interface Props {
  hasMore: boolean;
  loading?: boolean;
}

export default function LoadMore({ hasMore, loading }: Props) {
  return (
    <View style={{
      textAlign: 'center',
      padding: '16px 0',
      fontSize: '12px',
      color: 'var(--text-light)',
    }}>
      <Text>{loading ? '加载中...' : hasMore ? '上拉加载更多' : '没有更多了'}</Text>
    </View>
  );
}
