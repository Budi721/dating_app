package v1

import (
	"errors"
	"github.com/Budi721/dating_app/delivery/http_req"
	"github.com/Budi721/dating_app/delivery/http_resp"
	"github.com/Budi721/dating_app/entity"
	"github.com/Budi721/dating_app/usecase"
	"github.com/Budi721/dating_app/utils/base64"
	"github.com/Budi721/dating_app/utils/logger"
	"github.com/Budi721/dating_app/utils/read"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type MemberApi struct {
	BaseApi
	memberRegistrationUseCase usecase.MemberRegistrationUseCase
	memberProfileUseCase      usecase.MemberProfileUseCase
	memberPreferenceUseCase   usecase.MemberPreferenceUseCase
	path                      string
}

func (m *MemberApi) memberActivation() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(http_req.MemberIdReq)
		jsonResponse, err := m.ParseRequestQuery(ctx, newReq)
		logger.Log.Debug().Msg(newReq.MemberId)
		if err != nil {
			logger.Log.Error().Msg("Parse error")
			return jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", "Error parse request body"))
		}
		err = m.memberRegistrationUseCase.NewActivation(newReq.MemberId)
		if err != nil {
			logger.Log.Error().Msg("Activation failed")
			return jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", "Activation failed"))
		}

		return jsonResponse.SendData(http_resp.NewResponseMessage("", "Activation success", newReq.MemberId))
	}
}

func (m *MemberApi) memberRegistration() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(http_req.MemberRegistrationReq)
		jsonResponse, err := m.ParseRequestBody(ctx, newReq)
		if err != nil {
			logger.Log.Error().Msg("Parse error")
			return jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", "Error parse request body"))
		}
		newMember := newReq.ToUserAccessMemberRegistration()
		err = m.memberRegistrationUseCase.NewRegistration(newMember)
		if err != nil {
			logger.Log.Error().Msg("Registration failed")
			return jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", "Registration failed"))
		}

		return jsonResponse.SendData(http_resp.NewResponseMessage("", "Registration success", entity.MemberUserAccess{MemberId: newMember.MemberId}))
	}
}

func (m *MemberApi) updateProfile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(http_req.MemberProfileReq)
		jsonResponse, _ := m.ParseRequestBody(ctx, newReq)
		logger.Log.Debug().Msg(newReq.String())
		memberInfo := newReq.ToMember()
		member, err := m.memberProfileUseCase.UpdateProfile(memberInfo)
		if err != nil {
			logger.Log.Error().Msg("Update profile failed")
			return jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", "Update profile failed"))
		}

		return jsonResponse.SendData(http_resp.NewResponseMessage("", "Update profile success", member.PersonalInfo.MemberId))

	}
}

func (m *MemberApi) createPreference() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(http_req.MemberPreferenceReq)
		jsonResponse, _ := m.ParseRequestBody(ctx, newReq)
		logger.Log.Debug().Msg(newReq.String())

		memberPref := newReq.ToMemberPreference()
		memberInterest := newReq.ToMemberInterest()

		err := m.memberPreferenceUseCase.CreatePreference(memberPref, memberInterest)
		if err != nil {
			logger.Log.Error().Msg("create preference failed")
			return jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", "Create preference profile failed"))
		}

		return jsonResponse.SendData(http_resp.NewResponseMessage("", "Create preference success", memberPref.PreferenceId))
	}
}

func (m *MemberApi) uploadProfile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(http_req.MemberIdReq)
		jsonResp, _ := m.ParseRequestQuery(ctx, newReq)
		logger.Log.Debug().Msg(newReq.MemberId)

		file, err := ctx.FormFile("profile")
		if err != nil {
			logger.Log.Error().Msg("unable upload profile")
			return jsonResp.SendError(http_resp.NewErrorMessage(http.StatusInternalServerError, "", "upload profile failed"))
		}
		recentPhoto := m.path + newReq.MemberId + "_" + file.Filename
		err = m.memberProfileUseCase.UpdateRecentPhoto(recentPhoto, newReq.MemberId)
		if err != nil {
			logger.Log.Error().Msg("upload photo failed")
			return jsonResp.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", "Upload photo failed"))
		}

		err = ctx.SaveFile(file, recentPhoto)
		if err != nil {
			logger.Log.Error().Msg(err.Error())
			return jsonResp.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", err.Error()))
		}

		buf := read.ReadFile(recentPhoto)
		imageBase64 := base64.ToBase64(buf)

		return jsonResp.SendData(http_resp.NewResponseMessage("", "Upload photo success", imageBase64))
	}
}

func (m *MemberApi) getProfile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(http_req.MemberIdReq)
		jsonResp, _ := m.ParseRequestQuery(ctx, newReq)
		logger.Log.Debug().Msg(newReq.MemberId)

		profile, err := m.memberProfileUseCase.GetProfile(newReq.MemberId)
		if err != nil {
			logger.Log.Error().Msg(err.Error())
			return jsonResp.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", err.Error()))
		}

		return jsonResp.SendData(http_resp.NewResponseMessage("", "Get photo success", profile))
	}
}

func NewMemberApi(
	rg fiber.Router,
	memberRegistrationUseCase usecase.MemberRegistrationUseCase,
	memberProfileUseCase usecase.MemberProfileUseCase,
	memberPreferenceUseCase usecase.MemberPreferenceUseCase,
	staticPath string,
) error {
	if memberRegistrationUseCase == nil {
		return errors.New("empty use case for member registration")
	}
	if memberPreferenceUseCase == nil {
		return errors.New("empty use case for member preference")
	}
	if memberProfileUseCase == nil {
		return errors.New("empty use case for profile info")
	}
	memberApi := &MemberApi{
		memberRegistrationUseCase: memberRegistrationUseCase,
		memberProfileUseCase:      memberProfileUseCase,
		memberPreferenceUseCase:   memberPreferenceUseCase,
		path:                      staticPath,
	}

	memberGroup := rg.Group("/member")
	memberGroup.Post("/registration", memberApi.memberRegistration())
	memberGroup.Put("/activation", memberApi.memberActivation())

	profileGroup := memberGroup.Group("/profile")
	profileGroup.Put("", memberApi.updateProfile())
	profileGroup.Get("", memberApi.getProfile())
	profileGroup.Post("", memberApi.uploadProfile())

	preferenceGroup := memberGroup.Group("/preference")
	preferenceGroup.Post("", memberApi.createPreference())
	return nil
}
