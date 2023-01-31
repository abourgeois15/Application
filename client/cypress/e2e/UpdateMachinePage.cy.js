/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("Update Machine Component", () => {
  beforeEach(() => {
      cy.intercept("GET", "http://localhost:8080/items", {fixture: "items_2.json",}).as("getMachines")
      cy.intercept("GET", "http://localhost:8080/machines/type", {fixture: "types.json",}).as("getTypes")
      cy.intercept("GET", "http://localhost:8080/machine/name/*", {fixture: "machine.json",}).as("getMachine")
      cy.visit("http://localhost:3000/updateMachine/Iron%20plate")
  })

  it("AC1: Check that all required components are here", () => {

      cy.get('[data-cy="update-page"]').should("exist");
      cy.get('[data-cy="header"]').should("exist");
      cy.get('[data-cy="A-goback-button"]').should("exist");
      cy.get('[data-cy="machine-form"]').should("exist");
  });

  it("AC2: Click on Back button and go to machine list page", () => {

      cy.intercept("GET", "http://localhost:8080/machines", {fixture: "machines.json",})
      cy.get('[data-cy="A-goback-button"]').click();
      cy.url().should("include", "/fullMachines");
      cy.get('[data-cy="machine-list-page"]').should("exist");
  });


  it("AC3: Update the Machine with the form and submit", () => {

      cy.intercept("PUT", "http://localhost:8080/machine", {body: 1}).as("putMachine")
      cy.wait('@getMachines').wait('@getTypes').wait('@getMachine')
      cy.get('[data-cy="name"]').clear().type("Assembling machine 2");
      cy.get('[data-cy="time"]').clear().type("0.5");
      cy.get('[data-cy="number0"]').clear().type("3");
      cy.get('[data-cy="select-item0"]').select("Electronic circuit");
      cy.get('[data-cy="number1"]').clear().type("5");
      cy.get('[data-cy="select-item1"]').select("Iron gear wheel");
      cy.get('[data-cy="number2"]').clear().type("2");
      cy.get('[data-cy="select-item2"]').select("Steel plate");
      cy.get('[data-cy="select-type"]').select("Assembling");
      cy.get('[data-cy="speed"]').clear().type("0.75");
      cy.get('[data-cy="submit"]').click();
      cy.wait('@putMachine').then(({request}) => {
          const machineJSON = JSON.parse(request.body)
          expect(machineJSON.name).to.eq("Assembling machine 2")
          expect(machineJSON.time).to.eq(0.5)
          expect(machineJSON.recipe[0].number).to.eq(3)
          expect(machineJSON.recipe[0].item).to.eq("Electronic circuit")
          expect(machineJSON.recipe[1].number).to.eq(5)
          expect(machineJSON.recipe[1].item).to.eq("Iron gear wheel")
          expect(machineJSON.recipe[2].number).to.eq(2)
          expect(machineJSON.recipe[2].item).to.eq("Steel plate")
          expect(machineJSON.speed).to.eq(0.75)
          expect(machineJSON.type).to.eq("Assembling")
      });
  });
});