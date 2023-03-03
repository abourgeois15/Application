import React from "react";
import IngredientForm from "./IngredientForm";

const RecipeForm = ({recipe, handleChange, addIngredient, deleteIngredient}) => {
  return (
    <div className="form-recipe">
      <label>Recipe:</label>
        {recipe.map((ingredient, index) => (
          (ingredient.number !== -1) && <IngredientForm ingredient={ingredient} handleChange={handleChange} key={index} deleteIngredient={deleteIngredient} id={String(index)}/>
        ))}
      <input type="button" onClick={addIngredient} value="Add Ingredient"/>
    </div>
  )
};

export default RecipeForm;