describe('login', () => {
  it('passes', () => {
    cy.visit('https://172.21.148.164/')
    //cy.visit('http://localhost:3000/')
    cy.get('input').eq(0).type("Tom").should('have.value', 'Tom')
    cy.get('input').eq(1).type("P@ssword").should('have.value', "P@ssword")
    cy.get('button').eq(7).click()
  })
})

describe('openDiscoverPage', () => {
    it('passes', () => {
      cy.get('button').eq(0).click()
      cy.url().should('include', '/discover') // => true
    })
  })

describe('openActivityPage', () => {
    it('passes', () => {
          cy.contains('Food').click()
          cy.url().should('include', '/activity') // => true
    })
  })

describe('reviewCheck', () => {
    it('passes', () => {
        cy.get('input').eq(0).type('Review Title').should('have.value','Review Title')
        cy.get('textarea').eq(0).type('Review Description').should('have.value','Review Description')
        cy.get('[data-cy]="starratings"').eq(4).click({force: true})
        cy.get('button').contains('Submit').click().should('have.value', 'Successfully uploaded your review.')
    })
  })