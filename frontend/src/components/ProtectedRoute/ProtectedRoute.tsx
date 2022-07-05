import React, { ReactElement } from 'react';
import { Navigate, Outlet } from 'react-router-dom';

interface ProtectedRouteProps {
  isAllowed: boolean;
  redirectPath?: string;
  children?: ReactElement;
}

export function ProtectedRoute({ isAllowed, redirectPath = '/', children }: ProtectedRouteProps) {
  if (!isAllowed) {
    return <Navigate to={redirectPath} replace />;
  }

  return children || <Outlet />;
}

ProtectedRoute.defaultProps = {
  redirectPath: '/',
  children: undefined
};
