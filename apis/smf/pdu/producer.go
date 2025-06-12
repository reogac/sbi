/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package pdu

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnRetrievePduSession(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionRef is required"))
		return
	}

	// decode request body
	body := new(models.RetrieveData)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleRetrievePduSession(pduSessionRef, body)

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

func OnRetrieveSmContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smContextRef is required"))
		return
	}

	// decode request body
	var body *models.SmContextRetrieveData
	if ctx.HasRequestBody() {
		body = new(models.SmContextRetrieveData)
		if err = ctx.DecodeRequest(body); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleRetrieveSmContext(smContextRef, body)

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

func OnUpdateSmContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smContextRef is required"))
		return
	}

	// decode request body
	body := new(models.UpdateSmContextRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleUpdateSmContext(smContextRef, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromUpdateSmContextErrorResponse(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnUpdatePduSession(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionRef is required"))
		return
	}

	// decode request body
	body := new(models.UpdatePduSessionRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleUpdatePduSession(pduSessionRef, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromUpdatePduSessionErrorResponse(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnPostPduSessions(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.PostPduSessionsRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandlePostPduSessions(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromPostPduSessionsErrorResponse(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnReleasePduSession(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionRef is required"))
		return
	}

	// decode request body
	var body *models.ReleasePduSessionRequest
	if ctx.HasRequestBody() {
		body = new(models.ReleasePduSessionRequest)
		if err = ctx.DecodeRequest(body); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleReleasePduSession(pduSessionRef, body)

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

	// success
	ctx.WriteResponse(204, nil)

}

func OnTransferMoData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionRef is required"))
		return
	}

	// decode request body
	body := new(models.TransferMoDataRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleTransferMoData(pduSessionRef, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnPostSmContexts(ctx sbi.RequestContext, handler any) {
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
	body := new(models.PostSmContextsRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandlePostSmContexts(callback, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromPostSmContextsErrorResponse(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnReleaseSmContext(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smContextRef is required"))
		return
	}

	// decode request body
	var body *models.ReleaseSmContextRequest
	if ctx.HasRequestBody() {
		body = new(models.ReleaseSmContextRequest)
		if err = ctx.DecodeRequest(body); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleReleaseSmContext(smContextRef, body)

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

	// success
	ctx.WriteResponse(204, nil)

}

func OnSendMoData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smContextRef is required"))
		return
	}

	// decode request body
	body := new(models.SendMoDataRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	ersp, prob := prod.HandleSendMoData(smContextRef, body)

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromExtProblemDetails(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

type Producer interface {
	HandleRetrievePduSession(string, *models.RetrieveData) (*models.RetrievedData, *models.ProblemDetails)

	HandleRetrieveSmContext(string, *models.SmContextRetrieveData) (*models.SmContextRetrievedData, *models.ProblemDetails)

	HandleUpdateSmContext(string, *models.UpdateSmContextRequest) (*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails)

	HandleUpdatePduSession(string, *models.UpdatePduSessionRequest) (*models.UpdatePduSessionResponse, *models.UpdatePduSessionErrorResponse, *models.ProblemDetails)

	HandlePostPduSessions(*models.PostPduSessionsRequest) (*models.PostPduSessionsResponse, *models.PostPduSessionsErrorResponse, *models.ProblemDetails)

	HandleReleasePduSession(string, *models.ReleasePduSessionRequest) (*models.ReleasePduSessionResponse, *models.ProblemDetails)

	HandleTransferMoData(string, *models.TransferMoDataRequest) *models.ProblemDetails

	HandlePostSmContexts(*models.EndpointInfo, *models.PostSmContextsRequest) (*models.PostSmContextsResponse, *models.PostSmContextsErrorResponse, *models.ProblemDetails)

	HandleReleaseSmContext(string, *models.ReleaseSmContextRequest) (*models.SmContextReleasedData, *models.ProblemDetails)

	HandleSendMoData(string, *models.SendMoDataRequest) (*models.ExtProblemDetails, *models.ProblemDetails)
}
