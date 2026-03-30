import { View, Text } from '@tarojs/components';

interface Props {
  text?: string;
}

export default function EmptyState({ text = '暂无数据' }: Props) {
  return (
    <View style={{
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center',
      justifyContent: 'center',
      padding: '60px 0',
    }}>
      <View style={{
        width: '80px',
        height: '80px',
        borderRadius: '50%',
        background: 'var(--border)',
        marginBottom: '16px',
      }} />
      <Text style={{ fontSize: '14px', color: 'var(--text-light)' }}>{text}</Text>
    </View>
  );
}
