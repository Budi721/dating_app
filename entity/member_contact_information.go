package entity

type MemberContactInformation struct {
	MemberContactId   string `json:"member_contact_id,omitempty"`
	MemberId          string `json:"member_id,omitempty"`
	MobilePhoneNumber string `json:"mobile_phone_number,omitempty"`
	InstagramId       string `json:"instagram_id,omitempty"`
	TwitterId         string `json:"twitter_id,omitempty"`
	Email             string `json:"email,omitempty"`
}
