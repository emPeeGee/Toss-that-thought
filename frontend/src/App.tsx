import React, { useState } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { ColorScheme, ColorSchemeProvider, MantineProvider } from '@mantine/core';
import { Login } from 'features/authentication/Login';
import { Header } from 'components/layout/Header/Header';
import { GlobalStyles } from 'assets/styles/globalStyles';
import { Footer } from 'components/layout/Footer/Footer';
import { AppShell } from 'components/layout/AppShell/AppShell';
import { ThoughtMetadata, ThoughtCreate, ThoughtBurn, ThoughtView } from 'features/thoughts';
import { NotFound } from 'components/layout/NotFound/NotFound';
import { NotificationsProvider } from '@mantine/notifications';

function App() {
  const [colorScheme, setColorScheme] = useState<ColorScheme>('light');
  const toggleColorScheme = (value?: ColorScheme) =>
    setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'));

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
            <AppShell>
              <Header />
              <Routes>
                <Route path="/" element={<ThoughtCreate />} />
                <Route path="/login" element={<Login />} />
                <Route path="/metadata/:metadataKey" element={<ThoughtMetadata />} />
                <Route path="/thought/:thoughtKey" element={<ThoughtView />} />
                <Route path="/thought/:metadataKey/burn" element={<ThoughtBurn />} />
                <Route path="*" element={<NotFound />} />
              </Routes>
              <Footer />
            </AppShell>
          </BrowserRouter>
        </NotificationsProvider>
      </MantineProvider>
    </ColorSchemeProvider>
  );
}

export default App;
