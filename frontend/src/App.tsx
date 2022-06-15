import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { MantineProvider } from '@mantine/core';
import { Login } from 'features/authentication/Login';
import { Header } from 'components/layout/Header/Header';
import { GlobalStyles } from 'assets/styles/globalStyles';
import { ThoughtCreate } from 'features/thoughts/ThoughtCreate/ThoughtCreate';
import { ThoughtMetadata } from 'features/thoughts/ThougthMetadata/ThoughtMetadata';
import { Footer } from 'components/layout/Footer/Footer';
import { AppShell } from 'components/layout/AppShell/AppShell';

function App() {
  return (
    <MantineProvider
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
          </Routes>
          <Footer />
        </AppShell>
      </BrowserRouter>
    </MantineProvider>
  );
}

export default App;
