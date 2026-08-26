/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:02:44 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package conf

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnGetConfiguration(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.ConfigurationRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetConfiguration(ctx.Context(), body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

type Producer interface {
	HandleGetConfiguration(context.Context, *models.ConfigurationRequest) (*models.ConfigurationResponse, *models.ProblemDetails)
}
