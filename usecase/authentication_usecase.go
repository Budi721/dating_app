package usecase

import (
	"github.com/Budi721/dating_app/delivery/http_req"
	"github.com/Budi721/dating_app/entity"
	"github.com/Budi721/dating_app/repository"
)

type AuthenticationUseCase interface {
	Login(loginReq *http_req.LoginReq) (*entity.MemberUserAccess, error)
}

type authenticationUseCase struct {
	repo repository.MemberAccessRepo
}

func (a *authenticationUseCase) Login(loginReq *http_req.LoginReq) (*entity.MemberUserAccess, error) {
	return a.repo.FindByUsernamePasswordVerified(loginReq.UserName, loginReq.Password)
}

func NewAuthenticationUseCase(repo repository.MemberAccessRepo) AuthenticationUseCase {
	return &authenticationUseCase{
		repo: repo,
	}
}
