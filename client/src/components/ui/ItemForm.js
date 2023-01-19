/* eslint-disable no-unreachable */
import React from "react";
import "./item.css";

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
          <label>Ingredient:</label>
          <div className="form-ingredient">
            <label>Number:</label>
            <input type="number" name="number" id="0" onChange={handleChangeRecipe} value={item.recipe[0].number} />
            <label>Item:</label>
            <input type="text" name="item" id="0" onChange={handleChangeRecipe} value={item.recipe[0].item} />
          </div>
          <div className="form-ingredient">
            <label>Number:</label>
            <input type="number" name="number" id="1" onChange={handleChangeRecipe} value={item.recipe[1].number} />
            <label>Item:</label>
            <input type="text" name="item" id="1" onChange={handleChangeRecipe} value={item.recipe[1].item} />
          </div>
          <div className="form-ingredient">
            <label>Number:</label>
            <input type="number" name="number" id="2" onChange={handleChangeRecipe} value={item.recipe[2].number} />
            <label>Item:</label>
            <input type="text" name="item" id="2" onChange={handleChangeRecipe} value={item.recipe[2].item} />
          </div>
        </div>
      </div>
      <div className="form-group">
        <label>Result:</label>
        <input type="number" name="result" onChange={handleChangeItem} value={item.result} />
      </div>
      <div className="form-group">
        <label>Type Of Machine:</label>
        <select name="machineType" value={item.machineType} onChange={handleChangeItem}>
          {typeOptions.map((type, index) => (
            <option value={type} key={index}>{type}</option>
          ))}
        </select>
      </div>
      <input type="submit" />
    </form>
  )
};

export default ItemForm;