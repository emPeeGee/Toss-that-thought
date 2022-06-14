import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Thought } from './features/thoughts/Thought';
import { Login } from './features/authentication/Login';

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Thought />} />
          <Route path='/login' element={<Login />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
