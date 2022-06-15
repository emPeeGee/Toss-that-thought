import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { MantineProvider } from '@mantine/core';
import { Login } from 'features/authentication/Login';
import { Header } from 'components/layout/Header/Header';
import { GlobalStyles } from 'assets/styles/globalStyles';
import { ThoughtCreate } from 'features/thoughts/ThoughtCreate/ThoughtCreate';
import { ThoughtMetadata } from 'features/thoughts/ThougthMetadata/ThoughtMetadata';

function App() {
  return (
    <MantineProvider
      theme={{
        fontFamily: 'Open Sans, sans serif'
      }}>
      <BrowserRouter>
        <GlobalStyles />
        <Header />
        <Routes>
          <Route path="/" element={<ThoughtCreate />} />
          <Route path="/login" element={<Login />} />
          <Route path="/metadata/:metadataKey" element={<ThoughtMetadata />} />
        </Routes>
      </BrowserRouter>
    </MantineProvider>
  );
}

export default App;
