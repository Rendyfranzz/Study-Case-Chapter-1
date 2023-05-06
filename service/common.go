package service

import "net/http"

type CommonRequest struct {
	AuthToken string `json:"auth_token"`
}

type CommonResponse struct {
	StatusCode   int    `json:"code,omitempty"`
	Msg          string `json:"message,omitempty"`
	SetAuthToken string `json:"-"`
}

func (res *CommonResponse) SetMsg(statusCode int, msg string) {
	res.StatusCode = statusCode
	res.Msg = msg
}

func (res *CommonResponse) SetOK() {
	res.StatusCode = http.StatusOK
	res.Msg = "OK"
}
