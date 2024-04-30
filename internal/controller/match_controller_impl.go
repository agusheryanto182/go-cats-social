package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/agusheryanto182/go-social-media/internal/dto"
	"github.com/agusheryanto182/go-social-media/internal/helper"
	"github.com/agusheryanto182/go-social-media/internal/model/web"
	"github.com/agusheryanto182/go-social-media/internal/service"
	"github.com/go-playground/validator/v10"
)

type MatchControllerImpl struct {
	matchSvc service.MatchService
	catSvc   service.CatService
	valid    *validator.Validate
}

// Approve implements MatchController.
func (c *MatchControllerImpl) Approve(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("CurrentUser").(dto.UserResWithID)

	matchID := &dto.MatchIdReq{}

	if err := helper.ReadFromRequestBody(r, matchID); err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}

	if err := c.valid.Struct(matchID); err != nil {
		helper.WriteResponse(w, web.BadRequestResponse("validation error", err))
		return
	}

	matchID.MatchIdInt, _ = strconv.ParseUint(matchID.MatchID, 10, 64)

	isMatchExist, _ := c.matchSvc.IsMatchExist(r.Context(), matchID.MatchIdInt, currentUser.ID)
	if isMatchExist == nil {
		helper.WriteResponse(w, web.NotFoundResponse("not found", errors.New("match not found")))
		return
	}

	if isMatchExist.DeletedAt != nil {
		helper.WriteResponse(w, web.BadRequestResponse("bad request", errors.New("match id is no longer valid")))
		return
	}

	if err := c.matchSvc.ApproveTheMatch(r.Context(), matchID.MatchIdInt, currentUser.ID); err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}
	err := c.matchSvc.DeleteRequestByCatID(r.Context(), isMatchExist.MatchCatID, isMatchExist.UserCatID)
	if err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}
	err = c.catSvc.DoubleUpdateHasMatched(r.Context(), isMatchExist.MatchCatID, isMatchExist.UserCatID)
	if err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}

	helper.WriteResponse(w, web.OkResponse("success", "match approved"))

}

// GetMatch implements MatchController.
func (c *MatchControllerImpl) GetMatch(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("CurrentUser").(dto.UserResWithID)

	matchRes, err := c.matchSvc.GetMatch(r.Context(), currentUser.ID)
	if err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}

	helper.WriteResponse(w, web.OkResponse("successfully get match requests", matchRes))
}

// Match implements MatchController.
func (c *MatchControllerImpl) Match(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value("CurrentUser").(dto.UserResWithID)

	matchReq := &dto.MatchReq{}
	if err := helper.ReadFromRequestBody(r, matchReq); err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}

	matchReq.IssuedBy = currentUser.ID

	if err := c.valid.Struct(matchReq); err != nil {
		helper.WriteResponse(w, web.BadRequestResponse("validation error", err))
		return
	}

	matchCatIdInt, _ := strconv.ParseUint(matchReq.MatchCatID, 10, 64)
	userCatIdInt, _ := strconv.ParseUint(matchReq.UserCatID, 10, 64)
	matchCatDetail, _ := c.catSvc.GetByID(r.Context(), matchCatIdInt)
	userCatDetail, _ := c.catSvc.GetByIdAndUserID(r.Context(), userCatIdInt, currentUser.ID)
	matchReq.MatchCatInt = matchCatIdInt
	matchReq.UserCatInt = userCatIdInt
	matchReq.ReceiverBy = matchCatDetail.UserID

	isReqExist, _ := c.matchSvc.IsRequestExist(r.Context(), uint64(matchCatIdInt), uint64(userCatIdInt))
	if isReqExist {
		helper.WriteResponse(w, web.BadRequestResponse("bad request", errors.New("already matched or requested")))
		return
	}

	if matchCatDetail == nil || userCatDetail == nil {
		helper.WriteResponse(w, web.NotFoundResponse("not found", errors.New("cat not found")))
		return
	}

	if matchReq.MatchCatID == matchReq.UserCatID {
		helper.WriteResponse(w, web.BadRequestResponse("bad request", errors.New("you can't match with your own cat")))
		return
	}

	if matchCatDetail.UserID == userCatDetail.UserID {
		helper.WriteResponse(w, web.BadRequestResponse("bad request", errors.New("you can't match with your own cat")))
		return
	}

	if matchCatDetail.Sex == userCatDetail.Sex {
		helper.WriteResponse(w, web.BadRequestResponse("bad request", errors.New("you can't match with same sex cat")))
		return
	}

	if matchCatDetail.HasMatched || userCatDetail.HasMatched {
		helper.WriteResponse(w, web.BadRequestResponse("bad request", errors.New("you can't match with already matched cat")))
		return
	}

	if err := c.matchSvc.Create(r.Context(), matchReq); err != nil {
		helper.WriteResponse(w, web.InternalServerErrorResponse("internal server error", err))
		return
	}

	helper.WriteResponse(w, web.OkResponse("successfully send match request", "success"))
}

func NewMatchController(matchSvc service.MatchService, catSvc service.CatService, valid *validator.Validate) MatchController {
	return &MatchControllerImpl{
		matchSvc: matchSvc,
		catSvc:   catSvc,
		valid:    valid,
	}
}
