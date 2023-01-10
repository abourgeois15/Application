import { BASE_URL } from "./constants";

export const getItems = async () => {

  const url = `${BASE_URL}/items`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const getItemById = async (id) => {
 
  const url = `${BASE_URL}/items/${id}`;
  console.log(url)
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const services = {
    getItems,
    getItemById
};