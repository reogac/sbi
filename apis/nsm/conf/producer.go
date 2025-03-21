/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Mar 21 10:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package conf

import (
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnGetNssfConfiguration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// call application handler
	rsp, prob := prod.HandleGetNssfConfiguration()

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

type Producer interface {
	HandleGetNssfConfiguration() (*models.NssfConfiguration, *models.ProblemDetails)
}
