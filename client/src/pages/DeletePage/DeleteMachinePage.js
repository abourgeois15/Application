import React from "react";
import { useParams } from 'react-router-dom'
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import { useNavigate } from "react-router-dom";

export const DeleteMachinePage = () => {
  const navigate = useNavigate();
  const navigateToMachines = () => {
      // 👇️ navigate to /contacts
    navigate("/fullMachines");
    };
  const params = useParams()
  useApi(services.deleteMachine, [], params.machine_name);

  return (
    <div data-cy="delete-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Machine {params.machine_name} Deleted</h1>
        <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToMachines}>Go Back To Machines Page</button>
    </div>
  );
};