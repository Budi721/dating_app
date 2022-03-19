package http_req

type MemberPartnerReq struct {
	MemberId  string `json:"member_id,omitempty"`
	PartnerId string `json:"partner_id,omitempty"`
}
