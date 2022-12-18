package people

type Person struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Birthday   string `json:"birthday,omitempty"`
	Phone      string `json:"phone,omitempty"`
	University string `json:"university,omitempty"`

	// Vietnam
	// Chung minh nhan dan
	VNCMND string `json:"vn_cmnd,omitempty"`
	VNCCCD string `json:"vn_cccd,omitempty"`
	// Ma so thue
	VNMST string `json:"vn_mst,omitempty"`
	// Bao hiem xa hoi
	VNBHXH string `json:"vn_bhxh,omitempty"`

	// Company
	CompanyVNG string `json:"company_vng,omitempty"`

	// Social networks
	SocialFacebook  string `json:"social_facebook,omitempty"`
	SocialInstagram string `json:"social_instagram,omitempty"`
	SocialTiktok    string `json:"social_tiktok,omitempty"`
}

type WrapPeople struct {
	People []*Person `json:"people"`
}
