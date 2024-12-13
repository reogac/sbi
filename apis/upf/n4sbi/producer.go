/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:35 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n4sbi

import (
	"github.com/reogac/pfcp/message"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnSessionModification(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'seid'
	var seid int64
	seidStr := ctx.Param("seid")
	if len(seidStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "seid is required"))
		return
	}

	if seid, err = models.Int64FromString(seidStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse seid failed: %+v", err)))
		return
	}

	// decode request body
	body := new(message.PFCPSessionModificationRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleSessionModification(seid, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnSessionDeletion(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'seid'
	var seid int64
	seidStr := ctx.Param("seid")
	if len(seidStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "seid is required"))
		return
	}

	if seid, err = models.Int64FromString(seidStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse seid failed: %+v", err)))
		return
	}

	// decode request body
	body := new(message.PFCPSessionDeletionRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleSessionDeletion(seid, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnAssociationRequest(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'callback'
	var callback *models.EndpointInfo
	callbackStr := ctx.Header("callback")
	if len(callbackStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "callback is required"))
		return
	}

	if callback, err = models.EndpointInfoFromString(callbackStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse callback failed: %+v", err)))
		return
	}

	// decode request body
	body := new(message.PFCPAssociationSetupRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleAssociationRequest(callback, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnDisassociationRequest(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'smfId'
	var smfId string
	smfId = ctx.Param("smfId")
	if len(smfId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smfId is required"))
		return
	}

	// call application handler
	prob := prod.HandleDisassociationRequest(smfId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(201, nil)

}

func OnSessionEstablishment(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'smfId'
	var smfId string
	smfId = ctx.Param("smfId")
	if len(smfId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smfId is required"))
		return
	}

	// decode request body
	body := new(message.PFCPSessionEstablishmentRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleSessionEstablishment(smfId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

type Producer interface {
	HandleSessionModification(int64, *message.PFCPSessionModificationRequest) (*message.PFCPSessionModificationResponse, *models.ProblemDetails)

	HandleSessionDeletion(int64, *message.PFCPSessionDeletionRequest) (*message.PFCPSessionDeletionResponse, *models.ProblemDetails)

	HandleAssociationRequest(*models.EndpointInfo, *message.PFCPAssociationSetupRequest) (*message.PFCPAssociationSetupResponse, *models.ProblemDetails)

	HandleDisassociationRequest(string) *models.ProblemDetails

	HandleSessionEstablishment(string, *message.PFCPSessionEstablishmentRequest) (*message.PFCPSessionEstablishmentResponse, *models.ProblemDetails)
}
