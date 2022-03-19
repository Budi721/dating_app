package manager

import "github.com/Budi721/dating_app/usecase"

type UseCaseManager interface {
	MemberRegistrationUseCase() usecase.MemberRegistrationUseCase
	AuthUseCase() usecase.AuthenticationUseCase
	MemberProfileUseCase() usecase.MemberProfileUseCase
	MemberPreferenceUseCase() usecase.MemberPreferenceUseCase
	PartnerFinderUseCase() usecase.PartnerFinderUseCase
}

type useCaseManager struct {
	repo RepositoryManager
}

func (u *useCaseManager) PartnerFinderUseCase() usecase.PartnerFinderUseCase {
	return usecase.NewPartnerFinderUseCase(u.repo.PartnerRepo(), u.repo.MemberPreferenceRepo())
}

func (u *useCaseManager) MemberPreferenceUseCase() usecase.MemberPreferenceUseCase {
	return usecase.NewMemberPreferenceUseCase(u.repo.MemberPreferenceRepo())
}

func (u *useCaseManager) MemberProfileUseCase() usecase.MemberProfileUseCase {
	return usecase.NewMemberProfileUseCase(u.repo.MemberInfoRepo())
}

func (u *useCaseManager) AuthUseCase() usecase.AuthenticationUseCase {
	return usecase.NewAuthenticationUseCase(u.repo.MemberAccessRepo())
}

func (u *useCaseManager) MemberRegistrationUseCase() usecase.MemberRegistrationUseCase {
	return usecase.NewMemberRegistrationUseCase(u.repo.MemberAccessRepo())
}

func NewUseCaseManager(repo RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repo: repo,
	}
}
