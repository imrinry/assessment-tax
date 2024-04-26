package utils

// import "fmt"

// func CalculateTax3(income float64) float64 {
// 	type TaxLevel struct {
// 		Level  string  `json:"level"`
// 		Amount float64 `json:"amount"`
// 	}
// 	levels := []TaxLevel{
// 		{"0-150,000", 0},
// 	}

// 	var totalTax float64
// 	var taxLevels []TaxLevel

// 	personalDeduct := 60000.0

// 	taxableIncome := income - personalDeduct

// 	taxLevels = append(taxLevels, TaxLevel{})
// 	if taxableIncome > 150000 {
// 		nextLevel := min(taxableIncome, 500000) - 150000
// 		taxLevels = append(taxLevels, TaxLevel{Amount: nextLevel * 0.10, Level: "150,001-500,000"})
// 	}
// 	if taxableIncome > 500000 {
// 		nextLevel := min(taxableIncome, 1000000) - 500000
// 		taxLevels = append(taxLevels, TaxLevel{Amount: nextLevel * 0.15, Level: "500,001-1,000,000"})
// 	}
// 	if taxableIncome > 1000000 {
// 		nextLevel := min(taxableIncome, 2000000) - 1000000
// 		taxLevels = append(taxLevels, TaxLevel{Amount: nextLevel * 0.20, Level: "1,000,001-2,000,000"})
// 	}
// 	if taxableIncome > 2000000 {
// 		nextLevel := taxableIncome - 2000000
// 		taxLevels = append(taxLevels, TaxLevel{Amount: nextLevel * 0.35}, TaxLevel{Level: "2,000,001 ขึ้นไป"})
// 	}

// 	for _, level := range taxLevels {
// 		totalTax += level.Amount
// 		fmt.Println("Tax level:", level.Amount)
// 	}

// 	return totalTax
// }

// func TaxCalulate(income float64) float64 {
// 	income -= 60000

// 	switch {
// 	case income <= 150000:
// 		return 0
// 	case income <= 500000:
// 		return (income - 150000) * 0.1
// 	case income <= 1000000:
// 		return 35000 + (income-500000)*0.15
// 	case income <= 2000000:
// 		return 110000 + (income-1000000)*0.2
// 	default:
// 		return 310000 + (income-2000000)*0.35
// 	}
// }
