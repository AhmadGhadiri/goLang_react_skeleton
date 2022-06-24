import React from "react";
import { Routes, Route } from "react-router-dom";
// import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";
import Layout from './components/Layout/Layout';
import HomePage from './pages/HomePage';
import AuthPage from './pages/AuthPage';
import PostsPage from './pages/PostsPage';
import { useContext } from 'react';
import AuthContext from './store/auth-context';
import OtherPages from "./pages/OtherPages";
function App() {
  const authContext = useContext(AuthContext);

  return (
    <Layout>
      <Routes>
        <Route path="/" element={<HomePage />} />
        {!authContext.loggedIn && (
          <Route path="/auth" element={<AuthPage />} />
        )}
        <Route path="/posts" element={<PostsPage />} />
        <Route path="*" element={<OtherPages />}>
        </Route>
      </Routes>
    </Layout>
  );
}

export default App;