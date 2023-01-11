import { BASE_URL } from "./constants";

export const getItems = async () => {

  const url = `${BASE_URL}/items`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const getItemByName = async (name) => {
 
  const url = `${BASE_URL}/items/${name}`;
  console.log(url)
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const getMachines = async () => {

  const url = `${BASE_URL}/machines`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const getMachineById = async (id) => {
 
  const url = `${BASE_URL}/machines/${id}`;
  console.log(url)
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const services = {
    getItems,
    getItemByName,
    getMachines,
    getMachineById
};