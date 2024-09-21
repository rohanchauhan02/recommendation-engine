package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	dto "github.com/rohanchauhan02/recommendation-engine/dto/medicine"
	"github.com/rohanchauhan02/recommendation-engine/modules/medicine"
)

type handler struct {
	usecase medicine.Usecase
}

func NewMedicineHandler(e *echo.Echo, usecase *medicine.Usecase) {
	handler := handler{
		usecase: *usecase,
	}
	api := e.Group("/api/v1")

	api.GET("/hello", handler.HelloWorld)
	api.PUT("/medicine", handler.AddMedicine)
}

func (h *handler) HelloWorld(c echo.Context) error {
	return c.JSON(200, "Hello Rohan!")
}

func (h *handler) AddMedicine(c echo.Context) error {
	req := &dto.CreateMedicineRequest{}
	if err := c.Bind(req); err != nil {
		errMsg := fmt.Errorf("failed to validate request body")
		return errMsg
	}
	err := h.usecase.AddMedicine(req)
	if err != nil {
		return err
	}
	return c.JSON(200, "Success to add medicine data")
}
