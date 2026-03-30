import { Text } from '@tarojs/components';

interface Props {
  text: string;
  color: string;
}

export default function StatusTag({ text, color }: Props) {
  return (
    <Text style={{
      fontSize: '11px',
      padding: '2px 8px',
      borderRadius: '8px',
      background: `${color}18`,
      color,
    }}>
      {text}
    </Text>
  );
}
