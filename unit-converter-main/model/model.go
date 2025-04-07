package model

type Request struct {
	Input    float64 `json:"input"`
	UnitForm string  `json:"unitForm"`
	UnitTo   string  `json:"unitTo"`
}

type Response struct {
	Result float64 `json:"result"`
	Unit   string  `json:"unit"`
}
