package entity

type MemberAddressInformation struct {
	MemberAddressId string `json:"member_address_id,omitempty"`
	MemberId        string `json:"member_id,omitempty"`
	Address         string `json:"address,omitempty"`
	City            string `json:"city,omitempty"`
	PostalCode      string `json:"postal_code,omitempty"`
}
