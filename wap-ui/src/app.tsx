import { PropsWithChildren, useEffect } from 'react';
import { useAuthStore } from './store/auth';
import './app.scss';

function App({ children }: PropsWithChildren) {
  useEffect(() => {
    useAuthStore.getState().loadFromStorage();
  }, []);

  return <>{children}</>;
}

export default App;
