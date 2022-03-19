package entity

import "time"

type MemberPersonalInformation struct {
	PersonalInformationId string    `json:"personal_information_id,omitempty"`
	MemberId              string    `json:"member_id,omitempty"`
	Name                  string    `json:"name,omitempty"`
	Bod                   time.Time `json:"bod"`
	Gender                string    `json:"gender,omitempty"`
	RecentPhotoPath       string    `json:"recent_photo_path,omitempty"`
	SelfDescription       string    `json:"self_description,omitempty"`
}
