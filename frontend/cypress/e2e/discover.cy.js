
describe('login', () => {
  it('passes', () => {
    cy.visit('https://172.21.148.164/')
    //cy.visit('http://localhost:3000/')
    cy.get('input').eq(0).type("Tom").should('have.value', 'Tom')
    cy.get('input').eq(1).type("P@ssword").should('have.value', "P@ssword")
    cy.wait(1000)
    cy.get('button').eq(0).click()
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

describe('opendiscoverpage', () => {
  it('passes', () => {
    cy.get('button').eq(0).click()
    cy.url().should('include', '/discover') // => true
  })
})


describe('searchbar', () => {
  it('passes', () => {
    cy.get('input').eq(0).type("Food").should('have.value', 'Food')
    cy.get('button').contains('Search').click()
    cy.wait(2000) // wait for 2 seconds
  })
})


describe('openactivitypage', () => {
  it('passes', () => {
        cy.contains('Food').click()
        cy.url().should('include', '/activity') // => true
  })
})

describe('inactivereport', () => {
  it('passes', () => {
    cy.get('button').contains("Report inactive").click()
    cy.get('button').contains("Remove inactive report").click()
    cy.get('button').contains("Report inactive").click()
    cy.get('button').contains("Remove inactive report").click()
  })
})


describe('back', () => {
  it('passes', () => {
    cy.go('back')    
  })
})
