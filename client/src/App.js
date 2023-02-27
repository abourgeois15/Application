import './App.css';
import { BrowserRouter as Router, Route, Routes} from "react-router-dom";
import { HomePage, ItemListPage, ItemPage, MachineListPage, MachinePage, MachineTypePage, DeleteItemPage, CreateItemPage, UpdateItemPage, DeleteMachinePage, CreateMachinePage, UpdateMachinePage, CraftPlannerPage } from './pages';


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />

        <Route path="/fullItems" element={<ItemListPage />} />
        <Route path="/fullItems/:item_name" element={<ItemPage />} />
        <Route path="/deleteItem/:item_name" element={<DeleteItemPage />} />
        <Route path="/createItem" element={<CreateItemPage />} />
        <Route path="/updateItem/:item_name" element={<UpdateItemPage />} />
        <Route path="/craftPlanner/:item_name" element={<CraftPlannerPage />} />

        <Route path="/fullMachines" element={<MachineListPage />} />
        <Route path="/fullMachines/name/:machine_name" element={<MachinePage />} />
        <Route path="/fullMachines/type/:machine_type" element={<MachineTypePage />} />
        <Route path="/deleteMachine/:machine_name" element={<DeleteMachinePage />} />
        <Route path="/createMachine" element={<CreateMachinePage />} />
        <Route path="/updateMachine/:machine_name" element={<UpdateMachinePage />} />
      </Routes>
    </Router>
  )
}

export default App;
