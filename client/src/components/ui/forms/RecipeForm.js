import React from "react";
import IngredientForm from "./IngredientForm";

const RecipeForm = ({recipe, handleChange}) => {
  return (
    <div className="form-recipe">
      <label>Recipe:</label>
        {recipe.map((ingredient, index) => (
          <IngredientForm ingredient={ingredient} handleChange={handleChange} key={index} id={String(index)}/>
        ))}
    </div>
  )
};

export default RecipeForm;