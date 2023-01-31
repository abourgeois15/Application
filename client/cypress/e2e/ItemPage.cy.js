/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("Item Component", () => {
    beforeEach(() => {
        cy.intercept("GET", "http://localhost:8080/item/Iron%20plate", {fixture: "item_1.json",})
        cy.visit("http://localhost:3000/fullItems/Iron%20plate")
    })

    it("AC1: Check that all required components are here", () => {

        cy.get('[data-cy="item-page"]').should("exist");
        cy.get('[data-cy="header"]').should("exist");
        cy.get('[data-cy="A-goback-button"]').should("exist");
        cy.get('[data-cy="A-delete-button"]').should("exist");
        cy.get('[data-cy="A-update-button"]').should("exist");
        cy.get('[data-cy="item-container"]').should("exist");
        cy.get('[data-cy="item"]').should("exist");
        cy.get('[data-cy="name"]').should("exist");
        cy.get('[data-cy="machine-type"]').should("exist");
        cy.get('[data-cy="time"]').should("exist");
        cy.get('[data-cy="recipe-container"]').should("exist");
        cy.get('[data-cy="ingredient"]').should("exist");
        cy.get('[data-cy="result"]').should("exist");
    });

    it("AC2: Click on Back button and go to item list page", () => {

        cy.intercept("GET", "http://localhost:8080/items", {fixture: "items_1.json",})
        cy.get('[data-cy="A-goback-button"]').click();
        cy.url().should("include", "/fullItems");
        cy.get('[data-cy="item-list-page"]').should("exist");
    });
  

    it("AC3: Click on Delete Item button and go to delete page", () => {

        cy.intercept("DELETE", "http://localhost:8080/item/*", {body: 1})
        cy.get('[data-cy="A-delete-button"]').click();
        cy.url().should("include", "/deleteItem/Iron%20plate"); 
        cy.get('[data-cy="delete-page"]').should("exist");
    });
    
    it("AC4: Click on Modify Item button and go to update item page", () => {

        cy.intercept("PUT", "http://localhost:8080/item", {body: 1})
        cy.intercept("GET", "http://localhost:8080/items", {fixture: "items_1.json",})
        cy.intercept("GET", "http://localhost:8080/machines/type", {fixture: "types.json",})
        cy.get('[data-cy="A-update-button"]').click();
        cy.url().should("include", "/updateItem/Iron%20plate");
        cy.get('[data-cy="update-page"]').should("exist");
    });

    it("AC5: Click on Item in recipe go to other item page", () => {

        cy.intercept("GET", "http://localhost:8080/item/Iron%20ore", {fixture: "item_2.json",})
        cy.get('[data-cy="ingredient0"]').click();
        cy.url().should("include", "/fullItems/Iron%20ore");
        cy.get('[data-cy="item-page"]').should("exist");
    });

    it("AC6: Click on Machine Type and go to other Machine Type List Page", () => {

        cy.intercept("GET", "http://localhost:8080//machine/type/*", {fixture: "machines_type.json",})
        cy.get('[data-cy="machine-type"]').click();
        cy.url().should("include", "/fullMachines/type/Furnace");
        cy.get('[data-cy="machine-type-page"]').should("exist");
    });
  });
  