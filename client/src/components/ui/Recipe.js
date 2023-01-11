/* eslint-disable no-unreachable */
import React from "react";
import { useNavigate } from "react-router-dom";
import "./item.css";

const Recipe = ({ingredients}) => {
  const navigate = useNavigate();
  const navigateToItem = (item) => {
    // 👇️ navigate to /
    console.log(item)
    navigate("/fullItems/" + item);
  };
  return (
    <div data-cy="article-container" className="container">
      Items Required: 
      {ingredients && ingredients.map((ingredient, index) => (
        <div data-cy="item" className="itemContainer" key={index}>
          <p data-cy="name" className="content" onClick={() => {navigateToItem(ingredient.item)}}>{ingredient.number} {ingredient.item}</p>
        </div>
      ))}
    </div>
  );
};

export default Recipe;