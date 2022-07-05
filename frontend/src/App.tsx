import React, { useEffect, useMemo, useRef, useState } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { ColorScheme, ColorSchemeProvider, MantineProvider } from '@mantine/core';
import { GlobalStyles } from 'assets/styles/globalStyles';
import { api } from 'services/http';
import { NotificationsProvider, showNotification } from '@mantine/notifications';
import { useNetworkStatus } from 'hooks/use-network-status';
import { DateUnit } from 'utils/date';
import { AppShell, Footer, NotFound, ProtectedRoute, Offline, Header } from 'components';
import { Profile, SignIn, tokenIdentifier, UserModel, UserContext } from 'features/authentication';
import { ThoughtMetadata, ThoughtCreate, ThoughtBurn, ThoughtView } from 'features/thoughts';
import { RecentThoughts } from 'features/thoughts/RecentThoughts/RecentThoughts';

function App() {
  const [colorScheme, setColorScheme] = useState<ColorScheme>('light');
  const toggleColorScheme = (value?: ColorScheme) =>
    setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'));

  const { isOnline } = useNetworkStatus();
  const isFirstRun = useRef(true);
  const [user, setUser] = useState<UserModel | null>(null);
  const [token, setToken] = useState<string | null>(localStorage.getItem(tokenIdentifier));
  const value = useMemo(() => ({ user, setUser, token, setToken }), [user, token]);
  const [isLogged, setIsLogged] = useState(false);
  // TODO: is used to prepare user
  const [isAppReady, setIsAppReady] = useState(false);

  const logout = () => {
    setUser(null);
    setToken(null);
    // TODO: set token in localstorage
    localStorage.removeItem(tokenIdentifier);
    setIsAppReady(true);
  };

  useEffect(() => {
    if (user) {
      setIsLogged(!!user);
      setIsAppReady(true);
    }
  }, [user]);

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
    setIsAppReady(false);

    if ((token ?? '').length > 0) {
      api
        .get<UserModel>({
          url: `user`,
          token
        })
        .then((response) => setUser(response))
        .catch(() => logout());
    } else {
      setIsAppReady(true);
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
              {isAppReady && (
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
                      <Route element={<ProtectedRoute isAllowed={isLogged} />}>
                        <Route path="/profile" element={<Profile />} />
                        <Route path="/profile/recent" element={<RecentThoughts />} />
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
