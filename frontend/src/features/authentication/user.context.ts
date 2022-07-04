import React from 'react';
import { UserModel } from 'features/authentication/authentication.model';

export interface UserContextModel {
  user: UserModel | null;
  setUser: React.Dispatch<React.SetStateAction<UserModel | null>>;
}

export const UserContext = React.createContext<UserContextModel | null>(null);
