package context

import (
	"net/http"
)

var defaultUserId = "";
type Ctx struct {
	Request *http.Request
	Response http.ResponseWriter
	UserId string 
}