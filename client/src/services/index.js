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

export const getMachineByName = async (name) => {
 
  const url = `${BASE_URL}/machines/${name}`;
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
    getMachineByName
};