/*
describe('login', () => {
  it('passes', () => {
    cy.visit('https://172.21.148.164/')
    //cy.visit('http://localhost:3000/')
    cy.get('input').eq(0).type("Tom").should('have.value', 'Tom')
    cy.get('input').eq(1).type("P@ssword", { delay: 200 }).should('have.value', "P@ssword")
    cy.get('button').eq(0).click()
  })
})
*/

describe('openlandingpage', () => {
  it('passes', () => {
    cy.visit('https://172.21.148.164/')
    //cy.visit('http://localhost:3000/')
  })
})

describe('empty spec', () => {
  it('passes', () => {
    cy.url().should('include', '/welcome') // => true
    cy.contains('Hiking')
    cy.get('button').contains("Plan my itinerary!").click()
    cy.get('button').contains("Save changes made").click()
  })
})
