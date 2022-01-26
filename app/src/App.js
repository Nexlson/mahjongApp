import Calculator from "../src/views/Calculator"
import Docs from "../src/views/Docs"
import Landing from "../src/views/Landing"
import Logger from "../src/views/Logger"
import {BrowserRouter as Router, Routes, Route} from "react-router-dom"

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Landing/>} />
        <Route path="/docs" element={<Docs/>} />
        <Route path="/logger" element={<Logger/>} />
        <Route path="/calculator" element={<Calculator/>} />
      </Routes>
    </Router>
  );
}

export default App;
