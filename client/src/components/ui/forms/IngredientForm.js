/* eslint-disable no-unreachable */
import React from "react";
import SelectItem from "../selects/SelectItem";

const IngredientForm = ({ingredient, handleChange, deleteIngredient, id}) => {
  return (
    <div className="form-ingredient">
      <label>Number:</label>
      <input data-cy={"number"+id} type="number" name="number" id={id} value={ingredient.number} onChange={handleChange} />
      <label>Item:</label>
      <SelectItem value={ingredient.item} id={id} handleChange={handleChange} />
      <input type="button" onClick={() => {deleteIngredient(ingredient)}} value="Delete Ingredient"/>
    </div>
  )
};

export default IngredientForm;