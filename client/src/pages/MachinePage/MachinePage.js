import React from "react";
import { useParams } from 'react-router-dom'
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import { useNavigate } from "react-router-dom";
import Machine from "../../components/ui/Machine";

export const MachinePage = () => {
  const navigate = useNavigate();
  const navigateToMachines = () => {
    // 👇️ navigate to /contacts
    navigate("/fullMachines");
  };

  const navigateToDelete = (name) => {
    // 👇️ navigate to /contacts
  navigate("/deleteMachine/" + name);
  };

  const navigateToUpdate = (name) => {
    // 👇️ navigate to /contacts
    navigate("/updateMachine/" + name);
  };
  const params = useParams()
  const {state: machine} = useApi(services.getMachineByName, [], params.machine_name);

  return (
    <div data-cy="machine-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Machine</h1>
        <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToMachines}>Go Back Machine Page</button>
        <button data-cy="A-delete-button" className="buttonA buttonAA" onClick={() => {navigateToDelete(params.machine_name)}}>Delete Machine</button>
        <button data-cy="A-update-button" className="buttonA buttonAA" onClick={() => {navigateToUpdate(params.machine_name)}}>Modify Machine</button>
        <Machine machine={machine}/>
    </div>
  );
};