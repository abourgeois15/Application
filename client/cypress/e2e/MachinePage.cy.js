/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("Machine Component", () => {
  beforeEach(() => {
      cy.intercept("GET", "http://localhost:8080/machine/name/*", {fixture: "machine.json",})
      cy.visit("http://localhost:3000/fullMachines/name/Electric%20furnace")
  })

  it("AC1: Check that all required components are here", () => {

      cy.get('[data-cy="machine-page"]').should("exist");
      cy.get('[data-cy="header"]').should("exist");
      cy.get('[data-cy="A-goback-button"]').should("exist");
      cy.get('[data-cy="A-delete-button"]').should("exist");
      cy.get('[data-cy="A-update-button"]').should("exist");
      cy.get('[data-cy="machine-container"]').should("exist");
      cy.get('[data-cy="machine"]').should("exist");
      cy.get('[data-cy="name"]').should("exist");
      cy.get('[data-cy="type"]').should("exist");
      cy.get('[data-cy="time"]').should("exist");
      cy.get('[data-cy="recipe-container"]').should("exist");
      cy.get('[data-cy="ingredient"]').should("exist");
      cy.get('[data-cy="speed"]').should("exist");
  });

  it("AC2: Click on Back button and go to machine list page", () => {

      cy.intercept("GET", "http://localhost:8080/machines", {fixture: "machines.json",})
      cy.get('[data-cy="A-goback-button"]').click();
      cy.url().should("include", "/fullMachines");
      cy.get('[data-cy="machine-list-page"]').should("exist");
  });


  it("AC3: Click on Delete Machine button and go to delete page", () => {

      cy.intercept("DELETE", "http://localhost:8080/machine/*", {body: 1})
      cy.get('[data-cy="A-delete-button"]').click();
      cy.url().should("include", "/deleteMachine/Electric%20furnace"); 
      cy.get('[data-cy="delete-page"]').should("exist");
  });
  
  it("AC4: Click on Modify Machine button and go to update machine page", () => {

      cy.intercept("PUT", "http://localhost:8080/machine", {body: 1})
      cy.intercept("GET", "http://localhost:8080/machines", {fixture: "machines.json",})
      cy.intercept("GET", "http://localhost:8080/machines/type", {fixture: "types.json",})
      cy.get('[data-cy="A-update-button"]').click();
      cy.url().should("include", "/updateMachine/Electric%20furnace");
      cy.get('[data-cy="update-page"]').should("exist");
  });

  it("AC6: Click on Machine Type and go to other Machine Type List Page", () => {

    cy.intercept("GET", "http://localhost:8080//machine/type/*", {fixture: "machines_type.json",})
    cy.get('[data-cy="type"]').click();
    cy.url().should("include", "/fullMachines/type/Furnace");
    cy.get('[data-cy="machine-type-page"]').should("exist");
});
});
