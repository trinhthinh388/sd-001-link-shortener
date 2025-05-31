import { createContext, type PropsWithChildren } from 'react';
import { customAlphabet } from 'nanoid';
import { ANONYMOUS_ID_LS_KEY } from '~/constants/anonymous';

export type AnonymousProviderContextType = {
  id: string;
};

export const AnonymousContext = createContext<AnonymousProviderContextType>({
  id: '',
});

const nanoid = customAlphabet('1234567890abcdef', 24);

export const AnonymousProvider = ({ children }: PropsWithChildren) => {
  const readOrGenerateAnonymousId = () => {
    try {
      let id = window.localStorage.getItem(ANONYMOUS_ID_LS_KEY);
      if (!id) {
        id = nanoid();
        window.localStorage.setItem(ANONYMOUS_ID_LS_KEY, id);
      }
      return id;
    } catch {
      const id = nanoid();
      window.localStorage.setItem(ANONYMOUS_ID_LS_KEY, id);
      return id;
    }
  };

  return (
    <AnonymousContext.Provider value={{ id: readOrGenerateAnonymousId() }}>
      {children}
    </AnonymousContext.Provider>
  );
};
