package handler

import (
	"net/http"
	"strings"

	"lsy/internal/svc"
	"lsy/internal/types"
)

func listItemsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		items := ctx.Store.List()
		writeJSON(w, http.StatusOK, types.ListItemsResp{Items: items})
	}
}

func createItemHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateItemReq
		if err := decodeJSON(w, r, &req); err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		name := strings.TrimSpace(req.Name)
		if name == "" {
			writeError(w, http.StatusBadRequest, "name is required")
			return
		}

		item := ctx.Store.Create(name, strings.TrimSpace(req.Description))
		writeJSON(w, http.StatusCreated, item)
	}
}

func getItemHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := itemIDFromPath(r.URL.Path)
		if !ok {
			writeError(w, http.StatusBadRequest, "invalid item id")
			return
		}

		item, found := ctx.Store.Get(id)
		if !found {
			writeError(w, http.StatusNotFound, "item not found")
			return
		}

		writeJSON(w, http.StatusOK, item)
	}
}

func updateItemHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := itemIDFromPath(r.URL.Path)
		if !ok {
			writeError(w, http.StatusBadRequest, "invalid item id")
			return
		}

		var req types.UpdateItemReq
		if err := decodeJSON(w, r, &req); err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		if req.Name == nil && req.Description == nil {
			writeError(w, http.StatusBadRequest, "no fields to update")
			return
		}

		if req.Name != nil {
			trimmed := strings.TrimSpace(*req.Name)
			if trimmed == "" {
				writeError(w, http.StatusBadRequest, "name cannot be empty")
				return
			}
			req.Name = &trimmed
		}

		if req.Description != nil {
			trimmed := strings.TrimSpace(*req.Description)
			req.Description = &trimmed
		}

		item, found := ctx.Store.Update(id, req.Name, req.Description)
		if !found {
			writeError(w, http.StatusNotFound, "item not found")
			return
		}

		writeJSON(w, http.StatusOK, item)
	}
}

func deleteItemHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := itemIDFromPath(r.URL.Path)
		if !ok {
			writeError(w, http.StatusBadRequest, "invalid item id")
			return
		}

		if !ctx.Store.Delete(id) {
			writeError(w, http.StatusNotFound, "item not found")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
