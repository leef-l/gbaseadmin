import { PropsWithChildren } from 'react';
import { useAuthStore } from './store/auth';
import './app.scss';

useAuthStore.getState().loadFromStorage();

function App({ children }: PropsWithChildren) {
  return <>{children}</>;
}

export default App;
