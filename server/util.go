package server

import (
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

//NewCoockieSession NewCoockieSession
func NewCoockieSession() (sess echo.MiddlewareFunc) {
	cookieStore := session.NewCookieStore([]byte("gateway-session-cookie"))
	sess = session.Sessions("gateway-session", cookieStore)
	return
}
