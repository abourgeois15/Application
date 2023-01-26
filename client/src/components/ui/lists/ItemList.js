/* eslint-disable no-unreachable */
import React from "react";
import { useNavigate } from "react-router-dom";
import "../item.css";

const ItemList = ({names, value}) => {
  const navigate = useNavigate();
  const navigateToItem = (name) => {
    // 👇️ navigate to /
    navigate("./" + name);
  };
  return names.map((name, index) => (
    <div data-cy="item-container" className="container" key={index}>
      <div data-cy="item" className="itemContainer">
        {(name.toLowerCase().includes(value.toLowerCase())) && <p data-cy={name+"_cy"} className="title" onClick={() => {navigateToItem(name)}}>{name}</p>}
      </div>
    </div>
  ));
};

export default ItemList;