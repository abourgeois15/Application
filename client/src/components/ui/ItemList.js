/* eslint-disable no-unreachable */
import React from "react";
import { useNavigate } from "react-router-dom";
import "./item.css";

const ItemList = ({items}) => {
  const navigate = useNavigate();
  const navigateToItem = (name) => {
    // 👇️ navigate to /
    console.log(name)
    navigate("./" + name);
  };
  return items.map((item, index) => (
    <div data-cy="article-container" className="container" key={index}>
      <div data-cy="item" className="itemContainer">
        <p data-cy="name" className="title" onClick={() => {navigateToItem(item.name)}}>{item.name}</p>
      </div>
    </div>
  ));
};

export default ItemList;