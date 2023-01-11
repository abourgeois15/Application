import './App.css';
import { BrowserRouter as Router, Route, Routes} from "react-router-dom";
import { HomePage, ItemListPage, ItemPage, MachineListPage, MachinePage, MachineTypePage } from './pages';


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />

        <Route path="/fullItems" element={<ItemListPage />} />
        <Route path="/fullItems/:item_name" element={<ItemPage />} />

        <Route path="/fullMachines" element={<MachineListPage />} />
        <Route path="/fullMachines/name/:machine_name" element={<MachinePage />} />
        <Route path="/fullMachines/type/:machine_type" element={<MachineTypePage />} />
      </Routes>
    </Router>
  )

}

export default App;
