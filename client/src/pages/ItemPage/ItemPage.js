import React from "react";
import { useParams } from 'react-router-dom'
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import { useNavigate } from "react-router-dom";
import Item from "../../components/ui/Item";

export const ItemPage = () => {
  const navigate = useNavigate();
  const navigateToItems = () => {
    // 👇️ navigate to /contacts
    navigate("/fullItems");
  };

  const navigateToDelete = (name) => {
    // 👇️ navigate to /contacts
    navigate("/deleteItem/" + name);
  };

  const navigateToUpdate = (name) => {
    // 👇️ navigate to /contacts
    navigate("/updateItem/" + name);
  };

  const navigateToPlanner = (name) => {
    // 👇️ navigate to /contacts
    navigate("/craftPlanner/" + name);
  };
  const params = useParams()
  const {state: item} = useApi(services.getItemByName, [], params.item_name);

  return (
    <div data-cy="item-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Item</h1>
        <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToItems}>Go Back To Items Page</button>
        <button data-cy="A-delete-button" className="buttonA buttonAA" onClick={() => {navigateToDelete(params.item_name)}}>Delete Item</button>
        <button data-cy="A-update-button" className="buttonA buttonAA" onClick={() => {navigateToUpdate(params.item_name)}}>Modify Item</button>
        <button data-cy="A-plan-button" className="buttonA buttonAA" onClick={() => {navigateToPlanner(params.item_name)}}>Craft Planner</button>
        <Item item={item}/>
    </div>
  );
};