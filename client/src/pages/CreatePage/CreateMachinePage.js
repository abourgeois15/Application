import React, {useState} from "react";
import { useApi } from "../../hooks/useApi";
import { useNavigate } from "react-router-dom";
import { services } from "../../services";
import MachineForm from "../../components/ui/MachineForm";

export const CreateMachinePage = () => {
  const navigate = useNavigate();
  const navigateToMachines = () => {
      // 👇️ navigate to /contacts
    navigate("/fullMachines");
  };
  const [machine, setMachine] = useState({name: "", time: 0.0, recipe: [{number: 0, machine: ""}, {number: 0, machine: ""}, {number: 0, machine: ""}], type: "", speed: 0});
  const [post, setPost] = useState(false);

  useApi(services.createMachine, [], machine, post)

  const handleSubmit = (event) => {
    event.preventDefault();
    machine.time = Number(machine.time)
    machine.recipe[0].number = Number(machine.recipe[0].number)
    machine.recipe[1].number = Number(machine.recipe[0].number)
    machine.recipe[2].number = Number(machine.recipe[0].number)
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
      <MachineForm machine={machine} handleSubmit={handleSubmit} handleChangeMachine={handleChangeMachine} handleChangeRecipe={handleChangeRecipe}/>
    </div>
  );
};