import React, {useState} from "react";
import { useApi } from "../../hooks/useApi";
import { useNavigate } from "react-router-dom";
import { services } from "../../services";

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
        <form className="form-vertical" onSubmit={handleSubmit}>
          <div className="form-group">
            <label>Name:</label>
            <input type="text" name="name" onChange={handleChangeItem} value={item.name} />
          </div>
          <div className="form-group">
            <label>Time:</label>
            <input type="number" name="time" onChange={handleChangeItem} value={item.time} />
          </div>
          <div className="form-group">
            <label>Recipe:</label>
            <div className="form-recipe">
              <label>Ingredient:</label>
              <div className="form-ingredient">
                <label>Number:</label>
                <input type="number" name="number" id="0" onChange={handleChangeRecipe} value={item.recipe[0].number} />
                <label>Item:</label>
                <input type="text" name="item" id="0" onChange={handleChangeRecipe} value={item.recipe[0].item} />
              </div>
              <div className="form-ingredient">
                <label>Number:</label>
                <input type="number" name="number" id="1" onChange={handleChangeRecipe} value={item.recipe[1].number} />
                <label>Item:</label>
                <input type="text" name="item" id="1" onChange={handleChangeRecipe} value={item.recipe[1].item} />
              </div>
              <div className="form-ingredient">
                <label>Number:</label>
                <input type="number" name="number" id="2" onChange={handleChangeRecipe} value={item.recipe[2].number} />
                <label>Item:</label>
                <input type="text" name="item" id="2" onChange={handleChangeRecipe} value={item.recipe[2].item} />
              </div>
            </div>
          </div>
          <div className="form-group">
            <label>Result:</label>
            <input type="number" name="result" onChange={handleChangeItem} value={item.result} />
          </div>
          <div className="form-group">
            <label>Type Of Machine:</label>
            <input type="text" name="machineType" onChange={handleChangeItem} value={item.machineType} />
          </div>
          <input type="submit" />
        </form>
    </div>
  );
};