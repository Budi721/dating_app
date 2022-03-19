package http_req

type LimitReq struct {
	PageNo   int `query:"page_no"`
	PageSize int `query:"page_size"`
}
