package usecase

import (
	"github.com/Budi721/dating_app/entity"
	"github.com/Budi721/dating_app/repository"
)

type MemberPreferenceUseCase interface {
	CreatePreference(memberPref *entity.MemberPreferences, memberInterest []*entity.MemberInterest) error
}

type memberPreferenceUseCase struct {
	repo repository.MemberPreferenceRepo
}

func (m *memberPreferenceUseCase) CreatePreference(memberPref *entity.MemberPreferences, memberInterest []*entity.MemberInterest) error {
	return m.repo.Create(memberPref, memberInterest)
}

func NewMemberPreferenceUseCase(repo repository.MemberPreferenceRepo) MemberPreferenceUseCase {
	return &memberPreferenceUseCase{
		repo: repo,
	}
}
