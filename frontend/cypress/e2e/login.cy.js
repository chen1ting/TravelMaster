describe('signup', () => {
  it('passes', () => {
    cy.visit('https://172.21.148.164/')
    //cy.visit('http://localhost:3000/')
    cy.contains("Don't have an account? Sign Up Here").click()
    cy.url().should('include', '/Signup') // => true
    cy.get('input').eq(0).type("Tom").should('have.value', 'Tom')
    cy.get('input').eq(1).type("tom1122@gmail.com").should('have.value', "tom1122@gmail.com")
    cy.get('input').eq(2).type("P@ssword").should('have.value', "P@ssword")
    cy.get('input').eq(3).type("P@ssword", { delay: 200 }).should('have.value', "P@ssword")
    cy.get('input[type=file]').selectFile('./public/logo512.png')
    cy.get('button').contains("Create Account").click()
  })
})

describe('login', () => {
  it('passes', () => {
    cy.visit('https://172.21.148.164/')
    //cy.visit('http://localhost:3000/')
    cy.get('input').eq(0).type("Tom").should('have.value', 'Tom')
    cy.get('input').eq(1).type("P@ssword", { delay: 200 }).should('have.value', "P@ssword")
    cy.get('button').eq(0).click()
  })
})

