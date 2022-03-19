package http_resp

type ResponseMessage struct {
	Status      string      `json:"status,omitempty"`
	Code        string      `json:"code,omitempty"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

type ErrorDescription struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type ErrorMessage struct {
	HttpCode         int              `json:"http_code,omitempty"`
	ErrorDescription ErrorDescription `json:"error_description,omitempty"`
}

func NewResponseMessage(code string, description string, data interface{}) *ResponseMessage {
	return &ResponseMessage{
		Status:      "Success",
		Code:        code,
		Description: description,
		Data:        data,
	}
}

func NewErrorMessage(httpCode int, errCode string, message string) *ErrorMessage {
	return &ErrorMessage{
		HttpCode: httpCode,
		ErrorDescription: ErrorDescription{
			Code:        errCode,
			Description: message,
		},
	}
}
