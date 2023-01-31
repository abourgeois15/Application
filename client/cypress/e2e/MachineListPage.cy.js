/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("Machine List Component", () => {
  beforeEach(() => {
      cy.intercept("GET", "http://localhost:8080/machines", {fixture: "machines.json",})
      cy.visit("http://localhost:3000/fullMachines")
  })

  it("AC1: Check that all required components are here", () => {

      cy.get('[data-cy="machine-list-page"]').should("exist");
      cy.get('[data-cy="header"]').should("exist");
      cy.get('[data-cy="gohome-button"]').should("exist");
      cy.get('[data-cy="A-create-button"]').should("exist");
      cy.get('[data-cy="machine-container"]').should("exist");
      cy.get('[data-cy="machine"]').should("exist");
  });

  it("AC2: Click on Back button and go to home page", () => {

      cy.get('[data-cy="gohome-button"]').click();
      cy.url().should("include", "/");
      cy.get('[data-cy="home-page"]').should("exist");
  });

  it("AC3: Click on machine and go to machine page", () => {

      cy.intercept("GET", "http://localhost:8080/machine/*", {fixture: "machine.json",})
      cy.get('[data-cy="Electric furnace_cy"]').click();
      cy.url().should("include", "/fullMachines/name/Electric%20furnace"); 
      cy.get('[data-cy="machine-page"]').should("exist");
  });
  
  it("AC4: Click on Create Machine button and go to create machine page", () => {

      cy.intercept("GET", "http://localhost:8080/machines/type", {fixture: "types.json",})
      cy.get('[data-cy="A-create-button"]').click();
      cy.url().should("include", "/createMachine");
      cy.get('[data-cy="create-page"]').should("exist");
  });
});
