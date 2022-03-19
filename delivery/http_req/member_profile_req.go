package http_req

import (
	"fmt"
	"github.com/Budi721/dating_app/entity"
	"github.com/Budi721/dating_app/utils/date"
)

type MemberProfileReq struct {
	MemberId        string `json:"member_id,omitempty"`
	Name            string `json:"name,omitempty"`
	Bod             string `json:"bod,omitempty"`
	Gender          string `json:"gender,omitempty"`
	SelfDescription string `json:"self_description,omitempty"`
	Instagram       string `json:"instagram,omitempty"`
	Twitter         string `json:"twitter,omitempty"`
	MobilePhone     string `json:"mobile_phone,omitempty"`
	Address         string `json:"address,omitempty"`
	City            string `json:"city,omitempty"`
	PostalCode      string `json:"postal_code,omitempty"`
}

func (m *MemberProfileReq) String() string {
	return fmt.Sprintf("LoginReq => member id: %s", m.MemberId)
}

func (m *MemberProfileReq) ToMember() *entity.Member {
	return &entity.Member{
		PersonalInfo: entity.MemberPersonalInformation{
			MemberId:        m.MemberId,
			Name:            m.Name,
			Bod:             date.StringToDate(m.Bod),
			Gender:          m.Gender,
			SelfDescription: m.SelfDescription,
		},
		AddressInfo: entity.MemberAddressInformation{
			MemberId:   m.MemberId,
			Address:    m.Address,
			City:       m.City,
			PostalCode: m.PostalCode,
		},
		ContactInfo: entity.MemberContactInformation{
			MemberId:          m.MemberId,
			MobilePhoneNumber: m.MobilePhone,
			InstagramId:       m.Instagram,
			TwitterId:         m.Twitter,
		},
	}
}
