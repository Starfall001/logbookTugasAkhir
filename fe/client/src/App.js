import React from 'react'
import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import ViewUser from "./users/ViewUser";

function App() {
  return (
      <Router>
        <Routes>
          <Route  index element={<ViewUser/>} />
        </Routes>
      </Router>
  )
}

export default App;
