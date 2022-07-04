import React from 'react';
import { UserModel } from 'features/authentication/authentication.model';

export interface UserContextModel {
  token: string | null;
  setToken: React.Dispatch<React.SetStateAction<string | null>>;
  user: UserModel | null;
  setUser: React.Dispatch<React.SetStateAction<UserModel | null>>;
}

export const UserContext = React.createContext<UserContextModel | null>(null);
