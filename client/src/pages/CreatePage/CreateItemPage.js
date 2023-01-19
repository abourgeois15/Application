import React, {useState} from "react";
import { useApi } from "../../hooks/useApi";
import { useNavigate } from "react-router-dom";
import { services } from "../../services";
import ItemForm from "../../components/ui/ItemForm";

export const CreateItemPage = () => {
  const navigate = useNavigate();
  const navigateToItems = () => {
      // 👇️ navigate to /contacts
    navigate("/fullItems");
  };
  const [item, setItem] = useState({name: "", time: 0.0, recipe: [{number: 0, item: ""}, {number: 0, item: ""}, {number: 0, item: ""}], result: 1, machineType: ""});
  const [post, setPost] = useState(false);

  useApi(services.createItem, [], item, post)

  const handleSubmit = (event) => {
    event.preventDefault();
    item.time = Number(item.time)
    item.result = Number(item.result)
    item.recipe[0].number = Number(item.recipe[0].number)
    item.recipe[1].number = Number(item.recipe[0].number)
    item.recipe[2].number = Number(item.recipe[0].number)
    setPost(true)
  }

  const handleChangeItem = (event) => {
    setPost(false);
    setItem({
      ...item,
      [event.target.name]: event.target.value});
  };

  const handleChangeRecipe = (event) => {
    setPost(false);
    let recipe = item.recipe
    recipe[Number(event.target.id)][event.target.name] = event.target.value;
    setItem({
      ...item,
      ["recipe"]: recipe});
  };

  return (
    <div data-cy="detail-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Item Creation</h1>
      <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToItems}>Go Back To Items Page</button>
      <ItemForm item={item} handleSubmit={handleSubmit} handleChangeItem={handleChangeItem} handleChangeRecipe={handleChangeRecipe}/>
    </div>
  );
};