/* eslint-disable no-unreachable */
import React from "react";
import { useNavigate } from "react-router-dom";
import "./item.css";

const MachineList = ({machines}) => {
  const navigate = useNavigate();
  const navigateToMachine = (name) => {
    // 👇️ navigate to /
    console.log(name)
    navigate("./" + name);
  };
  return machines.machineList && machines.machineList.map((machine, index) => (
    <div data-cy="article-container" className="container" key={index}>
      <div data-cy="item" className="itemContainer">
        <p data-cy="name" className="title" onClick={() => {navigateToMachine(machine.name)}}>{machine.name}</p>
      </div>
    </div>
  ));
};

export default MachineList;