package models

type TaxRequest struct {
	TotalIncome float64     `json:"totalIncome"`
	WHT         float64     `json:"wht"`
	Allowances  []Allowance `json:"allowances"`
}

type Allowance struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

type TaxResponse struct {
	Tax      float64      `json:"tax"`
	TaxLevel []TaxBracket `json:"taxLevel,omitempty"`
}

type TaxBracket struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}

type AdminSettings struct {
	PersonalDeduction float64 `json:"personalDeduction"`
	KReceipt          float64 `json:"kReceipt"`
}
