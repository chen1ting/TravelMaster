describe('empty spec', () => {
  it('passes', () => {
    cy.visit('https://172.21.148.164/')
    //cy.visit('http://localhost:3000/')
    cy.get('input').eq(0).type("Tom")
    cy.get('input').eq(1).type("P@ssword")
    cy.get('button').eq(0).click()
  })
})