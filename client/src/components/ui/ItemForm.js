/* eslint-disable no-unreachable */
import React from "react";
import "./item.css";
import IngredientForm from "./IngredientForm";
import SelectType from "./SelectType";

const ItemForm = ({item, handleSubmit, handleChangeItem, handleChangeRecipe}) => {
  const typeOptions = ["Assembling", "Chemical", "Furnace", "Mining"]
  return (
    <form className="form-vertical" onSubmit={handleSubmit}>
      <div className="form-group">
        <label>Name:</label>
        <input type="text" name="name" onChange={handleChangeItem} value={item.name} />
      </div>
      <div className="form-group">
        <label>Time:</label>
        <input type="number" name="time" onChange={handleChangeItem} value={item.time} />
      </div>
      <div className="form-group">
        <label>Recipe:</label>
        <div className="form-recipe">
          <IngredientForm ingredient={item.recipe[0]} handleChange={handleChangeRecipe} id="0"/>
          <IngredientForm ingredient={item.recipe[1]} handleChange={handleChangeRecipe} id="1"/>
          <IngredientForm ingredient={item.recipe[2]} handleChange={handleChangeRecipe} id="2"/>
        </div>
      </div>
      <div className="form-group">
        <label>Result:</label>
        <input type="number" name="result" onChange={handleChangeItem} value={item.result} />
      </div>
      <div className="form-group">
        <label>Type Of Machine:</label>
        <SelectType name="machineType" type={item.machineType} handleChange={handleChangeItem}/>
      </div>
      <input type="submit" />
    </form>
  )
};

export default ItemForm;