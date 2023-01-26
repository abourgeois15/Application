/* eslint-disable no-unreachable */
import React from "react";
import { useApi } from "../../../hooks/useApi";
import { services } from "../../../services";

const SelectItem = ({value, handleChange, id}) => {
  const {state: names} = useApi(services.getItems, [])
  return (
    <select data-cy={"select-item"+id} name="item" id={id} value={value} onChange={handleChange}>
      <option value="" key={0}></option>
      {names.map((name, index) => (
        <option value={name} key={index+1}>{name}</option>
      ))}
    </select>
  );
};

export default SelectItem;