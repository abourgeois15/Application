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
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const createItem = async (item, post) => {
 
  const url = `${BASE_URL}/item`;
  const response = post && await fetch(url, {method: "POST", body: JSON.stringify(item)});
  if (response.ok) {
    return response.json();
  }
};

export const deleteItem = async (name) => {
 
  const url = `${BASE_URL}/item/${name}`;
  const response = await fetch(url, {method: "DELETE"});
  if (response.ok) {
    return response.json();
  }
};

export const updateItem = async (item, post) => {
 
  const url = `${BASE_URL}/item`;
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

export const getMachineByName = async (name) => {
 
  const url = `${BASE_URL}/machine/name/${name}`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const getMachineByType = async (type) => {
 
  const url = `${BASE_URL}/machine/type/${type}`;
  const response = await fetch(url);
  if (response.ok) {
    return response.json();
  }
};

export const createMachine = async (machine, post) => {
 
  const url = `${BASE_URL}/machine`;
  const response = post && await fetch(url, {method: "POST", body: JSON.stringify(machine)});
  if (response.ok) {
    return response.json();
  }
};

export const deleteMachine = async (name) => {
 
  const url = `${BASE_URL}/machine/${name}`;
  const response = await fetch(url, {method: "DELETE"});
  if (response.ok) {
    return response.json();
  }
};

export const updateMachine = async (machine, post) => {
 
  const url = `${BASE_URL}/machine`;
  const response = post && await fetch(url, {method: "PUT", body: JSON.stringify(machine)});
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
    getMachineByName,
    getMachineByType,
    createMachine,
    deleteMachine,
    updateMachine
};