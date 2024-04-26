package services_test

import (
	"context"
	"testing"

	"github.com/imrinry/assessment-tax/errs"
	"github.com/imrinry/assessment-tax/models"
	"github.com/imrinry/assessment-tax/repositories"
	"github.com/imrinry/assessment-tax/services" // Import the package containing the TaxCalculations function
	"github.com/stretchr/testify/assert"
)

func TestTaxCalculations(t *testing.T) {

	type testCase struct {
		income        float64
		wht           float64
		allowances    []models.Allowance
		expected      models.TaxResponse
		expectedLevel []models.TaxBracket
	}

	repo := repositories.NewTaxRepoMock()
	service := services.New(repo)

	testCases := []testCase{
		{income: 200000.0, wht: 0.0, allowances: []models.Allowance{}, expected: models.TaxResponse{Tax: 0, TaxRefund: 0}},
		{income: 1000000.0, wht: 50000.0, allowances: []models.Allowance{}, expected: models.TaxResponse{Tax: 51000, TaxRefund: 0}},
		{income: 1100000.0, wht: 0.0, allowances: []models.Allowance{}, expected: models.TaxResponse{Tax: 118000, TaxRefund: 0}},
		{income: 10000000.0, wht: 0.0, allowances: []models.Allowance{}, expected: models.TaxResponse{Tax: 3089000, TaxRefund: 0}},
		{income: 210000.0, wht: 2000.0, allowances: []models.Allowance{}, expected: models.TaxResponse{Tax: 0, TaxRefund: 2000}},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			response, err := service.TaxCalculations(context.Background(), tc.income, tc.wht, tc.allowances)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected.Tax, response.Tax)
			assert.Equal(t, tc.expected.TaxRefund, response.TaxRefund)
		})

	}

	// test error
	t.Run("invalid allowance type", func(t *testing.T) {
		_, err := service.TaxCalculations(context.Background(), 100000.0, 0.0, []models.Allowance{{AllowanceType: "xyz", Amount: 0}})
		assert.ErrorIs(t, err, errs.NewValidationError("invalid allowance type"))
	})

}

// returns error for invalid WHT value greater than income
func TestInvalidWHTValue(t *testing.T) {
	// Initialize test data
	income := 500000.0
	wht := 500001.0
	allowances := []models.Allowance{
		{AllowanceType: models.KReceiptType, Amount: 20000},
		{AllowanceType: models.DonationType, Amount: 5000},
	}
	repo := repositories.NewTaxRepoMock()
	s := services.New(repo)

	// Invoke code under test
	_, err := s.TaxCalculations(context.Background(), income, wht, allowances)

	// Assert the expected error
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid allowance wht value must be less than income: 500001.000000")
}

// Deducts valid donation allowance from taxable income
func TestDeductValidDonationAllowance(t *testing.T) {

	type testCase struct {
		taxableIncome float64
		allowances    []models.Allowance
		expected      float64
		errors        error
	}

	testCases := []testCase{
		{1000000.0, []models.Allowance{{AllowanceType: models.DonationType, Amount: 150000}}, 900000.0, nil},
		{1000000.0, []models.Allowance{{AllowanceType: models.DonationType, Amount: 50000}}, 950000.0, nil},
		{1000000.0, []models.Allowance{{AllowanceType: models.DonationType, Amount: 1}}, 999999.0, nil},
		{1000000.0, []models.Allowance{{AllowanceType: models.KReceiptType, Amount: 49000}}, 951000.0, nil},
		{1000000.0, []models.Allowance{{AllowanceType: models.KReceiptType, Amount: 55000}}, 950000.0, nil},
	}

	for _, tc := range testCases {
		t.Run("TestDeductValidDonationAllowance", func(t *testing.T) {
			services.DeductTaxAllowances(tc.allowances, &tc.taxableIncome)
			assert.NoError(t, tc.errors)
			assert.Equal(t, tc.expected, tc.taxableIncome)
		})
		// }
		// for _, tc := range testCases {
		// 	t.Run("error test", func(t *testing.T) {
		// 		err := services.DeductTaxAllowances(tc.allowances, &tc.taxableIncome)
		// 		assert.Error(t, err)

		// })
	}

}
