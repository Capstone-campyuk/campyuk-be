package handler

import (
	"campyuk-api/features/user"
	"campyuk-api/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv user.UserService) user.UserHandler {
	return &userControll{
		srv: srv,
	}
}

func (uc *userControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}
		if input.Username == "" {
			return c.JSON(helper.ErrorResponse("username is empty"))
		} else if input.Password == "" {
			return c.JSON(helper.ErrorResponse("password is empty"))
		}

		token, res, err := uc.srv.Login(input.Username, input.Password)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success login", ToResponse(res), token))
	}
}

func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		res, err := uc.srv.Register(*ReqToCore(input))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}
		log.Println(res)
		return c.JSON(helper.SuccessResponse(http.StatusCreated, "success create account"))
	}
}

func (uc *userControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := uc.srv.Profile(token)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success show profile", GetToResponse(res)))
	}
}

func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := UpdateRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		formHeader, err := c.FormFile("user_image")
		if err != nil {
			log.Println(err)
		}

		_, err = uc.srv.Update(c.Get("user"), formHeader, *ReqToCore(input))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err.Error()))
		}

		return c.JSON(helper.SuccessResponse(http.StatusOK, "success update profile"))
	}
}

func (uc *userControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := uc.srv.Delete(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "internal server error",
			})
		}
		return c.NoContent(204)
	}
}
