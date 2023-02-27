import React, {useState} from "react";
import { useParams } from 'react-router-dom'
import { useApi } from "../../hooks/useApi";
import { useNavigate } from "react-router-dom";
import { services } from "../../services";
import ItemForm from "../../components/ui/forms/ItemForm";

export const UpdateItemPage = () => {
  const navigate = useNavigate();
  const navigateToItems = () => {
      // 👇️ navigate to /contacts
    navigate("/fullItems");
  };
  const params = useParams()
  const {state: item, setState: setItem} = useApi(services.getItemByName, {name: "", time: 0.0, recipe: [{number: 0, item: ""}, {number: 0, item: ""}, {number: 0, item: ""}], result: 1, machineType: ""}, params.item_name);
  const [post, setPost] = useState(false);


  useApi(services.updateItem, [], item, post)

  const handleSubmit = (event) => {
    event.preventDefault();
    item.time = Number(item.time)
    item.result = Number(item.result)
    item.recipe[0].number = Number(item.recipe[0].number)
    item.recipe[1].number = Number(item.recipe[1].number)
    item.recipe[2].number = Number(item.recipe[2].number)
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
    const recipe = item.recipe.map((ingredient, index) => {
      if (index === Number(event.target.id)) {
        return {
          ...ingredient,
          [event.target.name]: event.target.value
        } 
      }
      else {
        return ingredient;
      }
    })
    setItem({
      ...item,
      "recipe": recipe});
  };
  
  return (
    <div data-cy="update-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Item Creation</h1>
      <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToItems}>Go Back To Items Page</button>
      <ItemForm item={item} handleSubmit={handleSubmit} handleChangeItem={handleChangeItem} handleChangeRecipe={handleChangeRecipe}/>
    </div>
  );
};