/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ampol

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnCreateIndividualAMPolicyAssociation(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PolicyAssociationRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleCreateIndividualAMPolicyAssociation(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnReadIndividualAMPolicyAssociation(ctx sbi.RequestContext, prod Producer) {

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadIndividualAMPolicyAssociation(polAssoId)

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

func OnDeleteIndividualAMPolicyAssociation(ctx sbi.RequestContext, prod Producer) {

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteIndividualAMPolicyAssociation(polAssoId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReportObservedEventTriggersForIndividualAMPolicyAssociation(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "polAssoId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PolicyAssociationUpdateRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(polAssoId, body)

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
	HandleCreateIndividualAMPolicyAssociation(*models.PolicyAssociationRequest) (map[string]string, *models.PolicyAssociation, *models.ProblemDetails)

	HandleReadIndividualAMPolicyAssociation(string) (*models.PolicyAssociation, *models.ProblemDetails)

	HandleDeleteIndividualAMPolicyAssociation(string) *models.ProblemDetails

	HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(string, *models.PolicyAssociationUpdateRequest) (*models.PolicyUpdate, *models.ProblemDetails)
}
