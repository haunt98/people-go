package people

type Person struct {
	ID       string  `json:"id"`
	Name     *string `json:"name"`
	Birthday *string `json:"birthday"`
	Phone    *string `json:"phone"`

	// Vietnam
	CMND       *string `json:"cmnd"`
	MST        *string `json:"mst"`  // Ma so thue
	BHXH       *string `json:"bhxh"` // Bao hiem xa hoi
	Address    *string `json:"address"`
	University *string `json:"university"`

	// Company
	VNG *string `json:"vng"`

	// Social networks
	Facebook  *string `json:"facebook"`
	Instagram *string `json:"instagram"`
	Tiktok    *string `json:"tiktok"`
}
