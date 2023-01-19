/* eslint-disable no-unreachable */
import React from "react";
import "./item.css";
import IngredientForm from "./IngredientForm";
import SelectType from "./SelectType";

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
          <IngredientForm ingredient={machine.recipe[0]} handleChange={handleChangeRecipe} id="0"/>
          <IngredientForm ingredient={machine.recipe[1]} handleChange={handleChangeRecipe} id="1"/>
          <IngredientForm ingredient={machine.recipe[2]} handleChange={handleChangeRecipe} id="2"/>
        </div>
      </div>
      <div className="form-group">
        <label>Type Of Machine:</label>
        <SelectType name="type" type={machine.type} handleChange={handleChangeMachine}/>
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