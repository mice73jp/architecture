package controller

import (
	"architecture/hexagonal-architecture/hex-arch-sample1-project/internal/service3"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Controller3 ...
type Controller3 struct {
	p *service3.Provider
}

// NewController3 ...
func NewController3(p *service3.Provider) *Controller3 {
	return &Controller3{p}
}

// HandleAccountOpen ...
func (ctrl *Controller3) HandleAccountOpen(c echo.Context) error {
	in := struct {
		Amount int `json:"amount"`
	}{}

	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// TODO: implement
	// if err := c.Validate(&in); err != nil {
	//  return c.JSON(http.StatusUnprocessableEntity, err.Error())
	// }

	ctx := c.Request().Context()
	psn, err := ctrl.p.OpenAccount(ctx, in.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, psn)
}

// HandleMoneyTransfer ...
func (ctrl *Controller3) HandleMoneyTransfer(c echo.Context) error {
	in := struct {
		FromAccountID int64 `json:"fromId"`
		ToAccountID   int64 `json:"toId"`
		Amount        int   `json:"amount"`
	}{}

	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// TODO: implement
	// if err := c.Validate(&in); err != nil {
	//  return c.JSON(http.StatusUnprocessableEntity, err.Error())
	// }

	ctx := c.Request().Context()
	from, to, err := ctrl.p.Transfer(ctx, in.Amount, in.FromAccountID, in.ToAccountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"from": from, "to": to})
}
