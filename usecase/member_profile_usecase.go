package usecase

import (
	"github.com/Budi721/dating_app/entity"
	"github.com/Budi721/dating_app/repository"
)

type MemberProfileUseCase interface {
	UpdateProfile(member *entity.Member) (*entity.Member, error)
	GetProfile(id string) (*entity.Member, error)
	UpdateRecentPhoto(path string, id string) error
}

type memberProfileUseCase struct {
	repo repository.MemberInfoRepo
}

func (m *memberProfileUseCase) GetProfile(id string) (*entity.Member, error) {
	return m.repo.FindById(id)
}

func (m *memberProfileUseCase) UpdateRecentPhoto(path string, id string) error {
	return m.repo.Update(path, id)
}

func (m *memberProfileUseCase) UpdateProfile(member *entity.Member) (*entity.Member, error) {
	return m.repo.Create(member)
}

func NewMemberProfileUseCase(repo repository.MemberInfoRepo) MemberProfileUseCase {
	return &memberProfileUseCase{
		repo: repo,
	}
}
