package entity

// 组件1：基金实体
type Fund struct {
	Code              string    `json:"code"`
	Name              string    `json:"name"`
	OriginalFeeRate   float64   `json:"originalFeeRate"`
	CurrentFeeRate    float64   `json:"currentFeeRate"`
	MinPurchaseAmount float64   `json:"minPurchaseAmount"`
	HoldingStocks     []string  `json:"holdingStocks"`
	HoldingBonds      []string  `json:"holdingBonds"`
	ReturnRates       []float64 `json:"returnRates"`
}
