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
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const createItem = async (item, post) => {
 
  const url = `${BASE_URL}/items`;
  const response = post && await fetch(url, {method: "POST", body: JSON.stringify(item)});
  if (response.ok) {
    return response.json();
  }
};

export const deleteItem = async (name) => {
 
  const url = `${BASE_URL}/items/${name}`;
  const response = await fetch(url, {method: "DELETE"});
  if (response.ok) {
    return response.json();
  }
};

export const updateItem = async (item, post) => {
  console.log(post)
  const url = `${BASE_URL}/items/${item.id}`;
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

export const getTypes = async () => {

  const url = `${BASE_URL}/machine-types`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const getMachineById = async (id) => {
 
  const url = `${BASE_URL}/machines/${id}`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const getMachineByType = async (type) => {
 
  const url = `${BASE_URL}/machines/${type}`;
  const response = type && await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const createMachine = async (machine, post) => {
 
  const url = `${BASE_URL}/machines`;
  const response = post && await fetch(url, {method: "POST", body: JSON.stringify(machine)});
  if (response.ok) {
    return response.json();
  }
};

export const deleteMachine = async (id) => {
 
  const url = `${BASE_URL}/machines/${id}`;
  const response = await fetch(url, {method: "DELETE"});
  if (response.ok) {
    return response.json();
  }
};

export const updateMachine = async (machine, post) => {
 
  const url = `${BASE_URL}/machines/${machine.id}`;
  const response = post && await fetch(url, {method: "PUT", body: JSON.stringify(machine)});
  if (response.ok) {
    return response.json();
  }
};

export const services = {
    getItems,
    getItemById,
    createItem,
    deleteItem,
    updateItem,
    getMachines,
    getTypes,
    getMachineById,
    getMachineByType,
    createMachine,
    deleteMachine,
    updateMachine
};