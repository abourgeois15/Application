import './App.css';
import { BrowserRouter as Router, Route, Routes} from "react-router-dom";
import { HomePage, ItemListPage, ItemPage } from './pages';


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/fullItems" element={<ItemListPage />} />
        <Route path="/fullItems/:item_id" element={<ItemPage />} />
      </Routes>
    </Router>
  )

}

export default App;
