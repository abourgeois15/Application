import React from "react";
import { useApi } from "../../hooks/useApi";
import { services } from "../../services";
import { useNavigate } from "react-router-dom";
import "./itemlistpage.css";
import ItemList from "../../components/ui/ItemList";

export const ItemListPage = () => {
  const navigate = useNavigate();

  const navigateToHome = () => {
    // 👇️ navigate to /
    navigate("/");
  };

  const [state] = useApi(services.getItems, []);

  return (
    <div data-cy="detail-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>List of Items</h1>
      <button data-cy="gohome-button" className="buttonG buttonGG" onClick={navigateToHome}>
        Go Back To Home
      </button>
      <ItemList items={state}/>
    </div>
  );
};