/* eslint-disable no-unreachable */
import React from "react";
import SelectType from "../selects/SelectType";
import RecipeForm from "./RecipeForm";

const ItemForm = ({item, handleSubmit, handleChangeItem, handleChangeRecipe}) => {
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
        <RecipeForm recipe={item.recipe} handleChange={handleChangeRecipe}/>
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