import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Thought } from './features/thoughts/Thought';
import { Login } from './features/authentication/Login';
import { Header } from './components/layout/Header/Header';
import { GlobalStyles } from './assets/styles/globalStyles';

function App() {
  return (
    <BrowserRouter>
      <GlobalStyles />
      <Header />
      <Routes>
        <Route path="/" element={<Thought />} />
        <Route path="/login" element={<Login />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
