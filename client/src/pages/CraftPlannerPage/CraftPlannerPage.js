import React, {useState} from "react";
import { useParams } from 'react-router-dom'
import { useApi } from "../../hooks/useApi";
import { useNavigate } from "react-router-dom";
import { services } from "../../services";
import {CraftPlannerFormMain, CraftPlannerFormChild} from "../../components/ui/forms/CraftPlannerForm";

export const CraftPlannerPage = () => {
  const navigate = useNavigate();
  const navigateToItems = () => {
      // 👇️ navigate to /contacts
    navigate("/fullItems");
  };
  const params = useParams()
  const [plans, setPlans] = useState([{parentId: -1, item: params.item_name, number: 0, time: "s", machine: "", machines: [], numberMachine: 0.0, recipe: [{number: 0.0, item: ""}, {number: 0.0, item: ""}, {number: 0.0, item: ""}]}]);
  const {state: planResps} = useApi(services.getCraftPlanner, [{parentId: -1, item: params.item_name, number: 0, time: "s", machine: "", machines: [], numberMachine: 0.0,  recipe: [{number: 0.0, item: ""}, {number: 0.0, item: ""}, {number: 0.0, item: ""}]}], plans);

  const handleChange = (event) => {
    const nextPlans = plans.map((plan, index) => {
      if (index === Number(event.target.id)) {
        if (event.target.name === "number") {
          return {
            ...plan,
            [event.target.name]: parseInt(event.target.value)
          };
        }
        else {
          return {
            ...plan,
            [event.target.name]: event.target.value
          };
        }
      }
      else {
        return plan;
      }
    })
    setPlans(nextPlans)
  };

  const addCraftPlan = (parentId, item, number, time) => {
    console.log(parentId)
    setPlans([
      ...plans,
      {parentId: parentId, item: item, number: number, time: time, machine: "", machines: [], numberMachine: 0.0, recipe: [{number: 0.0, item: ""}, {number: 0.0, item: ""}, {number: 0.0, item: ""}]}
    ])
  }

  return (
    <div data-cy="update-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Craft Planner</h1>
      <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToItems}>Go Back To Items Page</button>
      {planResps && planResps.map((planResp, index) => {
        if (index === 0) {
          return <CraftPlannerFormMain plan={plans[index]} planResp={planResp} handleChange={handleChange} addCraftPlan={addCraftPlan} key={index} id={index}/>
        }
        else {
          return <CraftPlannerFormChild plan={plans[index]} planResp={planResp} handleChange={handleChange} addCraftPlan={addCraftPlan} key={index} id={index}/>
        }
      })}
    </div>
  );
};