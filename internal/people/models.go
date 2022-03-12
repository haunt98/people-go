package people

type Person struct {
	ID       string  `json:"id"`
	Name     *string `json:"name"`
	Birthday *string `json:"birthday"`
	Phone    *string `json:"phone"`

	// Chung minh nhan dan
	CMND *string `json:"cmnd"`
	// Ma so thue
	MST *string `json:"mst"`
	// Bao hiem xa hoi
	BHXH       *string `json:"bhxh"`
	Address    *string `json:"address"`
	University *string `json:"university"`

	// Company
	VNG *string `json:"vng"`

	// Social networks
	Facebook  *string `json:"facebook"`
	Instagram *string `json:"instagram"`
	Tiktok    *string `json:"tiktok"`
}
