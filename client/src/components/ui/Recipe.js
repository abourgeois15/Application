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
    <div data-cy="recipe-container" className="container">
      Items Required: 
      {ingredients && ingredients.map((ingredient, index) => (
        <div data-cy="ingredient" className="itemContainer" key={index}>
          {(ingredient.number !== 0) && <p data-cy={"ingredient"+index} className="content" onClick={() => {navigateToItem(ingredient.item)}}>{ingredient.number} {ingredient.item}</p>}
        </div>
      ))}
    </div>
  );
};

export default Recipe;