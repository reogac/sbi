package upmf2activate

import (
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"github.com/reogac/sbi/models/n42"
	"net/http"
)

// sbi producer handler for activate, deactivate Upfs :
func UpfsActivate(ctx sbi.RequestContext, handler interface{}) {
	prod := handler.(Producer)
	var query *n42.UpfActivateQuery

	if err := ctx.DecodeRequest(&query); err == nil {
		if rsp, prob := prod.HandleActivateFe(query); prob == nil {
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

func UpfsDeActivate(ctx sbi.RequestContext, handler interface{}) {
	prod := handler.(Producer)
	var query *n42.UpfDeactivateQuery

	if err := ctx.DecodeRequest(&query); err == nil {
		if rsp, prob := prod.HandleDeactivateFe(query); prob == nil {
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
	//HandleGetUpfPath(query *n43.UpfPathQuery) (*n43.UpfPath, *models.ProblemDetails)
	HandleActivateFe(query *n42.UpfActivateQuery) (*n42.UpfActivate, *models.ProblemDetails)
	HandleDeactivateFe(query *n42.UpfDeactivateQuery) (*n42.UpfDeactivate, *models.ProblemDetails)
}
