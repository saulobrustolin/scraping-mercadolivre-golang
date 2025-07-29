package models

type Product struct {
	Title           string  `json:"title"`
	QuantityReviews int64   `json:"quantity_reviews"`
	Stars           float64 `json:"stars"`
	Price           float64 `json:"price"`
	AnchorPrice     float64 `json:"anchor_price"`
	URL             string  `json:"url"`
	Picture         string  `json:"picture"`
}
