import React from "react";

export const CraftPlannerFormMain = ({plan, planResp, handleChange, addCraftPlan, id}) => {
    return (
      <form data-cy="item-form" className="form-inline">
        <div className="form-group">
          <label>{planResp.item}</label>
          <input data-cy="number" type="number" name="number" onChange={handleChange} value={plan.number} id={id}/>
          /
          <select data-cy="select-time" name="time" value={plan.time} onChange={handleChange} id={id}>
            <option value="s">s</option>
            <option value="min">min</option>
            <option value="h">h</option>
          </select>
          {planResp && planResp.numberMachine}
          <select data-cy="select-machine" name="machine" value={plan.machine} onChange={handleChange} id={id}>
            {planResp.machines && planResp.machines.map((machine, index) => (
              <option value={machine} key={index}>{machine}</option>
            ))}
          </select>
          {planResp.recipe[0].item !== "" && <p data-cy="ingredient0" onClick={() => addCraftPlan(Number(id), planResp.recipe[0].item, planResp.recipe[0].number, plan.time)}>{planResp && planResp.recipe[0].number}/{plan.time} {planResp.recipe[0].item}</p>}
          {planResp.recipe[1].item !== "" && <p data-cy="ingredient0" onClick={() => addCraftPlan(Number(id), planResp.recipe[1].item, planResp.recipe[1].number, plan.time)}>{planResp && planResp.recipe[1].number}/{plan.time} {planResp.recipe[1].item}</p>}
          {planResp.recipe[2].item !== "" && <p data-cy="ingredient0" onClick={() => addCraftPlan(Number(id), planResp.recipe[2].item, planResp.recipe[2].number, plan.time)}>{planResp && planResp.recipe[2].number}/{plan.time} {planResp.recipe[2].item}</p>}
        </div>
      </form>
    )
}

export const CraftPlannerFormChild = ({plan, planResp, handleChange, addCraftPlan, id}) => {
  return (
    <form data-cy="item-form" className="form-inline">
      <div className="form-group">
        <label>{planResp.item}</label>
        {planResp.number}
        /
        <select data-cy="select-time" name="time" value={plan.time} onChange={handleChange} id={id}>
          <option value="s">s</option>
          <option value="min">min</option>
          <option value="h">h</option>
        </select>
        {planResp && planResp.numberMachine}
        <select data-cy="select-machine" name="machine" value={plan.machine} onChange={handleChange} id={id}>
          {planResp.machines && planResp.machines.map((machine, index) => (
            <option value={machine} key={index}>{machine}</option>
          ))}
        </select>
        {planResp.recipe[0].item !== "" && <p data-cy="ingredient0" onClick={() => addCraftPlan(Number(id), planResp.recipe[0].item, planResp.recipe[0].number, plan.time)}>{planResp && planResp.recipe[0].number}/{plan.time} {planResp.recipe[0].item}</p>}
        {planResp.recipe[1].item !== "" && <p data-cy="ingredient0" onClick={() => addCraftPlan(Number(id), planResp.recipe[1].item, planResp.recipe[1].number, plan.time)}>{planResp && planResp.recipe[1].number}/{plan.time} {planResp.recipe[1].item}</p>}
        {planResp.recipe[2].item !== "" && <p data-cy="ingredient0" onClick={() => addCraftPlan(Number(id), planResp.recipe[2].item, planResp.recipe[2].number, plan.time)}>{planResp && planResp.recipe[2].number}/{plan.time} {planResp.recipe[2].item}</p>}
      </div>
    </form>
  )
}