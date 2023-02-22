import React, {useState} from 'react';
import './App.css';

import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import Layout from "./pages/Layout";
import Home from "./pages/Home";
import ViewExperiments from "./pages/ViewExperiments";
import ViewParticipants from "./pages/ViewParticipants";
import ViewRejectedSignUps from "./pages/ViewRejectedSignUps";
import ViewSignUps from "./pages/ViewSignUps";
import Error from "./pages/Error";

const App = () => {
  const[user, setUser] = useState({})
  return(
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout user={user}/>}>
          <Route index element={<Home user={user}/>} />
          <Route path="/experiments" element={<ViewExperiments user={user}/>} />
          <Route path="/participants" element={<ViewParticipants  user={user}/>} />
          <Route path="/rejected-sign-ups" element={<ViewRejectedSignUps  user={user}/>} />
          <Route path="/sign-ups" element={<ViewSignUps  user={user}/>} />
          <Route path="*" element={<Error />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
