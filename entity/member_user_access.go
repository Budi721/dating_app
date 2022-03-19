package entity

import "time"

type MemberUserAccess struct {
	Username           string    `json:"username,omitempty"`
	Password           string    `json:"password,omitempty"`
	MemberId           string    `json:"member_id,omitempty"`
	JoinDate           time.Time `json:"join_date"`
	VerificationStatus string    `json:"verification_status,omitempty"`
}
