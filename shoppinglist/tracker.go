package shoppinglist

type Item struct {
	ID int  `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	Bought bool `json:"bought"`

}
