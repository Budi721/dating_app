package manager

import "github.com/Budi721/dating_app/repository"

type RepositoryManager interface {
	MemberAccessRepo() repository.MemberAccessRepo
	MemberInfoRepo() repository.MemberInfoRepo
	MemberPreferenceRepo() repository.MemberPreferenceRepo
	PartnerRepo() repository.PartnerRepo
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) PartnerRepo() repository.PartnerRepo {
	return repository.NewPartnerRepo(r.infraManager.SqlDatabase())
}

func (r *repositoryManager) MemberPreferenceRepo() repository.MemberPreferenceRepo {
	return repository.NewMemberPreferenceRepo(r.infraManager.SqlDatabase())
}

func (r *repositoryManager) MemberInfoRepo() repository.MemberInfoRepo {
	return repository.NewMemberInfoRepo(r.infraManager.SqlDatabase())
}

func (r *repositoryManager) MemberAccessRepo() repository.MemberAccessRepo {
	return repository.NewMemberAccessRepo(r.infraManager.SqlDatabase())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{infraManager: infra}
}
