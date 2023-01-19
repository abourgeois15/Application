/* eslint-disable no-unreachable */
import React from "react";
import { useNavigate } from "react-router-dom";
import "./item.css";

const MachineList = ({names}) => {
  const navigate = useNavigate();
  const navigateToMachine = (name) => {
    // 👇️ navigate to /
    console.log(name)
    navigate("/fullMachines/name/" + name);
  };
  return names.map((name, index) => (
    <div data-cy="article-container" className="container" key={index}>
      <div data-cy="item" className="itemContainer">
        <p data-cy="name" className="title" onClick={() => {navigateToMachine(name)}}>{name}</p>
      </div>
    </div>
  ));
};

export default MachineList;