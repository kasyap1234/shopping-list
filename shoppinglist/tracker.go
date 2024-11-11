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
type getItemResponse struct {
	Item Item `json:"item"`
}

//encore:api public method=GET path=/items/:id
func (s *Service) GetItem(ctx context.Context, id int) (*getItemResponse, error) {
	var item Item
	err := s.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &getItemResponse{Item: item}, nil
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

//encore:api public method=POST path=/items
func (s *Service) CreateItem(ctx context.Context, item Item) (bool, error) {
	err := s.db.Create(&item).Error
	if err != nil {
		return false, err
	}
	return true, nil

}

//encore:api public method=PUT path=/items/:id
func (s *Service) UpdateItem(ctx context.Context, id int, item Item) (*getItemResponse, error) {
	err := s.db.Model(&item).Where("id = ?", id).Updates(item).Error
	if err != nil {
		return nil, err
	}
	return &getItemResponse{Item: item}, nil
}
