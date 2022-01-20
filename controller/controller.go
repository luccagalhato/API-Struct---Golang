package controller

import (
	"net/http"
	"vendas/server"
)

//Controllers ...
func (c *Controller) Controllers() map[string]server.Handler {
	return map[string]server.Handler{
		"/cadastro_page": {
			Method: http.MethodGet,
			Fn:     nil,
		},
	}
}
