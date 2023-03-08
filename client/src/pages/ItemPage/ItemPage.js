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

  const navigateToDelete = (id) => {
    // 👇️ navigate to /contacts
    navigate("/deleteItem/" + id);
  };

  const navigateToUpdate = (id) => {
    // 👇️ navigate to /contacts
    navigate("/updateItem/" + id);
  };

  const params = useParams()
  const {state: item} = useApi(services.getItemById, [], params.item_id);

  return (
    <div data-cy="item-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Item</h1>
        <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToItems}>Go Back To Items Page</button>
        <button data-cy="A-delete-button" className="buttonA buttonAA" onClick={() => {navigateToDelete(params.item_id)}}>Delete Item</button>
        <button data-cy="A-update-button" className="buttonA buttonAA" onClick={() => {navigateToUpdate(params.item_id)}}>Modify Item</button>
        <Item item={item}/>
    </div>
  );
};