/* eslint-disable no-unreachable */
import React from "react";
import SelectType from "../selects/SelectType";
import RecipeForm from "./RecipeForm";

const MachineForm = ({machine, handleSubmit, handleChangeMachine, handleChangeRecipe}) => {
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
        <RecipeForm recipe={machine.recipe} handleChange={handleChangeRecipe}/>
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