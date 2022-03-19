package http_req

import (
	"github.com/Budi721/dating_app/entity"
	"github.com/google/uuid"
	"time"
)

type MemberRegistrationReq struct {
	Email    string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (m *MemberRegistrationReq) ToUserAccessMemberRegistration() *entity.MemberUserAccess {
	joinDate := time.Now().Local()
	return &entity.MemberUserAccess{
		Username:           m.Email,
		Password:           m.Password,
		MemberId:           uuid.New().String(),
		JoinDate:           joinDate,
		VerificationStatus: "N",
	}
}
