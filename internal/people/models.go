package people

import (
	"fmt"
)

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
	SocialLinkedin  string `json:"social_linkedin,omitempty"`
}

func (p *Person) Pretty(prefix string) string {
	template := "%sID: %s\n"
	values := make([]any, 0, 20)
	values = append(values, prefix, p.ID)

	if p.Name != "" {
		template += "%sName: %s\n"
		values = append(values, prefix, p.Name)
	}

	if p.Birthday != "" {
		template += "%sBirthday: %s\n"
		values = append(values, prefix, p.Birthday)
	}

	if p.Phone != "" {
		template += "%sPhone: %s\n"
		values = append(values, prefix, p.Phone)
	}

	if p.University != "" {
		template += "%sUniversity: %s\n"
		values = append(values, prefix, p.University)
	}

	if p.VNCMND != "" {
		template += "%sVNCMND: %s\n"
		values = append(values, prefix, p.VNCMND)
	}

	if p.VNCCCD != "" {
		template += "%sVNCCCD: %s\n"
		values = append(values, prefix, p.VNCCCD)
	}

	if p.VNMST != "" {
		template += "%sVNMST: %s\n"
		values = append(values, prefix, p.VNMST)
	}

	if p.VNBHXH != "" {
		template += "%sVNBHXH: %s\n"
		values = append(values, prefix, p.VNBHXH)
	}

	if p.CompanyVNG != "" {
		template += "%sCompanyVNG: %s\n"
		values = append(values, prefix, p.CompanyVNG)
	}

	if p.SocialFacebook != "" {
		template += "%sSocialFacebook: %s\n"
		values = append(values, prefix, p.SocialFacebook)
	}

	if p.SocialInstagram != "" {
		template += "%sSocialInstagram: %s\n"
		values = append(values, prefix, p.SocialInstagram)
	}

	if p.SocialTiktok != "" {
		template += "%sSocialTiktok: %s"
		values = append(values, prefix, p.SocialTiktok)
	}

	if p.SocialLinkedin != "" {
		template += "%sSocialLinkedin: %s"
		values = append(values, prefix, p.SocialLinkedin)
	}

	return fmt.Sprintf(template, values...)
}

type WrapPeople struct {
	People []*Person `json:"people"`
}
