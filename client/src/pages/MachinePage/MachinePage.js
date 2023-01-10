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
  const params = useParams()
  const [state] = useApi(services.getMachineByName, [], params.machine_name);

  return (
    <div data-cy="detail-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Machine</h1>
        <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToMachines}>Go Back Machine Page</button>
        <Machine machine={state}/>
    </div>
  );
};