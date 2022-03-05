package driver

type Book struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MediumPrice float32 `json:"medium_price"`
	Author      string  `json:"author"`
	ImageURL    string  `json:"img_url"`
}
