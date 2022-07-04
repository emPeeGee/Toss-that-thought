import React, { useEffect, useMemo, useRef, useState } from 'react';
import { BrowserRouter, Route, Routes, Outlet, Navigate } from 'react-router-dom';
import { ColorScheme, ColorSchemeProvider, MantineProvider, Title } from '@mantine/core';
import { Header } from 'components/layout/Header/Header';
import { GlobalStyles } from 'assets/styles/globalStyles';
// TODO, create index for component
import { Footer } from 'components/layout/Footer/Footer';
import { AppShell } from 'components/layout/AppShell/AppShell';
import { ThoughtMetadata, ThoughtCreate, ThoughtBurn, ThoughtView } from 'features/thoughts';
import { NotFound } from 'components/layout/NotFound/NotFound';
import { NotificationsProvider, showNotification } from '@mantine/notifications';
import { useNetworkStatus } from 'hooks/use-network-status';
import { Offline } from 'components/Offline/Offline';
import { DateUnit } from 'utils/date';
// TODO: index
import { SignIn } from 'features/authentication/SignIn/SignIn';
import { Profile } from 'features/authentication/Profile/Profile';
import { UserContext } from 'features/authentication/user.context';
import { api } from 'services/http';
import { UserModel } from 'features/authentication/authentication.model';
import { tokenIdentifier } from './features/authentication/constants';

interface ProtectedRouteProps {
  isAllowed: boolean;
  redirectPath?: string;
  children?: any;
}

// TODO: To be relocated
function ProtectedRoute({ isAllowed, redirectPath = '/', children }: ProtectedRouteProps) {
  if (!isAllowed) {
    return <Navigate to={redirectPath} replace />;
  }

  return children || <Outlet />;
}

ProtectedRoute.defaultProps = {
  redirectPath: '/',
  children: undefined
};

function App() {
  const [colorScheme, setColorScheme] = useState<ColorScheme>('light');
  const toggleColorScheme = (value?: ColorScheme) =>
    setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'));

  const { isOnline } = useNetworkStatus();
  const isFirstRun = useRef(true);
  const [user, setUser] = useState<UserModel | null>(null);
  const [token, setToken] = useState<string | null>(localStorage.getItem(tokenIdentifier));
  const value = useMemo(() => ({ user, setUser, token, setToken }), [user, token]);

  // TODO: is used to prepare user
  const [isReady, setIsReady] = useState(false);

  useEffect(() => {
    if (isFirstRun.current) {
      isFirstRun.current = false;
      return;
    }

    showNotification({
      title: isOnline ? 'You are online' : 'Oops. No internet connection.',
      message: isOnline
        ? 'Connection restored.'
        : 'Make sure wifi or cellular data is turned on and then try again.',
      color: isOnline ? 'green' : 'red',
      autoClose: DateUnit.second * 5
    });
  }, [isOnline]);

  useEffect(() => {
    console.log('Token', token);
    // TODO: set token in localstorage

    if ((token ?? '').length > 0) {
      api
        .get<UserModel>({
          url: `user`,
          token
        })
        .then((response) => {
          setUser(response);
        })
        .catch((err) => {
          console.log(err);
          setUser(null);
          setToken(null);
          localStorage.removeItem(tokenIdentifier);
        })
        .finally(() => {
          setIsReady(true);
        });
    } else {
      setIsReady(true);
    }
  }, [token]);

  return (
    <ColorSchemeProvider colorScheme={colorScheme} toggleColorScheme={toggleColorScheme}>
      <MantineProvider
        withGlobalStyles
        withNormalizeCSS
        defaultProps={{
          Container: {
            sizes: {
              xs: 540,
              sm: 720,
              md: 960,
              lg: 1140,
              xl: 1320
            }
          },
          Button: { tabIndex: 0 },
          Anchor: { tabIndex: 0 }
        }}
        styles={{
          Button: (theme) => ({
            root: {
              '&:focus': {
                outline: `2px solid ${theme.colors.orange[5]} !important`
              }
            }
          }),
          ActionIcon: (theme) => ({
            root: {
              '&:focus': {
                outline: `2px solid ${theme.colors.orange[5]} !important`
              }
            }
          })
        }}
        theme={{
          colorScheme,
          fontFamily: 'Open Sans, sans serif'
        }}>
        <NotificationsProvider>
          <BrowserRouter>
            <GlobalStyles />
            <UserContext.Provider value={value}>
              {isReady && (
                <AppShell>
                  <Header />
                  {!isOnline ? (
                    <Offline />
                  ) : (
                    <Routes>
                      <Route path="/" element={<ThoughtCreate />} />
                      <Route path="/sign-in" element={<SignIn />} />
                      <Route path="/metadata/:metadataKey" element={<ThoughtMetadata />} />
                      <Route path="/thought/:thoughtKey" element={<ThoughtView />} />
                      <Route path="/thought/:metadataKey/burn" element={<ThoughtBurn />} />
                      <Route element={<ProtectedRoute isAllowed={!!user} />}>
                        <Route path="/profile" element={<Profile />} />
                        <Route path="/profile/recent" element={<Title>Recent</Title>} />
                      </Route>
                      <Route path="*" element={<NotFound />} />
                    </Routes>
                  )}
                  <Footer />
                </AppShell>
              )}
            </UserContext.Provider>
          </BrowserRouter>
        </NotificationsProvider>
      </MantineProvider>
    </ColorSchemeProvider>
  );
}

export default App;
