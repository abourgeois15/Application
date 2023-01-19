/* eslint-disable no-unreachable */
import React from "react";
import "./item.css";
import SelectItem from "./SelectItem";

const IngredientForm = ({ingredient, handleChange, id}) => {
  return (
    <div className="form-ingredient">
      <label>Number:</label>
      <input type="number" name="number" id={id} value={ingredient.number} onChange={handleChange} />
      <label>Item:</label>
      <SelectItem value={ingredient.item} id={id} handleChange={handleChange} />
    </div>
  )
};

export default IngredientForm;