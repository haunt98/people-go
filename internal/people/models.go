package people

type Person struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday,omitempty"`
	Phone    string `json:"phone,omitempty"`

	// Chung minh nhan dan
	CMND string `json:"cmnd,omitempty"`
	// Ma so thue
	MST string `json:"mst,omitempty"`
	// Bao hiem xa hoi
	BHXH       string `json:"bhxh,omitempty"`
	University string `json:"university,omitempty"`

	// Company
	VNG string `json:"vng,omitempty"`

	// Social networks
	Facebook  string `json:"facebook,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Tiktok    string `json:"tiktok,omitempty"`
}

type WrapPeople struct {
	People []Person `json:"people"`
}
