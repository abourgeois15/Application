/* eslint-disable no-unreachable */
import React from "react";
import { useNavigate } from "react-router-dom";
import "./item.css";
import Recipe from "./Recipe";

const Item = ({item}) => {
  console.log(item)
  const navigate = useNavigate();
  const navigateToMachineType = (type) => {
    // 👇️ navigate to /
    console.log(type)
    navigate("/fullMachines/type/" + type);
  };
  return (
    <div data-cy="item-container" className="container">
      <div data-cy="item" className="itemContainer">
        <p data-cy="name" className="title">{item.name}</p>
        <p data-cy="machine-type" className="content" onClick={() => {navigateToMachineType(item.machineType)}}>Machine type: {item.machineType}</p>
        {item.time !== 0 && <p data-cy="time" className="content">Crafting time: {item.time}</p>}
        {item.recipe && <Recipe ingredients={item.recipe}/>}
        <p data-cy="result" className="content">Result: {item.result}</p>
      </div>
    </div>
  );
};

export default Item;