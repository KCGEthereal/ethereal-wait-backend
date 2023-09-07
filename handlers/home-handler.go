package handlers

import (
	"github.com/esportsclub/entity-service-golang/common"
	"net/http"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	common.NewResponse(common.SUCCESS, "It works").
		SetContent(h.Service.Home()).
		Respond(w)
}
