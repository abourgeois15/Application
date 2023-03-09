/* eslint-disable no-unreachable */
import React from "react";
import { useNavigate } from "react-router-dom";
import "../item.css";

const MachineList = ({machines}) => {
  const navigate = useNavigate();
  const navigateToMachine = (id) => {
    // 👇️ navigate to /
    navigate("/fullMachines/id/" + id);
  };
  return machines.map((machine, index) => (
    <div data-cy="machine-container" className="container" key={index}>
      <div data-cy="machine" className="itemContainer">
        <p data-cy={machine.name+"_cy"} className="title" onClick={() => {navigateToMachine(machine.id)}}>{machine.name}</p>
      </div>
    </div>
  ));
};

export default MachineList;