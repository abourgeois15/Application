/* eslint-disable no-unreachable */
import React from "react";
import "./item.css";
import Recipe from "./Recipe";

const Machine = ({machine}) => {
  console.log(machine)
  return (
    <div data-cy="article-container" className="container">
      <div data-cy="item" className="itemContainer">
        <p data-cy="name" className="title">{machine.name}</p>
        <p data-cy="name" className="content">Machine type: {machine.machineType}</p>
        <p data-cy="name" className="content">Crafting speed: {machine.speed}</p>
        <p data-cy="name" className="content">Crafting time: {machine.time}</p>
        <Recipe ingredients={machine.recipe}/>
      </div>
    </div>
  );
};

export default Machine;