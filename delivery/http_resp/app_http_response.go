package http_resp

type AppHttpResponse interface {
	SendData(message *ResponseMessage) error
	SendError(errorMessage *ErrorMessage) error
}
