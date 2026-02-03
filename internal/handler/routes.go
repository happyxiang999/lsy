package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"lsy/internal/svc"
)

func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {
	server.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/healthz",
			Handler: healthHandler(),
		},
		{
			Method:  http.MethodGet,
			Path:    "/readyz",
			Handler: readyHandler(),
		},
		{
			Method:  http.MethodGet,
			Path:    "/items",
			Handler: listItemsHandler(ctx),
		},
		{
			Method:  http.MethodPost,
			Path:    "/items",
			Handler: createItemHandler(ctx),
		},
		{
			Method:  http.MethodGet,
			Path:    "/items/:id",
			Handler: getItemHandler(ctx),
		},
		{
			Method:  http.MethodPut,
			Path:    "/items/:id",
			Handler: updateItemHandler(ctx),
		},
		{
			Method:  http.MethodDelete,
			Path:    "/items/:id",
			Handler: deleteItemHandler(ctx),
		},
	})
}
