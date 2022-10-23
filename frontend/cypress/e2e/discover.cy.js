
describe('login', () => {
  it('passes', () => {
    cy.visit('https://172.21.148.164/')
    //cy.visit('http://localhost:3000/')
    cy.get('input').eq(0).type("Tom").should('have.value', 'Tom')
    cy.get('input').eq(1).type("P@ssword").should('have.value', "P@ssword")
    cy.get('button').contains("Log").click()
    cy.wait(5000) // wait for 5 seconds
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
    cy.get('input').eq(0).click().type("Food").should('have.value', 'Food')
    cy.get('button').contains('Search').click()
  })
})

describe('opendiscoverpage', () => {
  it('passes', () => {
    cy.get('button').eq(0).click()
    cy.url().should('include', '/discover') // => true
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


describe('opendiscoverpage', () => {
  it('passes', () => {
    cy.get('button').eq(0).click()
    cy.url().should('include', '/discover') // => true
  })
})


describe('create new activity', () => {
  it('passes', () => {
    cy.get('button').contains("Create an activity").click()
    cy.wait(2000)
    cy.get('input').eq(1).click().type("Sentosa").should('have.value', 'Sentosa')
    cy.get('textarea').click().type("Island of fun").should('have.value', 'Island of fun')
    cy.get('input[type=file]').selectFile('./public/logo512.png')
    cy.get('button').eq(16).click()
    cy.get('button').eq(19).click()
    cy.get('button').eq(20).click()
    cy.get('button').eq(21).click()
    cy.get('input').eq(3).click().type("Sentosa, Singapore{enter}").should('have.value', 'Sentosa, Singapore')
    cy.get('button').contains("Save")
  })
})