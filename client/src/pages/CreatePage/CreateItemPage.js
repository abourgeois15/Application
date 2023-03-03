import React, {useState} from "react";
import { useApi } from "../../hooks/useApi";
import { useNavigate } from "react-router-dom";
import { services } from "../../services";
import ItemForm from "../../components/ui/forms/ItemForm";

export const CreateItemPage = () => {
  const navigate = useNavigate();
  const navigateToItems = () => {
      // 👇️ navigate to /contacts
    navigate("/fullItems");
  };
  const [item, setItem] = useState({name: "", time: 0.0, recipe: [], result: 1, machineType: ""});
  const [post, setPost] = useState(false);

  useApi(services.createItem, [], item, post)

  const handleSubmit = (event) => {
    event.preventDefault();
    item.time = Number(item.time)
    item.result = Number(item.result)
    item.recipe = item.recipe.map((ingredient) => (
      {
        ...ingredient,
        "number": Number(ingredient.number)
      }
    ))
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

  const addIngredient = () => {
    setPost(false);
    setItem({
      ...item,
      "recipe": [...item.recipe, {id: -1, number: 0, item: ""}]
    })
  }

  const deleteIngredient = (ingredient) => {
    setPost(false);
    if (ingredient.id === -1) {  
      setItem({
        ...item,
        "recipe": item.recipe.filter(_ingredient => _ingredient !== ingredient)
      })
    }
    else {
      const recipe = item.recipe.map((_ingredient) => {
        if (_ingredient === ingredient) {
          return {
            ..._ingredient,
            "number": -1
          } 
        }
        else {
          return _ingredient;
        }
      })
      setItem({
        ...item,
        "recipe": recipe});
    }
  }

  return (
    <div data-cy="create-page" style={{ textAlign: "center" }}>
      <h1  data-cy="header" style={{ textAlign: "center"}}>Item Creation</h1>
      <button data-cy="A-goback-button" className="buttonA buttonAA" onClick={navigateToItems}>Go Back To Items Page</button>
      <ItemForm item={item} handleSubmit={handleSubmit} handleChangeItem={handleChangeItem} handleChangeRecipe={handleChangeRecipe} addIngredient={addIngredient} deleteIngredient={deleteIngredient}/>
    </div>
  );
};