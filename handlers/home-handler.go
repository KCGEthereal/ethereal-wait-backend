package handlers

import (
	"net/http"

	"github.com/KCGEthereal/ethereal-wait-backend/common"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.Home()
	if err != nil {
		common.NewResponse(common.ERROR, "It does not work").
			SetError(err).
			Respond(w)
		return
	}

	common.NewResponse(common.SUCCESS, "It works").
		SetContent(result).
		Respond(w)
}
