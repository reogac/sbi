package upf2upmf

import (
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"github.com/reogac/sbi/models/n42"
	"net/http"
)

func OnHeartbeat(ctx sbi.RequestContext, handler interface{}) {
	prod := handler.(Producer)
	var req n42.HeartbeatRequest

	if err := ctx.DecodeRequest(&req); err == nil {
		if rsp, prob := prod.HandleHeartbeat(&req); prob != nil {
			ctx.WriteResponse(prob.Status, prob)
		} else {
			ctx.WriteResponse(200, rsp)
		}
	} else {
		ctx.WriteResponse(http.StatusBadRequest, &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Detail: err.Error(),
		})

	}
	return
}

func OnActivate(ctx sbi.RequestContext, handler interface{}) {
	prod := handler.(Producer)
	var query *n42.UpfActivateQuery

	if err := ctx.DecodeRequest(&query); err == nil {
		if rsp, prob := prod.HandleActivate(query); prob == nil {
			ctx.WriteResponse(200, rsp.Msg)
		} else {
			ctx.WriteResponse(prob.Status, prob)
		}
	} else {
		ctx.WriteResponse(http.StatusBadRequest, &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Detail: err.Error(),
		})
	}
	return
}

func OnDeactivate(ctx sbi.RequestContext, handler interface{}) {
	prod := handler.(Producer)
	var query *n42.UpfDeactivateQuery

	if err := ctx.DecodeRequest(&query); err == nil {
		if rsp, prob := prod.HandleDeactivate(query); prob == nil {
			ctx.WriteResponse(200, rsp.Msg)
		} else {
			ctx.WriteResponse(prob.Status, prob)
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
	HandleHeartbeat(req *n42.HeartbeatRequest) (*n42.HeartbeatResponse, *models.ProblemDetails)
	HandleActivate(req *n42.UpfActivateQuery) (*n42.UpfActivate, *models.ProblemDetails)
	HandleDeactivate(req *n42.UpfDeactivateQuery) (*n42.UpfDeactivate, *models.ProblemDetails)
}
