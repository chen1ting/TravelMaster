describe('login', () => {
    it('passes', () => {
      cy.visit('https://172.21.148.164/')
      //cy.visit('http://localhost:3000/')
      cy.get('input').eq(0).type("Tom").should('have.value', 'Tom')
      cy.get('input').eq(1).type("P@ssword").should('have.value', "P@ssword")
      cy.wait(1000)
      cy.get('button').eq(7).click()
    })
  })
  
  describe('openprofilepage', () => {
    it('passes', () => {
      cy.get('button').eq(4).click()
      //cy.get('menuitem').eq(0).click()
      //cy.get('menuitem.profile').parent().click()
      cy.get('button').contains('Profile').click()
      cy.url().should('include', '/profile') // => true
    })
  })