package models

type Store struct {
	ID          string `json:"_id"`
	StoreID     int    `json:"StoreID"`
	Name        string `json:"Name"`
	Type        string `json:"Type"`
	Description string `json:"Description"`
	Items       []Item `json:"Items"`
}

type Item struct {
	ItemID      int     `json:"ItemID"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	Available   bool    `json:"Available"`
	Price       float64 `json:"Price"`
}
