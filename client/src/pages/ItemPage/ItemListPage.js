import React, { useState } from "react";
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import { useNavigate } from "react-router-dom";
import ItemList from "../../components/ui/ItemList";
import SearchBox from "../../components/ui/SearchBox";

export const ItemListPage = () => {
  const navigate = useNavigate();

  const navigateToHome = () => {
    // 👇️ navigate to /
    navigate("/");
  };

  const navigateToCreate = () => {
    // 👇️ navigate to /
    navigate("/createItem");
  };

  const {state: names} = useApi(services.getItems, []);
  console.log(names)
  const [value, setValue] = useState("");

  //handleChange function for search
  const searchHandleChange = (event) => {
    setValue(event.target.value);
  };

  return (
    <div data-cy="detail-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>List of Items</h1>
      <SearchBox handleChange={searchHandleChange} value={value} />
      <button data-cy="gohome-button" className="buttonG buttonGG" onClick={navigateToHome}>Go Back To Home</button>
      <button data-cy="A-create-button" className="buttonA buttonAA" onClick={navigateToCreate}>Create New Item</button>
      <ItemList names={names} value={value} />
    </div>
  );
};