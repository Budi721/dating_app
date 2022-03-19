package http_req

import (
	"fmt"
	"github.com/Budi721/dating_app/entity"
)

type MemberPreferenceReq struct {
	MemberId         string           `json:"member_id,omitempty"`
	GenderInterest   string           `json:"gender_interest,omitempty"`
	DomicileInterest string           `json:"domicile_interest,omitempty"`
	StartAgeInterest int              `json:"start_age_interest,omitempty"`
	EndAgeInterest   int              `json:"end_age_interest,omitempty"`
	Interest         []MemberInterest `json:"interest,omitempty"`
}

type MemberInterest struct {
	InterestId string `json:"interest_id,omitempty"`
}

func (m *MemberPreferenceReq) String() string {
	return fmt.Sprintf("MemberPreferenceReq => MemberId : %v", m.MemberId)
}

func (m *MemberPreferenceReq) ToMemberPreference() *entity.MemberPreferences {
	return &entity.MemberPreferences{
		MemberId:           m.MemberId,
		LookingForGender:   m.GenderInterest,
		LookingForDomicile: m.DomicileInterest,
		LookingForStartAge: m.StartAgeInterest,
		LookingForEndAge:   m.EndAgeInterest,
	}
}

func (m *MemberPreferenceReq) ToMemberInterest() []*entity.MemberInterest {
	var interests []*entity.MemberInterest
	for _, interest := range interests {
		interests = append(interests, &entity.MemberInterest{
			InterestId: interest.InterestId,
			MemberId:   m.MemberId,
		})
	}

	return interests
}
