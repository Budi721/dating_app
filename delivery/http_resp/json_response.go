package http_resp

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type JsonResponse struct {
	ctx *fiber.Ctx
}

func (j *JsonResponse) SendData(message *ResponseMessage) error {
	return j.ctx.Status(http.StatusOK).JSON(message)
}

func (j *JsonResponse) SendError(errorMessage *ErrorMessage) error {
	return j.ctx.Status(errorMessage.HttpCode).JSON(errorMessage)
}

func NewJsonResponse(ctx *fiber.Ctx) AppHttpResponse {
	return &JsonResponse{
		ctx: ctx,
	}
}
