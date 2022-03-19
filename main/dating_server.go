package main

import (
	v1 "github.com/Budi721/dating_app/delivery/v1"
	"github.com/Budi721/dating_app/manager"
	"github.com/Budi721/dating_app/utils/logger"
	"github.com/gofiber/fiber/v2"
)

type DatingServer interface {
	Run()
}

type datingServer struct {
	infra      manager.InfraManager
	useCase    manager.UseCaseManager
	router     *fiber.App
	apiVersion string
}

func (d datingServer) handlers() {
	apiMainGroup := d.infra.Config().GetString("datingapp.api.group")

	api := d.router.Group(apiMainGroup)
	apiVersionGroup := api.Group(d.apiVersion)

	switch d.apiVersion {
	case "/v1":
		d.v1(apiVersionGroup)
	default:
		logger.Log.Fatal().Msg("Undefined api version")
	}
}

func (d datingServer) v1(rg fiber.Router) {
	staticPath := d.infra.Config().GetString("datingapp.static.uploads")

	err := v1.NewMemberApi(
		rg,
		d.useCase.MemberRegistrationUseCase(),
		d.useCase.MemberProfileUseCase(),
		d.useCase.MemberPreferenceUseCase(),
		staticPath,
	)

	if err != nil {
		logger.Log.Fatal().Msg("member registration failed to start")
	}

	err = v1.NewAuthApi(rg, d.useCase.AuthUseCase())
	if err != nil {
		logger.Log.Fatal().Err(err).Msg(err.Error())
	}

	err = v1.NewPartnerFinderApi(rg, d.useCase.PartnerFinderUseCase())
	if err != nil {
		logger.Log.Fatal().Err(err).Msg(err.Error())
	}
}

func (d datingServer) Run() {
	apiUrl := d.infra.Config().GetString("datingapp.api.url")
	d.handlers()
	logger.Log.Info().Msg("Server is running")
	err := d.router.Listen(apiUrl)
	if err != nil {
		logger.Log.Error().Msg(err.Error())
	}
}

func NewDatingServer() DatingServer {
	infra := manager.NewInfra()
	repo := manager.NewRepositoryManager(infra)
	useCase := manager.NewUseCaseManager(repo)
	router := fiber.New(fiber.Config{
		AppName:       infra.Config().GetString("datingapp.name"),
		CaseSensitive: true,
	})
	apiVersion := infra.Config().GetString("datingapp.api.version")

	return &datingServer{
		infra:      infra,
		useCase:    useCase,
		router:     router,
		apiVersion: apiVersion,
	}
}
