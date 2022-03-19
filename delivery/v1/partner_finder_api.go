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

type PartnerFinderApi struct {
	BaseApi
	PartnerFinderUseCase usecase.PartnerFinderUseCase
}

func (a *PartnerFinderApi) findPartner() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var findPartner struct {
			http_req.MemberIdReq
			http_req.LimitReq
		}

		jsonResponse, _ := a.ParseRequestQuery(ctx, findPartner)
		partnerInfo, err := a.PartnerFinderUseCase.ViewPartner(findPartner.MemberId, findPartner.PageNo, findPartner.PageSize)
		if err != nil {
			logger.Log.Error().Msg("find partner failed")
			return jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", "Find partner failed"))
		}

		return jsonResponse.SendData(http_resp.NewResponseMessage("", "Find partner success", partnerInfo))

	}
}

func (a *PartnerFinderApi) matchPartner() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(http_req.MemberPartnerReq)
		jsonResponse, _ := a.ParseRequestBody(ctx, newReq)
		logger.Log.Debug().Msgf(newReq.MemberId)

		err := a.PartnerFinderUseCase.MatchPartner(newReq.MemberId, newReq.PartnerId)
		if err != nil {
			logger.Log.Error().Msg("match partner failed")
			return jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusInternalServerError, "", "match partner failed"))
		}

		return jsonResponse.SendData(http_resp.NewResponseMessage("", "Match partner success", true))
	}
}

func (a *PartnerFinderApi) listMatchPartner() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(http_req.MemberIdReq)
		jsonResponse, _ := a.ParseRequestQuery(ctx, newReq)
		logger.Log.Debug().Msgf("list partner service id %s", newReq.MemberId)
		partners, err := a.PartnerFinderUseCase.ListPartner(newReq.MemberId)
		if err != nil {
			logger.Log.Error().Msg("list partner failed")
			return jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusInternalServerError, "", "list partner failed"))
		}

		return jsonResponse.SendData(http_resp.NewResponseMessage("", "List partner success", partners))
	}
}

func NewPartnerFinderApi(rg fiber.Router, partnerFinderUseCase usecase.PartnerFinderUseCase) error {
	if partnerFinderUseCase == nil {
		return errors.New("use case partner finder not available")
	}
	partnerFinderApi := &PartnerFinderApi{
		PartnerFinderUseCase: partnerFinderUseCase,
	}
	partnerGroup := rg.Group("/partner")
	partnerGroup.Get("/view", partnerFinderApi.findPartner())
	partnerGroup.Post("/match", partnerFinderApi.matchPartner())
	partnerGroup.Get("/list", partnerFinderApi.listMatchPartner())
	return nil
}
