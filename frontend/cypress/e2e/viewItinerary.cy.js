describe('login', () => {
  it('passes', () => {
    cy.visit('https://172.21.148.164/')
    //cy.visit('http://localhost:3000/')
    cy.get('input').eq(0).type("Tom").should('have.value', 'Tom')
    cy.get('input').eq(1).type("P@ssword").should('have.value', "P@ssword")
    cy.get('button').contains("Login").click()
  })
})

describe('createItineraryCheck', () => {
  it('passes', () => {
    cy.url().should('include', '/welcome') // => true
    cy.contains('Hiking')
    cy.get('button').contains("Plan my itinerary!").click()
    cy.wait(2000) // wait for 2 seconds
    cy.get('button').contains("Save changes made").click()
    cy.get('button').eq(7).click()
  })
})

describe('viewItinerary', () => {
  it('passes', () => {
    cy.get('button').eq(2).click()
    cy.url().should('include', '/itineraries') // => true
    cy.get('button').contains('From').click({force:true})
    //cy.url().should('include', '/edit-itinerary')
    cy.get('button').contains("Save changes made").click()
  })
})