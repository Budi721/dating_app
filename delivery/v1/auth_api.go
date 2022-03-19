package v1

import (
	"errors"
	"github.com/Budi721/dating_app/delivery/http_req"
	"github.com/Budi721/dating_app/delivery/http_resp"
	"github.com/Budi721/dating_app/usecase"
	"github.com/Budi721/dating_app/utils/logger"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthApi struct {
	BaseApi
	authUseCase usecase.AuthenticationUseCase
}

func (a *AuthApi) userLogin() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(http_req.LoginReq)
		jsonResponse, _ := a.ParseRequestBody(ctx, newReq)
		logger.Log.Debug().Msg(newReq.String())
		member, err := a.authUseCase.Login(newReq)
		if err != nil {
			logger.Log.Error().Msg("failed to login user")
			return jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusUnauthorized, "", "unauthorized"))
		}
		return jsonResponse.SendData(http_resp.NewResponseMessage("", "Login success", member.MemberId))
	}
}

func NewAuthApi(rg fiber.Router, authUseCase usecase.AuthenticationUseCase) error {
	if authUseCase == nil {
		return errors.New("use case auth not found")
	}
	authApi := AuthApi{
		authUseCase: authUseCase,
	}
	memberGroup := rg.Group("/auth")
	memberGroup.Post("/login", authApi.userLogin())

	return nil
}
