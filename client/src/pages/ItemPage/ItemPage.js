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
  const params = useParams()
  const [state] = useApi(services.getItemByName, [], params.item_name);

  return (
    <div data-cy="detail-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Item</h1>
        <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToItems}>Go Back Items Page</button>
        <Item item={state}/>
    </div>
  );
};