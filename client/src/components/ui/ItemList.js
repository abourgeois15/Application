/* eslint-disable no-unreachable */
import React from "react";
import { useNavigate } from "react-router-dom";
import "./item.css";

const ItemList = ({items}) => {
  const navigate = useNavigate();
  const navigateToItem = (id) => {
    // 👇️ navigate to /
    console.log(id)
    navigate("/fullItems/" + id);
  };
  return items.itemList && items.itemList.map((item, index) => (
    <div data-cy="article-container" className="container" key={index}>
      <div data-cy="item" className="itemContainer">
        <p data-cy="name" className="title" onClick={() => {navigateToItem(item.id)}}>{item.name}</p>
      </div>
    </div>
  ));
};

export default ItemList;