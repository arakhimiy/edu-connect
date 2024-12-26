package handler

import (
	"github.com/pocketbase/pocketbase/core"
	"net/http"
)

func (h *Handler) TestHandler(e *core.RequestEvent) error {
	return h.NewSuccessResponse(e, http.StatusOK, "ok")
}
