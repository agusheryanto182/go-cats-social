package controller

import (
	"net/http"
	"strconv"
	"strings"

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

// GetCat implements CatController.
func (c *CatControllerImpl) GetCat(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("CurrentUser").(dto.UserResWithID)

	userID := currentUser.ID

	idParam := r.URL.Query().Get("id")
	cleanedIdParam := strings.ReplaceAll(idParam, "\"", "")
	idINT, _ := strconv.Atoi(cleanedIdParam)

	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	race := r.URL.Query().Get("race")
	cleanedRace := strings.ReplaceAll(race, "\"", "")

	sex := r.URL.Query().Get("sex")
	cleanedSex := strings.ReplaceAll(sex, "\"", "")

	isAlreadyMatched := r.URL.Query().Get("isAlreadyMatched")

	ageInMonth := r.URL.Query().Get("ageInMonth")
	cleanedAgeInMonth := strings.ReplaceAll(ageInMonth, "\"", "")

	owned := r.URL.Query().Get("owned")

	search := r.URL.Query().Get("search")
	cleanedSearch := strings.ReplaceAll(search, "\"", "")

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	query := "SELECT id, user_id, name, race, sex, age_in_month, description, has_matched, image_urls, to_char(created_at AT TIME ZONE 'ASIA/JAKARTA', 'YYYY-MM-DD\"T\"HH24:MI:SS\"Z\"') AS created_at FROM cats WHERE 1=1"

	args := make([]interface{}, 0)

	if idINT > 0 {
		query += " AND id = $" + strconv.Itoa(len(args)+1)
		args = append(args, idINT)
	}

	if cleanedRace != "" {
		query += " AND race = $" + strconv.Itoa(len(args)+1)
		args = append(args, cleanedRace)
	}

	if cleanedSex != "" {
		query += " AND sex = $" + strconv.Itoa(len(args)+1)
		args = append(args, cleanedSex)
	}

	if isAlreadyMatched != "" {
		matched, _ := strconv.ParseBool(isAlreadyMatched)
		query += " AND has_matched = $" + strconv.Itoa(len(args)+1)
		args = append(args, matched)
	}

	if cleanedAgeInMonth != "" {
		var operator string
		switch cleanedAgeInMonth[0] {
		case '>':
			operator = ">"
		case '<':
			operator = "<"
		default:
			operator = "="
		}

		query += " AND age_in_month " + operator + " $" + strconv.Itoa(len(args)+1)
		args = append(args, cleanedAgeInMonth[1:])
	}

	if owned != "" {
		ownedBool, _ := strconv.ParseBool(owned)
		if ownedBool {
			query += " AND user_id = $" + strconv.Itoa(len(args)+1)
			args = append(args, int(userID))
		}
	}

	if cleanedSearch != "" {
		query += " AND name ILIKE $" + strconv.Itoa(len(args)+1)
		args = append(args, "%"+cleanedSearch+"%")
	}

	query += " AND deleted_at IS NULL"

	if limit == "" {
		limitInt = 5
	}

	if offset == "" {
		offsetInt = 0
	}

	query += " LIMIT $" + strconv.Itoa(len(args)+1)
	args = append(args, limitInt)

	query += " OFFSET $" + strconv.Itoa(len(args)+1)
	args = append(args, offsetInt)

	result, err := c.catSvc.GetByFilterAndArgs(r.Context(), query, args)
	if err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}

	helper.WriteResponse(w, web.OkResponse("successfully get cats", result))
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
