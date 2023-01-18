import { BASE_URL } from "./constants";

export const getItems = async () => {

  const url = `${BASE_URL}/items`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const getItemByName = async (name) => {
 
  const url = `${BASE_URL}/item/${name}`;
  console.log(url)
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const createItem = async (item, post) => {
 
  const url = `${BASE_URL}/item`;
  console.log(item)
  const response = post && await fetch(url, {method: "POST", body: JSON.stringify(item)});
  if (response.ok) {
    return response.json();
  }
};

export const deleteItem = async (name) => {
 
  const url = `${BASE_URL}/item/${name}`;
  console.log(url)
  const response = await fetch(url, {method: "DELETE"});
  if (response.ok) {
    return response.json();
  }
};

export const updateItem = async (item, post) => {
 
  const url = `${BASE_URL}/item`;
  console.log(url)
  const response = post && await fetch(url, {method: "PUT", body: JSON.stringify(item)});
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
    createItem,
    deleteItem,
    updateItem,
    getMachines,
    getMachineById
};