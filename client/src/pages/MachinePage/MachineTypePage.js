import React from "react";
import { useApi } from "../../hooks/useApi";
import { useParams } from 'react-router-dom'
import { services } from "../../services";
import { useNavigate } from "react-router-dom";
import MachineList from "../../components/ui/lists/MachineList";

export const MachineTypePage = () => {
  const navigate = useNavigate();

  const navigateToFullMachines = () => {
    // 👇️ navigate to /
    navigate("/fullMachines");
  };
  const params = useParams()
  const {state: names} = useApi(services.getMachineByType, [], params.machine_type);

  return (
    <div data-cy="machine-type-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>List Of {params.machine_type} Machines</h1>
      <button data-cy="A-goback-button" className="buttonG buttonGG" onClick={navigateToFullMachines}>
        Go Back To Full Machine List
      </button>
      <MachineList names={names}/>
    </div>
  );
};