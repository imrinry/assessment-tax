package services_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/imrinry/assessment-tax/models"
// 	"github.com/imrinry/assessment-tax/repositories"
// 	"github.com/imrinry/assessment-tax/services" // Import the package containing the TaxCalculations function
// 	"github.com/stretchr/testify/assert"
// )

// func TestValidInput(t *testing.T) {
// 	// Initialize test data
// 	income := 1000000.0
// 	wht := 50000.0
// 	allowances := []models.Allowance{}

// 	repo := repositories.NewTaxRepoMock()
// 	s := services.New(repo)

// 	response, err := s.TaxCalculations(context.Background(), income, wht, allowances)
// 	// Assert the expected tax response
// 	assert.NoError(t, err)
// 	assert.Equal(t, 51000.0, response.Tax)
// 	assert.Equal(t, 0.0, response.TaxRefund)
// 	assert.Equal(t, []models.TaxBracket{
// 		{"0-150,000", 0},
// 		{"150,001-500,000", 35000},
// 		{"500,001-1,000,000", 66000},
// 		{"1,000,001-2,000,000", 0},
// 		{"2,000,001 ขึ้นไป", 0},
// 	}, response.TaxLevel)
// }

// // returns error for invalid WHT value greater than income
// func TestInvalidWHTValue(t *testing.T) {
// 	// Initialize test data
// 	income := 500000.0
// 	wht := 500001.0
// 	allowances := []models.Allowance{
// 		{AllowanceType: models.KReceiptType, Amount: 20000},
// 		{AllowanceType: models.DonationType, Amount: 5000},
// 	}
// 	repo := repositories.NewTaxRepoMock()
// 	s := services.New(repo)

// 	// Invoke code under test
// 	_, err := s.TaxCalculations(context.Background(), income, wht, allowances)

// 	// Assert the expected error
// 	assert.Error(t, err)
// 	assert.EqualError(t, err, "invalid allowance wht value must be less than income: 500001.000000")
// }
