import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { MantineProvider } from '@mantine/core';
import { Thought } from './features/thoughts/Thought';
import { Login } from './features/authentication/Login';
import { Header } from './components/layout/Header/Header';
import { GlobalStyles } from './assets/styles/globalStyles';

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
          <Route path="/" element={<Thought />} />
          <Route path="/login" element={<Login />} />
        </Routes>
      </BrowserRouter>
    </MantineProvider>
  );
}

export default App;
