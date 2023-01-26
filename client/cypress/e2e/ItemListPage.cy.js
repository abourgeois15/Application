/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("Item List Component", () => {
    beforeEach(() => {
        cy.intercept("GET", "http://localhost:8080/items", {fixture: "items.json",})
        cy.visit("http://localhost:3000/fullItems")
    })

    it("AC1: Check that all required components are here", () => {

        cy.get('[data-cy="item-list-page"]').should("exist");
        cy.get('[data-cy="header"]').should("exist");
        cy.get('[data-cy="item-search-box"]').should("exist");
        cy.get('[data-cy="gohome-button"]').should("exist");
        cy.get('[data-cy="A-create-button"]').should("exist");
        cy.get('[data-cy="item-container"]').should("exist");
        cy.get('[data-cy="item"]').should("exist");
    });

    it("AC2: Click on Back button and go to home page", () => {

        cy.get('[data-cy="gohome-button"]').click();
        cy.url().should("include", "/");
        cy.get('[data-cy="home-page"]').should("exist");
    });
  
    it("AC3: Use the search box", () => {
        cy.get('[data-cy="item-search-box"]').type("iron");
        cy.get('[data-cy="Iron plate_cy"]').should("exist");
        cy.get('[data-cy="Iron ore_cy"]').should("exist");
        cy.get('[data-cy="Copper plate_cy"]').should("not.exist");
        cy.get('[data-cy="item-search-box"]').type(" plate");
        cy.get('[data-cy="Iron plate_cy"]').should("exist");
        cy.get('[data-cy="Iron ore_cy"]').should("not.exist");
        cy.get('[data-cy="Copper plate_cy"]').should("not.exist");
    });

    it("AC4: Click on item and go to item page", () => {

        cy.intercept("GET", "http://localhost:8080/item/*", {fixture: "item.json",})
        cy.get('[data-cy="Iron plate_cy"]').click();
        cy.url().should("include", "/fullItems/Iron%20plate"); 
        cy.get('[data-cy="item-page"]').should("exist");
      });
    
    it("AC5: Click on Create Item button and go to create item page", () => {

        cy.intercept("GET", "http://localhost:8080/machines/type", {fixture: "types.json",})
        cy.get('[data-cy="A-create-button"]').click();
        cy.url().should("include", "/createItem");
        cy.get('[data-cy="create-page"]').should("exist");
    });
  });
  