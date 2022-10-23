describe('login', () => {
  it('passes', () => {
    cy.visit('https://172.21.148.164/')
    //cy.visit('http://localhost:3000/')
    cy.get('input').eq(0).type("Tom").should('have.value', 'Tom')
    cy.get('input').eq(1).type("P@ssword").should('have.value', "P@ssword")
    cy.get('button').eq(7).click()
  })
})

/*
describe('openlandingpage', () => {
  it('passes', () => {
    cy.visit('https://172.21.148.164/')
    //cy.visit('http://localhost:3000/')
  })
})
*/


describe('welcomepagecheck', () => {
  it('passes', () => {
    cy.wait(2000) // wait for 2 seconds
    cy.url().should('include', '/welcome') // => true
    cy.contains('Hiking')
    cy.get('button').contains("Plan my itinerary!").click()
    cy.get('button').contains("Save changes made").click()
  })
})

