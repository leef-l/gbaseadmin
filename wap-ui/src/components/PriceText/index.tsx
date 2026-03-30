import { Text } from '@tarojs/components';

interface Props {
  price: number;
  unit?: string;
  size?: 'sm' | 'md' | 'lg';
}

const sizeMap = { sm: 14, md: 16, lg: 24 };

export default function PriceText({ price, unit, size = 'md' }: Props) {
  return (
    <Text style={{ color: 'var(--accent)', fontWeight: '700', fontSize: `${sizeMap[size]}px` }}>
      ¥{(price / 100).toFixed(2)}
      {unit && <Text style={{ fontSize: '11px', color: 'var(--text-light)', fontWeight: '400' }}>{unit}</Text>}
    </Text>
  );
}
