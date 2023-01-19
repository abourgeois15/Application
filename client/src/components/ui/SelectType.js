/* eslint-disable no-unreachable */
import React from "react";
import "./item.css";

const SelectType = ({name, type, handleChange}) => {
    const typeOptions = ["", "Assembling", "Chemical", "Furnace", "Mining"]
    return (
    <select name={name} value={type} onChange={handleChange}>
      {typeOptions.map((type, index) => (
        <option value={type} key={index}>{type}</option>
      ))}
    </select>
  );
};

export default SelectType;