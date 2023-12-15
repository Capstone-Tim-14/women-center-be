package handlers

import (
	"strings"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type CounselingSessionHandlers interface {
	ListCounselingSessionHandler(echo.Context) error
}

type CounselingSessionHandlerImpl struct {
	CounselingSession services.CounselingSessionService
}

func NewCounselingSessionHandlers(counselingSession CounselingSessionHandlerImpl) CounselingSessionHandlers {
	return &counselingSession
}

func (handler *CounselingSessionHandlerImpl) ListCounselingSessionHandler(ctx echo.Context) error {

	Response, err := handler.CounselingSession.GetListCounselingSession(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Counselor not found") {
			return exceptions.StatusInternalServerError(ctx, err)
		}
	}

	return responses.StatusOK(ctx, "Success get list counseling session", Response)
}
