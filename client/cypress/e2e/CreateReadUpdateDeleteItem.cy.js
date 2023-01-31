/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe('Create Read Update Delete Item', () => {
  beforeEach(() => {
    cy.request("POST", "http://localhost:8080/tables")
    cy.visit("http://localhost:3000/fullItems")
  })

  it("AC1: Go to Create Page, fill the form and submit to create item. Then go back to Items Page and click on the new item created to check", () => {

    cy.intercept("GET", "http://localhost:8080/items").as("getItems");
    cy.intercept("GET", "http://localhost:8080/machines/type").as("getTypes");
    cy.intercept("POST", "http://localhost:8080/item").as("postItem");
    cy.intercept("GET", "http://localhost:8080/item/*").as("getItem");

    // Go to create item page
    cy.get('[data-cy="A-create-button"]').click();
    cy.wait('@getItems').wait('@getTypes');

    // Type in form and submit
    cy.get('[data-cy="name"]').type("Inserter");
    cy.get('[data-cy="time"]').clear().type("0.5");
    cy.get('[data-cy="number0"]').clear().type("1");
    cy.get('[data-cy="select-item0"]').select("Electronic circuit");
    cy.get('[data-cy="number1"]').clear().type("1");
    cy.get('[data-cy="select-item1"]').select("Iron gear wheel");
    cy.get('[data-cy="number2"]').clear().type("1");
    cy.get('[data-cy="select-item2"]').select("Iron plate");
    cy.get('[data-cy="select-type"]').select("Assembling");
    cy.get('[data-cy="submit"]').click();
    cy.wait('@postItem');

    // Go back to item list page and check if new item is here
    cy.get('[data-cy="A-goback-button"]').click();
    cy.wait('@getItems');
    cy.get('[data-cy="Inserter_cy"]').should("exist");

    // Click on item and check if the information is as expected
    cy.get('[data-cy="Inserter_cy"]').click();
    cy.wait('@getItem');
    cy.contains('[data-cy="name"]', "Inserter")
    cy.contains('[data-cy="time"]', "0.5")
    cy.contains('[data-cy="result"]', "1")
    cy.contains('[data-cy="ingredient0"]', "1 Electronic circuit")
    cy.contains('[data-cy="ingredient1"]', "1 Iron gear wheel")
    cy.contains('[data-cy="ingredient2"]', "1 Iron plate")
    cy.contains('[data-cy="machine-type"]', "Assembling")

  });

  it("AC2: Go to Update Page, fill the form and submit to update item. Then go back to Items Page and click on the new item created to check", () => {

    cy.intercept("GET", "http://localhost:8080/items").as("getItems");
    cy.intercept("GET", "http://localhost:8080/machines/type").as("getTypes");
    cy.intercept("PUT", "http://localhost:8080/item").as("putItem");
    cy.intercept("GET", "http://localhost:8080/item/*").as("getItem");

    // Go to update item page
    cy.get('[data-cy="Transport belt_cy"]').click();
    cy.get('[data-cy="A-update-button"]').click();
    cy.wait('@getItem').wait('@getItems').wait('@getTypes');

    // Type in form and submit
    cy.get('[data-cy="name"]').clear().type("Electric furnace");
    cy.get('[data-cy="time"]').clear().type("5");
    cy.get('[data-cy="number0"]').clear().type("5");
    cy.get('[data-cy="select-item0"]').select("Advanced circuit");
    cy.get('[data-cy="number1"]').clear().type("10");
    cy.get('[data-cy="select-item1"]').select("Steel plate");
    cy.get('[data-cy="number2"]').clear().type("10");
    cy.get('[data-cy="select-item2"]').select("Stone brick");
    cy.get('[data-cy="select-type"]').select("Assembling");
    cy.get('[data-cy="submit"]').click();
    cy.wait('@putItem');

    // Go back to item list page and check if new item is here
    cy.get('[data-cy="A-goback-button"]').click();
    cy.wait('@getItems');
    cy.get('[data-cy="Electric furnace_cy"]').should("exist");
    cy.get('[data-cy="Transport belt_cy"]').should("not.exist");

    // Click on item and check if the information is as expected
    cy.get('[data-cy="Electric furnace_cy"]').click();
    cy.wait('@getItem');
    cy.contains('[data-cy="name"]', "Electric furnace")
    cy.contains('[data-cy="time"]', "5")
    cy.contains('[data-cy="result"]', "1")
    cy.contains('[data-cy="ingredient0"]', "5 Advanced circuit")
    cy.contains('[data-cy="ingredient1"]', "10 Steel plate")
    cy.contains('[data-cy="ingredient2"]', "10 Stone brick")
    cy.contains('[data-cy="machine-type"]', "Assembling")
  });

  it("AC3: Go to Item Page and delete. Then go back to Items Page and check that the item is not there anymore", () => {

    cy.intercept("GET", "http://localhost:8080/items").as("getItems");
    cy.intercept("DELETE", "http://localhost:8080/item/*").as("deleteItem");
    cy.intercept("GET", "http://localhost:8080/item/*").as("getItem");

    // Go to item page and delete
    cy.get('[data-cy="Advanced circuit_cy"]').click();
    cy.wait('@getItem');
    cy.get('[data-cy="A-delete-button"]').click();
    cy.wait('@deleteItem');

    // Go back to item list page and check if new item is here
    cy.get('[data-cy="A-goback-button"]').click();
    cy.wait('@getItems');
    cy.get('[data-cy="Advanced circuit_cy"]').should("not.exist");
  });
})