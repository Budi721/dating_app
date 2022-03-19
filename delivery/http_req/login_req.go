package http_req

import "fmt"

type LoginReq struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (p *LoginReq) String() string {
	return fmt.Sprintf("LoginReq => username : %s", p.UserName)
}
