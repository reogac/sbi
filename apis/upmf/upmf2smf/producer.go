package upmf2smf

import (
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"github.com/reogac/sbi/models/n43"
	"net/http"
)

func OnGetUpfPath(ctx sbi.RequestContext, handler interface{}) {
	prod := handler.(Producer)
	var query n43.UpfPathQuery

	if err := ctx.DecodeRequest(&query); err == nil {
		if rsp, prob := prod.HandleGetUpfPath(&query); prob != nil {
			ctx.WriteResponse(prob.Status, prob)
		} else {
			ctx.WriteResponse(http.StatusCreated, rsp)
		}
	} else {
		ctx.WriteResponse(http.StatusBadRequest, &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Detail: err.Error(),
		})

	}
	return
}

func OnSmfSubscribe(ctx sbi.RequestContext, handler interface{}) {
	prod := handler.(Producer)
	var req n43.SmfSubscriptionRequest

	if err := ctx.DecodeRequest(&req); err == nil {
		if rsp, prob := prod.HandleSmfSubscribe(&req); prob != nil {
			ctx.WriteResponse(prob.Status, prob)
		} else {
			ctx.WriteResponse(http.StatusCreated, rsp)
		}
	} else {
		ctx.WriteResponse(http.StatusBadRequest, &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Detail: err.Error(),
		})

	}
	return
}

type Producer interface {
	HandleGetUpfPath(query *n43.UpfPathQuery) (*n43.UpfPath, *models.ProblemDetails)
	HandleSmfSubscribe(req *n43.SmfSubscriptionRequest) (*n43.SmfSubscriptionResponse, *models.ProblemDetails)
}
