package shoppinglist

import "context"

type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Bought   bool    `json:"bought"`
}
type getItemsResponse struct {
	Items []Item `json:"items"`
}

//encore:api public method=GET path=/items
func (s *Service) GetItems(context.Context) (*getItemsResponse, error) {
	var items []Item
	err := s.db.Find(&items).Error
	if err != nil {
		return nil, err
	}
	return &getItemsResponse{Items: items}, nil

}
