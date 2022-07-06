import React from 'react';

export interface ThoughtViewContextModel {
  isLoading: boolean;
  setIsLoading: React.Dispatch<React.SetStateAction<boolean>>;
  isPassphrasePhasePassed: boolean;
  setIsPassphrasePhasePassed: React.Dispatch<React.SetStateAction<boolean>>;
  setThought: React.Dispatch<React.SetStateAction<string>>;
}

export const ThoughtViewContext = React.createContext<ThoughtViewContextModel | null>(null);
