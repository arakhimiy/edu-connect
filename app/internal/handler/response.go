package handler

import (
	"github.com/pocketbase/pocketbase/core"
	"gitlab.saidoff.uz/company/muslim-administration/mosque/back/internal/apperror"
)

type errorResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Content *string `json:"content"`
}
type successResponse struct {
	Success bool        `json:"success"`
	Message *string     `json:"message"`
	Content interface{} `json:"content"`
}

func (h *Handler) NewErrorResponse(e *core.RequestEvent, err *apperror.AppError) error {
	if err.IsDevErr {
		h.logger.Error(err.DeveloperMessage)
	}
	return e.JSON(err.StatusCode, errorResponse{false, err.Error(), nil})
}

func (h *Handler) NewSuccessResponse(e *core.RequestEvent, statusCode int, content interface{}) error {
	return e.JSON(statusCode, successResponse{true, nil, content})
}
