/* eslint-disable no-unreachable */
import React from "react";
import { useApi } from "../../../hooks/useApi";
import { services } from "../../../services";

const SelectItem = ({value, handleChange, id}) => {
  const {state: items} = useApi(services.getItems, [])
  return (
    <select data-cy={"select-item"+id} name="item" id={id} value={value} onChange={handleChange}>
      <option value="" key={0}>{""}</option>
      {items.map((item, index) => (
        <option value={item.name} key={index+1}>{item.name}</option>
      ))}
    </select>
  );
};

export default SelectItem;