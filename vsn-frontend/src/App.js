import React, {useState} from 'react';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import './App.css';

import Home from "./pages/Home";
import ViewExperiments from "./pages/ViewExperiments";
import ViewParticipants from "./pages/ViewParticipants";
import ViewRejectedSignUps from "./pages/ViewRejectedSignUps";
import ViewSignUps from "./pages/ViewSignUps";
import DownloadGame from "./pages/DownloadGame";
import NavBar from './components/NavBar';

const App = () => {
  const[activeItem, setActiveItem] = useState();
  return(
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<NavBar activatedLink={activeItem}/>}>
          <Route index element={<Home/>}/>
          <Route path="/experiments" element={<ViewExperiments/>}/>
          <Route path="/participants" element={<ViewParticipants/>} />
          <Route path="/rejected-sign-ups" element={<ViewRejectedSignUps/>} />
          <Route path="/sign-ups" element={<ViewSignUps/>} />
          <Route path="/download-game" element={<DownloadGame/>}/>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
