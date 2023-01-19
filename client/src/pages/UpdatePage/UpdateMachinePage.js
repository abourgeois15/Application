import React, {useState} from "react";
import { useParams } from 'react-router-dom'
import { useApi } from "../../hooks/useApi";
import { useNavigate } from "react-router-dom";
import { services } from "../../services";

export const UpdateMachinePage = () => {
  const navigate = useNavigate();
  const navigateToMachines = () => {
      // 👇️ navigate to /contacts
    navigate("/fullMachines");
  };
  const params = useParams()
  const {state: machine, setState: setMachine} = useApi(services.getMachineByName, {name: "", time: 0.0, recipe: [{number: 0, machine: ""}, {number: 0, machine: ""}, {number: 0, machine: ""}], type: 1, speed: ""}, params.machine_name);
  const [post, setPost] = useState(false);

  useApi(services.updateMachine, [], machine, post)

  const handleSubmit = (event) => {
    event.preventDefault();
    machine.time = Number(machine.time)
    machine.recipe[0].number = Number(machine.recipe[0].number)
    machine.recipe[1].number = Number(machine.recipe[1].number)
    machine.recipe[2].number = Number(machine.recipe[2].number)
    machine.speed = Number(machine.speed)
    setPost(true)
  }

  const handleChangeMachine = (event) => {
    setPost(false);
    setMachine({
      ...machine,
      [event.target.name]: event.target.value});
  };

  const handleChangeRecipe = (event) => {
    setPost(false);
    let recipe = machine.recipe
    recipe[Number(event.target.id)][event.target.name] = event.target.value;
    setMachine({
      ...machine,
      ["recipe"]: recipe});
  };
  
  return (
    <div data-cy="detail-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Machine Creation</h1>
        <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToMachines}>Go Back To Machines Page</button>
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
            <input type="text" name="type" onChange={handleChangeMachine} value={machine.type} />
          </div>
          <div className="form-group">
            <label>Speed:</label>
            <input type="number" name="speed" onChange={handleChangeMachine} value={machine.speed} />
          </div>
          <input type="submit" />
        </form>
    </div>
  );
};