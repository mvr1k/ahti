package api

import "net/http"

/*API interface tell that who ever is a API can be passed to the web server as a API handler*/

type API interface {
	GetRoutingList() []RouteInfo
	ModuleName() string
}

type RouteInfo struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func NewRouteInfo(path string, method string, handler func(writer http.ResponseWriter, request *http.Request)) RouteInfo {
	return RouteInfo{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}
