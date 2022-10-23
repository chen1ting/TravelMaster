describe('login', () => {
    it('passes', () => {
      cy.visit('https://172.21.148.164/')
      //cy.visit('http://localhost:3000/')
      cy.get('input').eq(0).type("Tom").should('have.value', 'Tom')
      cy.get('input').eq(1).type("P@ssword").should('have.value', "P@ssword")
      cy.wait(1000)
    })
  })
  
  describe('openfeedbackpage', () => {
    it('passes', () => {
      cy.get('button').eq(3).click()
      cy.url().should('include', '/feedback') // => true
      //cy.get('input').eq(0).type("Feedback Description").should('have.value', 'Feedback Description')
      cy.get('button').eq(5).click()

    })
  })