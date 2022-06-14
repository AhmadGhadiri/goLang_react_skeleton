import './App.css';
import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Dashboard from './components/Dashboard/Dashboard';
import Login from './components/Auth/Login';
import Preferences from './components/Preferences/Preferences';
import useToken from './useToken';
// import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  const { token, setToken } = useToken();
  console.log("token show1", token);

  if (!token) {
    return <Login setToken={setToken} />
  }


  return (
    <div className="wrapper">
      <h1>Application</h1>
      <BrowserRouter>
        <Routes>
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/preferences" element={<Preferences />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
