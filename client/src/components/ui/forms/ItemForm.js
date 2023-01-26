/* eslint-disable no-unreachable */
import React from "react";
import SelectType from "../selects/SelectType";
import RecipeForm from "./RecipeForm";

const ItemForm = ({item, handleSubmit, handleChangeItem, handleChangeRecipe}) => {
  return (
    <form data-cy="item-form" className="form-vertical" onSubmit={handleSubmit}>
      <div className="form-group">
        <label>Name:</label>
        <input data-cy="name" type="text" name="name" onChange={handleChangeItem} value={item.name} />
      </div>
      <div className="form-group">
        <label>Time:</label>
        <input data-cy="time" type="number" name="time" onChange={handleChangeItem} value={item.time} />
      </div>
      <div className="form-group">
        <RecipeForm recipe={item.recipe} handleChange={handleChangeRecipe}/>
      </div>
      <div className="form-group">
        <label>Result:</label>
        <input data-cy="result" type="number" name="result" onChange={handleChangeItem} value={item.result} />
      </div>
      <div className="form-group">
        <label>Type Of Machine:</label>
        <SelectType name="machineType" type={item.machineType} handleChange={handleChangeItem}/>
      </div>
      <input data-cy="submit" type="submit" />
    </form>
  )
};

export default ItemForm;