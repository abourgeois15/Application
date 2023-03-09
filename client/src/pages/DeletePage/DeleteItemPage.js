import React from "react";
import { useParams } from 'react-router-dom'
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import { useNavigate } from "react-router-dom";

export const DeleteItemPage = () => {
  const navigate = useNavigate();
  const navigateToItems = () => {
      // 👇️ navigate to /contacts
    navigate("/fullItems");
    };
  const params = useParams()
  useApi(services.deleteItem, [], params.item_id);

  return (
    <div data-cy="delete-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Item {params.item_name} Deleted</h1>
        <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToItems}>Go Back To Items Page</button>
    </div>
  );
};