package controller

import (
	"errors"
	"net/http"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/helper"
	"github.com/agusheryanto182/go-social-media/internal/model/web"
	"github.com/agusheryanto182/go-social-media/internal/service"
	"github.com/go-playground/validator/v10"
)

type UserControllerImpl struct {
	svc       service.UserService
	validator *validator.Validate
}

// Login implements UserController.
func (c *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	loginReq := &dto.UserLoginReq{}
	if err := helper.ReadFromRequestBody(r, loginReq); err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}

	if err := c.validator.Struct(loginReq); err != nil {
		helper.WriteResponse(w, web.BadRequestResponse("request doesn't pass validation", err))
		return
	}

	isUserExist, _ := c.svc.IsEmailExist(r.Context(), loginReq.Email)
	if !isUserExist {
		helper.WriteResponse(w, web.NotFoundResponse("not found", errors.New("user not found")))
		return
	}

	result, err := c.svc.Login(r.Context(), loginReq)
	if err != nil {
		switch {
		case err.Error() == "wrong password":
			helper.WriteResponse(w, web.BadRequestResponse("bad request", err))
			return
		default:
			helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
			return
		}
	}

	helper.WriteResponse(w, web.OkResponse("User successfully logged", result))
}

// Register implements UserController.
func (c *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	newUser := &dto.UserReq{}
	if err := helper.ReadFromRequestBody(r, newUser); err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}

	if err := c.validator.Struct(newUser); err != nil {
		helper.WriteResponse(w, web.BadRequestResponse("request doesn't pass validation", err))
		return
	}

	isEmailExist, _ := c.svc.IsEmailExist(r.Context(), newUser.Email)
	if isEmailExist {
		helper.WriteResponse(w, web.ConflictResponse("conflict", errors.New("email already exist")))
		return
	}

	result, err := c.svc.Create(r.Context(), newUser)
	if err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}
	helper.WriteResponse(w, web.CreatedResponse("User successfully registered", result))
}

func NewUserController(svc service.UserService, validator *validator.Validate) UserController {
	return &UserControllerImpl{
		svc:       svc,
		validator: validator,
	}
}
