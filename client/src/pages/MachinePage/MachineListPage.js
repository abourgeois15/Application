import React from "react";
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import { useNavigate } from "react-router-dom";
import MachineList from "../../components/ui/MachineList";

export const MachineListPage = () => {
  const navigate = useNavigate();

  const navigateToHome = () => {
    // 👇️ navigate to /
    navigate("/");
  };

  const navigateToCreate = () => {
    // 👇️ navigate to /
    navigate("/createMachine");
  };

  const {state: machines} = useApi(services.getMachines, []);

  return (
    <div data-cy="detail-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>List of Machines</h1>
      <button data-cy="gohome-button" className="buttonG buttonGG" onClick={navigateToHome}>Go Back To Home</button>
      <button data-cy="A-create-button" className="buttonA buttonAA" onClick={navigateToCreate}>Create New Machine</button>

      <MachineList machines={machines}/>
    </div>
  );
};