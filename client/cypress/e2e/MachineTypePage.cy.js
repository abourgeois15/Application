/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("Machine Component", () => {
  beforeEach(() => {
    cy.intercept("GET", "http://localhost:8080//machine/type/*", {fixture: "machines_type.json",})
      cy.visit("http://localhost:3000/fullMachines/type/Furnace")
  })

  it("AC1: Check that all required components are here", () => {

      cy.get('[data-cy="machine-type-page"]').should("exist");
      cy.get('[data-cy="header"]').should("exist");
      cy.get('[data-cy="A-goback-button"]').should("exist");
      cy.get('[data-cy="machine-container"]').should("exist");
      cy.get('[data-cy="machine"]').should("exist");
  });

  it("AC2: Click on Back button and go to machine list page", () => {

    cy.intercept("GET", "http://localhost:8080/machines", {fixture: "machines.json",})
    cy.get('[data-cy="A-goback-button"]').click();
    cy.url().should("include", "/fullMachines");
    cy.get('[data-cy="machine-list-page"]').should("exist");
  });

  it("AC3: Click on machine and go to machine page", () => {

    cy.intercept("GET", "http://localhost:8080/machine/*", {fixture: "machine.json",})
    cy.get('[data-cy="Electric furnace_cy"]').click();
    cy.url().should("include", "/fullMachines/name/Electric%20furnace"); 
    cy.get('[data-cy="machine-page"]').should("exist");
  });
});
