package v1

import (
	"github.com/Budi721/dating_app/delivery/http_resp"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type BaseApi struct {
}

func (b *BaseApi) ParseRequestBody(ctx *fiber.Ctx, requestBody interface{}) (http_resp.AppHttpResponse, error) {
	jsonResponse := http_resp.NewJsonResponse(ctx)
	if err := ctx.BodyParser(requestBody); err != nil {
		return nil, jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", "Request cannot be parsed"))
	}
	return jsonResponse, nil
}

func (b *BaseApi) ParseRequestQuery(ctx *fiber.Ctx, requestBody interface{}) (http_resp.AppHttpResponse, error) {
	jsonResponse := http_resp.NewJsonResponse(ctx)
	if err := ctx.QueryParser(requestBody); err != nil {
		return nil, jsonResponse.SendError(http_resp.NewErrorMessage(http.StatusBadRequest, "", "Query string cannot be parsed"))
	}
	return jsonResponse, nil
}

func (b *BaseApi) ParseRequest(ctx *fiber.Ctx) (http_resp.AppHttpResponse, error) {
	jsonResponse := http_resp.NewJsonResponse(ctx)
	return jsonResponse, nil
}
