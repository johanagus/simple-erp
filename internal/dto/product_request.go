package dto

type ProductRequest struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Type     string  `json:"type"`
	Brand    string  `json:"brand"`
	Price    float64 `json:"price"`
	Barcode  string  `json:"barcode"`
	Unit     string  `json:"unit"`
	IsSN     bool    `json:"is_sn"`
}
