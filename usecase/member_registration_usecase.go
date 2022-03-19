package usecase

import (
	"github.com/Budi721/dating_app/entity"
	"github.com/Budi721/dating_app/repository"
)

type MemberRegistrationUseCase interface {
	NewRegistration(userAccess *entity.MemberUserAccess) error
	NewActivation(id string) error
}

type memberRegistrationUseCase struct {
	accessRepo repository.MemberAccessRepo
}

func (m *memberRegistrationUseCase) NewRegistration(userAccess *entity.MemberUserAccess) error {
	return m.accessRepo.Create(userAccess)
}

func (m *memberRegistrationUseCase) NewActivation(id string) error {
	return m.accessRepo.UpdateVerification(id)
}

func NewMemberRegistrationUseCase(repo repository.MemberAccessRepo) MemberRegistrationUseCase {
	return &memberRegistrationUseCase{accessRepo: repo}
}
