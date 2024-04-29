package controller

import (
	"net/http"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/helper"
	"github.com/agusheryanto182/go-social-media/internal/model/web"
	"github.com/agusheryanto182/go-social-media/internal/service"
	"github.com/go-playground/validator/v10"
)

type CatControllerImpl struct {
	catSvc    service.CatService
	validator *validator.Validate
}

// Create implements CatController.
func (c *CatControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("CurrentUser").(dto.UserResWithID)

	newCat := &dto.CatReq{}
	if err := helper.ReadFromRequestBody(r, newCat); err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}

	newCat.UserID = currentUser.ID

	if err := c.validator.Struct(newCat); err != nil {
		helper.WriteResponse(w, web.BadRequestResponse("request doesn't pass validation", err))
		return
	}

	result, err := c.catSvc.Create(r.Context(), newCat)
	if err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}

	helper.WriteResponse(w, web.OkResponse("successfully add cat", result))
}

func NewCatController(catSvc service.CatService, validator *validator.Validate) CatController {
	return &CatControllerImpl{
		catSvc:    catSvc,
		validator: validator,
	}
}
