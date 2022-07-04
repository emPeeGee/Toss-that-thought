import { useState } from 'react';

const tokenIdentifier = 'token';

export function useToken() {
  const getToken = (): string | null => {
    return localStorage.getItem(tokenIdentifier);
  };

  const [token, setToken] = useState(getToken());

  const saveToken = (newToken: string) => {
    console.log('save token');
    localStorage.setItem(tokenIdentifier, newToken);
    setToken(newToken);
  };

  console.log('returning new');

  return {
    setToken: saveToken,
    token
  };
}
