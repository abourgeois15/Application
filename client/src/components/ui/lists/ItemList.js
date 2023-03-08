/* eslint-disable no-unreachable */
import React from "react";
import { useNavigate } from "react-router-dom";
import "../item.css";

const ItemList = ({items, value}) => {
  const navigate = useNavigate();
  const navigateToItem = (id) => {
    // 👇️ navigate to /
    navigate("./" + id);
  };
  return items.map((item, index) => (
    <div data-cy="item-container" className="container" key={index}>
      <div data-cy="item" className="itemContainer">
        {(item.name.toLowerCase().includes(value.toLowerCase())) && <p data-cy={item.name+"_cy"} className="title" onClick={() => {navigateToItem(item.id)}}>{item.name}</p>}
      </div>
    </div>
  ));
};

export default ItemList;