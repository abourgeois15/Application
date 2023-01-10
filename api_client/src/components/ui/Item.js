/* eslint-disable no-unreachable */
import React from "react";
import "./item.css";
import Recipe from "./Recipe";

const Item = ({item}) => {
  console.log(item)
  return (
    <div data-cy="article-container" className="container">
      <div data-cy="item" className="itemContainer">
        <p data-cy="name" className="title">{item.name}</p>
        {item.time !== 0 && <p data-cy="name" className="content">Crafting time: {item.time}</p>}
        <p data-cy="name" className="content">Machine type: {item.machineType}</p>
        {item.recipe && <Recipe ingredients={item.recipe}/>}
      </div>
    </div>
  );
};

export default Item;