import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import Home from './components/Home';
import ProjectForm from './components/ProjectForm';
import './App.css';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/create-bff" element={<ProjectForm title="Create New BFF" />} />
        <Route path="/create-mf" element={<ProjectForm title="Create New MF" />} />
      </Routes>
    </Router>
  );
}

export default App;
