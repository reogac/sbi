package upmf2fe

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/models/n43"
)

// sbi producer handler for Topo Path :
func OnGetTopo(ctx sbi.RequestContext, handler interface{}) {
	prob := handler.(Producer)
	if rsp, prob := prob.GetTopo(); prob != nil {
		ctx.WriteResponse(prob.Status, prob)
	} else {
		ctx.WriteResponse(200, rsp)
	}
	return
}

type Producer interface {
	GetTopo() (*n43.TopoUpf, *models.ProblemDetails)
}