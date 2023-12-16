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
	CounselingSessionDetailHandler(echo.Context) error
}

type CounselingSessionHandlerImpl struct {
	CounselingSession services.CounselingSessionService
}

func NewCounselingSessionHandlers(counselingSession CounselingSessionHandlerImpl) CounselingSessionHandlers {
	return &counselingSession
}

func (handler *CounselingSessionHandlerImpl) CounselingSessionDetailHandler(ctx echo.Context) error {

	response, err := handler.CounselingSession.GetCounselingSessionDetail(ctx)

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Success get counseling session detail", response)
}

func (handler *CounselingSessionHandlerImpl) ListCounselingSessionHandler(ctx echo.Context) error {

	Response, err := handler.CounselingSession.GetListCounselingSession(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Counseling session empty") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Success get list counseling session", Response)
}
