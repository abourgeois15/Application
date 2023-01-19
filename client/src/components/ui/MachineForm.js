/* eslint-disable no-unreachable */
import React from "react";
import "./item.css";

const MachineForm = ({machine, handleSubmit, handleChangeMachine, handleChangeRecipe}) => {
  const typeOptions = ["Assembling", "Chemical", "Furnace", "Mining"]
  return (
    <form className="form-vertical" onSubmit={handleSubmit}>
      <div className="form-group">
        <label>Name:</label>
        <input type="text" name="name" onChange={handleChangeMachine} value={machine.name} />
      </div>
      <div className="form-group">
        <label>Time:</label>
        <input type="number" name="time" onChange={handleChangeMachine} value={machine.time} />
      </div>
      <div className="form-group">
        <label>Recipe:</label>
        <div className="form-recipe">
          <label>Ingredient:</label>
          <div className="form-ingredient">
            <label>Number:</label>
            <input type="number" name="number" id="0" onChange={handleChangeRecipe} value={machine.recipe[0].number} />
            <label>Machine:</label>
            <input type="text" name="item" id="0" onChange={handleChangeRecipe} value={machine.recipe[0].item} />
          </div>
          <div className="form-ingredient">
            <label>Number:</label>
            <input type="number" name="number" id="1" onChange={handleChangeRecipe} value={machine.recipe[1].number} />
            <label>Machine:</label>
            <input type="text" name="item" id="1" onChange={handleChangeRecipe} value={machine.recipe[1].item} />
          </div>
          <div className="form-ingredient">
            <label>Number:</label>
            <input type="number" name="number" id="2" onChange={handleChangeRecipe} value={machine.recipe[2].number} />
            <label>Machine:</label>
            <input type="text" name="item" id="2" onChange={handleChangeRecipe} value={machine.recipe[2].item} />
          </div>
        </div>
      </div>
      <div className="form-group">
        <label>Type Of Machine:</label>
        <select name="type" value={machine.machineType} onChange={handleChangeMachine}>
          {typeOptions.map((type, index) => (
            <option value={type} key={index}>{type}</option>
          ))}
        </select>
      </div>
      <div className="form-group">
        <label>Speed:</label>
        <input type="number" name="speed" onChange={handleChangeMachine} value={machine.speed} />
      </div>
      <input type="submit" />
    </form>
  )
};

export default MachineForm;