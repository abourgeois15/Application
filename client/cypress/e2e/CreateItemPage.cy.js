/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("Create Item Component", () => {
    beforeEach(() => {
        cy.intercept("GET", "http://localhost:8080/items", {fixture: "items_1.json",}).as("getItems")
        cy.intercept("GET", "http://localhost:8080/machines/type", {fixture: "types.json",}).as("getTypes")
        cy.visit("http://localhost:3000/createItem")
    })

    it("AC1: Check that all required components are here", () => {

        cy.get('[data-cy="create-page"]').should("exist");
        cy.get('[data-cy="header"]').should("exist");
        cy.get('[data-cy="A-goback-button"]').should("exist");
        cy.get('[data-cy="item-form"]').should("exist");
    });

    it("AC2: Click on Back button and go to item list page", () => {

        cy.get('[data-cy="A-goback-button"]').click();
        cy.url().should("include", "/fullItems");
        cy.get('[data-cy="item-list-page"]').should("exist");
    });
  

    it("AC3: Create an Item with the form and submit", () => {

        cy.intercept("POST", "http://localhost:8080/item", {body: 1}).as("postItem")
        cy.wait('@getItems').wait('@getTypes')
        cy.get('[data-cy="name"]').type("Iron gear wheel");
        cy.get('[data-cy="time"]').clear().type("0.5");
        cy.get('[data-cy="number0"]').clear().type("2");
        cy.get('[data-cy="select-item0"]').select("Iron plate");
        cy.get('[data-cy="select-type"]').select("Assembling");
        cy.get('[data-cy="submit"]').click();
        cy.wait('@postItem').then(({request}) => {
            const itemJSON = JSON.parse(request.body)
            expect(itemJSON.name).to.eq("Iron gear wheel")
            expect(itemJSON.time).to.eq(0.5)
            expect(itemJSON.recipe[0].number).to.eq(2)
            expect(itemJSON.recipe[0].item).to.eq("Iron plate")
            expect(itemJSON.recipe[1].number).to.eq(0)
            expect(itemJSON.recipe[1].item).to.eq("")
            expect(itemJSON.recipe[2].number).to.eq(0)
            expect(itemJSON.recipe[2].item).to.eq("")
            expect(itemJSON.result).to.eq(1)
            expect(itemJSON.machineType).to.eq("Assembling")
        });
    });
});