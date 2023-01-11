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
    <div data-cy="article-container" className="container">
      <div data-cy="item" className="itemContainer">
        <p data-cy="name" className="title">{item.name}</p>
        <p data-cy="name" className="content" onClick={() => {navigateToMachineType(item.machineType)}}>Machine type: {item.machineType}</p>
        {item.time !== 0 && <p data-cy="name" className="content">Crafting time: {item.time}</p>}
        {item.recipe && <Recipe ingredients={item.recipe}/>}
      </div>
    </div>
  );
};

export default Item;