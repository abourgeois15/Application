/* eslint-disable no-unreachable */
import React from "react";
import SelectType from "../selects/SelectType";
import RecipeForm from "./RecipeForm";

const MachineForm = ({machine, handleSubmit, handleChangeMachine, handleChangeRecipe, addIngredient, deleteIngredient}) => {
  return (
    <form data-cy="machine-form" className="form-vertical" onSubmit={handleSubmit}>
      <div className="form-group">
        <label>Name:</label>
        <input data-cy="name" type="text" name="name" onChange={handleChangeMachine} value={machine.name} />
      </div>
      <div className="form-group">
        <label>Time:</label>
        <input data-cy="time" type="number" name="time" onChange={handleChangeMachine} value={machine.time} />
      </div>
      <div className="form-group">
        <RecipeForm recipe={machine.recipe} handleChange={handleChangeRecipe} addIngredient={addIngredient} deleteIngredient={deleteIngredient}/>
      </div>
      <div className="form-group">
        <label>Type Of Machine:</label>
        <input data-cy="type" type="text" name="type" onChange={handleChangeMachine} value={machine.type} />
        <SelectType name="type" type={machine.type} handleChange={handleChangeMachine}/>
      </div>
      <div className="form-group">
        <label>Speed:</label>
        <input data-cy="speed" type="number" name="speed" onChange={handleChangeMachine} value={machine.speed} />
      </div>
      <input data-cy="submit" type="submit" />
    </form>
  )
};

export default MachineForm;