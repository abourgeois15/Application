import React, {useState} from "react";
import { useParams } from 'react-router-dom'
import { useApi } from "../../hooks/useApi";
import { useNavigate } from "react-router-dom";
import { services } from "../../services";
import MachineForm from "../../components/ui/forms/MachineForm";

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
    machine.recipe = machine.recipe.map((ingredient) => (
      {
        ...ingredient,
        "number": Number(ingredient.number)
      }
    ))
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
    const recipe = machine.recipe.map((ingredient, index) => {
      if (index === Number(event.target.id)) {
        return {
          ...ingredient,
          [event.target.name]: event.target.value
        } 
      }
      else {
        return ingredient;
      }
    })
    setMachine({
      ...machine,
      "recipe": recipe});
  };

  const addIngredient = () => {
    setPost(false);
    setMachine({
      ...machine,
      "recipe": [...machine.recipe, {id: -1, number: 0, item: ""}]
    })
  }

  const deleteIngredient = (ingredient) => {
    setPost(false);
    if (ingredient.id === -1) {  
      setMachine({
        ...machine,
        "recipe": machine.recipe.filter(_ingredient => _ingredient !== ingredient)
      })
    }
    else {
      const recipe = machine.recipe.map((_ingredient) => {
        if (_ingredient === ingredient) {
          return {
            ..._ingredient,
            "number": -1
          } 
        }
        else {
          return _ingredient;
        }
      })
      setMachine({
        ...machine,
        "recipe": recipe});
    }
  }
  
  return (
    <div data-cy="update-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Machine Creation</h1>
      <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToMachines}>Go Back To Machines Page</button>
      <MachineForm machine={machine} handleSubmit={handleSubmit} handleChangeMachine={handleChangeMachine} handleChangeRecipe={handleChangeRecipe} addIngredient={addIngredient} deleteIngredient={deleteIngredient}/>
    </div>
  );
};