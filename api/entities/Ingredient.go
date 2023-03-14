package entities

type Ingredient struct {
	Id     int     `json:"id"`
	Number float32 `json:"number"`
	Item   string  `json:"item"`
}
