package handlers

import (
	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *handler) SignIn(ctx echo.Context) error {
	return ctx.Render(200, "auth.html", map[string]interface{}{"type": "signin"})
}

func (h *handler) SignUp(ctx echo.Context) error {
	return ctx.Render(200, "auth.html", map[string]interface{}{"type": "signup"})
}

func (h *handler) NewUser(ctx echo.Context) error {
	usr := models.User{
		Name:     ctx.FormValue("name"),
		Password: ctx.FormValue("password"),
	}
	if err := h.usecase.NewUser(usr); err != nil {
		return ctx.Render(401, "auth.html", map[string]interface{}{
			"type":  "signup",
			"error": err.Error(),
		})
	}
	return ctx.Redirect(301, "/")
}
