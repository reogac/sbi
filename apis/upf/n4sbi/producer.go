/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:47 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n4sbi

import (
	"context"
	"fmt"
	"github.com/reogac/pfcp/message"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnAssociationRequest(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'callback'
	var callback *models.EndpointInfo
	callbackStr := ctx.Header("callback")
	if len(callbackStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "callback is required"), nil)
		return
	}

	if callback, err = models.EndpointInfoFromString(callbackStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse callback failed: %+v", err)), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(message.PFCPAssociationSetupRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleAssociationRequest(ctx.Context(), callback, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnDisassociationRequest(ctx sbi.RequestContext, prod Producer) {

	// read 'smfId'
	var smfId string
	smfId = ctx.Param("smfId")
	if len(smfId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smfId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDisassociationRequest(ctx.Context(), smfId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(201, nil, nil)

}

func OnSessionEstablishment(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'smfId'
	var smfId string
	smfId = ctx.Param("smfId")
	if len(smfId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smfId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(message.PFCPSessionEstablishmentRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleSessionEstablishment(ctx.Context(), smfId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnSessionModification(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'seid'
	var seid int64
	seidStr := ctx.Param("seid")
	if len(seidStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "seid is required"), nil)
		return
	}

	if seid, err = models.Int64FromString(seidStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse seid failed: %+v", err)), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(message.PFCPSessionModificationRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleSessionModification(ctx.Context(), seid, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnSessionDeletion(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'seid'
	var seid int64
	seidStr := ctx.Param("seid")
	if len(seidStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "seid is required"), nil)
		return
	}

	if seid, err = models.Int64FromString(seidStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse seid failed: %+v", err)), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(message.PFCPSessionDeletionRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleSessionDeletion(ctx.Context(), seid, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

type Producer interface {
	HandleAssociationRequest(context.Context, *models.EndpointInfo, *message.PFCPAssociationSetupRequest) (*message.PFCPAssociationSetupResponse, *models.ProblemDetails)

	HandleDisassociationRequest(context.Context, string) *models.ProblemDetails

	HandleSessionEstablishment(context.Context, string, *message.PFCPSessionEstablishmentRequest) (*message.PFCPSessionEstablishmentResponse, *models.ProblemDetails)

	HandleSessionModification(context.Context, int64, *message.PFCPSessionModificationRequest) (*message.PFCPSessionModificationResponse, *models.ProblemDetails)

	HandleSessionDeletion(context.Context, int64, *message.PFCPSessionDeletionRequest) (*message.PFCPSessionDeletionResponse, *models.ProblemDetails)
}
