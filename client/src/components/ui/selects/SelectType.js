/* eslint-disable no-unreachable */
import React from "react";
import { useApi } from "../../../hooks/useApi";
import { services } from "../../../services";

const SelectType = ({name, type, handleChange}) => {
  const {state: types} = useApi(services.getTypes, [])
  return (
    <select name={name} value={type} onChange={handleChange}>
      <option value="" key={0}>{""}</option>
      {types.map((type, index) => (
        <option value={type} key={index+1}>{type}</option>
      ))}
    </select>
  );
};

export default SelectType;