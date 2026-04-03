import { create } from 'zustand';

interface CartState {
  goodsId: string;
  goodsName: string;
  coachId: string;
  coachName: string;
  price: number;
  quantity: number;
  couponId: string;
  couponAmount: number;
  remark: string;
  setOrder: (data: Partial<CartState>) => void;
  startOrder: (data: Pick<CartState, 'goodsId' | 'goodsName' | 'coachId' | 'coachName' | 'price'> & Partial<Pick<CartState, 'quantity'>>) => void;
  reset: () => void;
}

const initialState = {
  goodsId: '',
  goodsName: '',
  coachId: '',
  coachName: '',
  price: 0,
  quantity: 1,
  couponId: '',
  couponAmount: 0,
  remark: '',
};

export const useCartStore = create<CartState>((set) => ({
  ...initialState,
  setOrder: (data) => set((state) => ({ ...state, ...data })),
  startOrder: (data) => set({
    ...initialState,
    ...data,
    quantity: data.quantity && data.quantity > 0 ? data.quantity : 1,
  }),
  reset: () => set(initialState),
}));
