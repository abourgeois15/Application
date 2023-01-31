/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe('Create Read Update Delete Machine', () => {
  beforeEach(() => {
    cy.request("POST", "http://localhost:8080/tables")
    cy.visit("http://localhost:3000/fullMachines")
  })

  it("AC1: Go to Create Page, fill the form and submit to create machine. Then go back to Machines Page and click on the new machine created to check", () => {

    cy.intercept("GET", "http://localhost:8080/items").as("getItems");
    cy.intercept("GET", "http://localhost:8080/machines").as("getMachines");
    cy.intercept("GET", "http://localhost:8080/machines/type").as("getTypes");
    cy.intercept("POST", "http://localhost:8080/machine").as("postMachine");
    cy.intercept("GET", "http://localhost:8080/machine/name/*").as("getMachine");

    // Go to create machine page
    cy.get('[data-cy="A-create-button"]').click();
    cy.wait('@getItems').wait('@getTypes');

    // Type in form and submit
    cy.get('[data-cy="name"]').type("Chemical plant");
    cy.get('[data-cy="time"]').clear().type("5");
    cy.get('[data-cy="number0"]').clear().type("5");
    cy.get('[data-cy="select-item0"]').select("Electronic circuit");
    cy.get('[data-cy="number1"]').clear().type("5");
    cy.get('[data-cy="select-item1"]').select("Iron gear wheel");
    cy.get('[data-cy="number2"]').clear().type("5");
    cy.get('[data-cy="select-item2"]').select("Steel plate");
    cy.get('[data-cy="type"]').type("Chemical");
    cy.get('[data-cy="speed"]').clear().type("1");
    cy.get('[data-cy="submit"]').click();
    cy.wait('@postMachine');

    // Go back to machine list page and check if new machine is here
    cy.get('[data-cy="A-goback-button"]').click();
    cy.wait('@getMachines');
    cy.get('[data-cy="Chemical plant_cy"]').should("exist");

    // Click on machine and check if the information is as expected
    cy.get('[data-cy="Chemical plant_cy"]').click();
    cy.wait('@getMachine');
    cy.contains('[data-cy="name"]', "Chemical plant")
    cy.contains('[data-cy="time"]', "5")
    cy.contains('[data-cy="ingredient0"]', "5 Electronic circuit")
    cy.contains('[data-cy="ingredient1"]', "5 Iron gear wheel")
    cy.contains('[data-cy="ingredient2"]', "5 Steel plate")
    cy.contains('[data-cy="type"]', "Chemical")
    cy.contains('[data-cy="speed"]', "1")

  });

  it("AC2: Go to Update Page, fill the form and submit to update machine. Then go back to Machines Page and click on the new machine created to check", () => {

    cy.intercept("GET", "http://localhost:8080/items").as("getItems");
    cy.intercept("GET", "http://localhost:8080/machines").as("getMachines");
    cy.intercept("GET", "http://localhost:8080/machines/type").as("getTypes");
    cy.intercept("PUT", "http://localhost:8080/machine").as("putMachine");
    cy.intercept("GET", "http://localhost:8080/machine/name/*").as("getMachine");

    // Go to update machine page
    cy.get('[data-cy="Electric furnace_cy"]').click();
    cy.get('[data-cy="A-update-button"]').click();
    cy.wait('@getMachine').wait('@getMachines').wait('@getTypes');

    // Type in form and submit
    cy.get('[data-cy="name"]').clear().type("Pumpjack");
    cy.get('[data-cy="time"]').clear().type("5");
    cy.get('[data-cy="number0"]').clear().type("5");
    cy.get('[data-cy="select-item0"]').select("Electronic circuit");
    cy.get('[data-cy="number1"]').clear().type("10");
    cy.get('[data-cy="select-item1"]').select("Iron gear wheel");
    cy.get('[data-cy="number2"]').clear().type("5");
    cy.get('[data-cy="select-item2"]').select("Steel plate");
    cy.get('[data-cy="type"]').clear().type("Pumping");
    cy.get('[data-cy="speed"]').clear().type("1");
    cy.get('[data-cy="submit"]').click();
    cy.wait('@putMachine');

    // Go back to machine list page and check if new machine is here
    cy.get('[data-cy="A-goback-button"]').click();
    cy.wait('@getMachines');
    cy.get('[data-cy="Pumpjack_cy"]').should("exist");
    cy.get('[data-cy="Electric furnace_cy"]').should("not.exist");

    // Click on machine and check if the information is as expected
    cy.get('[data-cy="Pumpjack_cy"]').click();
    cy.wait('@getMachine');
    cy.contains('[data-cy="name"]', "Pumpjack")
    cy.contains('[data-cy="time"]', "5")
    cy.contains('[data-cy="ingredient0"]', "5 Electronic circuit")
    cy.contains('[data-cy="ingredient1"]', "10 Iron gear wheel")
    cy.contains('[data-cy="ingredient2"]', "5 Steel plate")
    cy.contains('[data-cy="type"]', "Pumping")
    cy.contains('[data-cy="speed"]', "1")
  });

  it("AC3: Go to Machine Page and delete. Then go back to Machines Page and check that the machine is not there anymore", () => {

    cy.intercept("GET", "http://localhost:8080/machines").as("getMachines");
    cy.intercept("DELETE", "http://localhost:8080/machine/*").as("deleteMachine");
    cy.intercept("GET", "http://localhost:8080/machine/name/*").as("getMachine");

    // Go to machine page and delete
    cy.get('[data-cy="Assembling machine 1_cy"]').click();
    cy.wait('@getMachine');
    cy.get('[data-cy="A-delete-button"]').click();
    cy.wait('@deleteMachine');

    // Go back to machine list page and check if new machine is here
    cy.get('[data-cy="A-goback-button"]').click();
    cy.wait('@getMachines');
    cy.get('[data-cy="Advanced circuit_cy"]').should("not.exist");
  });
})