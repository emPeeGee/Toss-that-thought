import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { MantineProvider } from '@mantine/core';
import { Login } from 'features/authentication/Login';
import { Header } from 'components/layout/Header/Header';
import { GlobalStyles } from 'assets/styles/globalStyles';
import { Footer } from 'components/layout/Footer/Footer';
import { AppShell } from 'components/layout/AppShell/AppShell';
import { ThoughtMetadata, ThoughtCreate, ThoughtBurn, ThoughtView } from 'features/thoughts';
import { NotFound } from 'components/layout/NotFound/NotFound';

function App() {
  return (
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
        }
      }}
      theme={{
        fontFamily: 'Open Sans, sans serif'
      }}>
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
    </MantineProvider>
  );
}

export default App;
