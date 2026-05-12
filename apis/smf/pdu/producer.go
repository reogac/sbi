/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package pdu

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"io"
)

func OnRetrievePduSession(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionRef is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.RetrieveData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleRetrievePduSession(ctx.Context(), pduSessionRef, body)

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

func OnSendMoData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smContextRef is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SendMoDataRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	ersp, prob := prod.HandleSendMoData(ctx.Context(), smContextRef, body)

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromExtProblemDetails(ersp), ersp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnPostPduSessions(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PostPduSessionsRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp, ersp, prob := prod.HandlePostPduSessions(ctx.Context(), body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromPostPduSessionsErrorResponse(ersp), ersp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnUpdatePduSession(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionRef is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UpdatePduSessionRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleUpdatePduSession(ctx.Context(), pduSessionRef, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromUpdatePduSessionErrorResponse(ersp), ersp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReleasePduSession(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionRef is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	var body *models.ReleasePduSessionRequest
	body = new(models.ReleasePduSessionRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	rsp, prob := prod.HandleReleasePduSession(ctx.Context(), pduSessionRef, body)

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

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnTransferMoData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionRef is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.TransferMoDataRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleTransferMoData(ctx.Context(), pduSessionRef, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnPostSmContexts(ctx sbi.RequestContext, prod Producer) {
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
	body := new(models.PostSmContextsRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp, ersp, prob := prod.HandlePostSmContexts(ctx.Context(), callback, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromPostSmContextsErrorResponse(ersp), ersp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnRetrieveSmContext(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smContextRef is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	var body *models.SmContextRetrieveData
	body = new(models.SmContextRetrieveData)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	rsp, prob := prod.HandleRetrieveSmContext(ctx.Context(), smContextRef, body)

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

func OnUpdateSmContext(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smContextRef is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UpdateSmContextRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleUpdateSmContext(ctx.Context(), smContextRef, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for defined error
	if ersp != nil {
		ctx.WriteResponse(models.StatusFromUpdateSmContextErrorResponse(ersp), ersp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnReleaseSmContext(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "smContextRef is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	var body *models.ReleaseSmContextRequest
	body = new(models.ReleaseSmContextRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	rsp, prob := prod.HandleReleaseSmContext(ctx.Context(), smContextRef, body)

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

	// success
	ctx.WriteResponse(204, nil, nil)

}

type Producer interface {
	HandleRetrievePduSession(context.Context, string, *models.RetrieveData) (*models.RetrievedData, *models.ProblemDetails)

	HandleSendMoData(context.Context, string, *models.SendMoDataRequest) (*models.ExtProblemDetails, *models.ProblemDetails)

	HandlePostPduSessions(context.Context, *models.PostPduSessionsRequest) (map[string]string, *models.PostPduSessionsResponse, *models.PostPduSessionsErrorResponse, *models.ProblemDetails)

	HandleUpdatePduSession(context.Context, string, *models.UpdatePduSessionRequest) (*models.UpdatePduSessionResponse, *models.UpdatePduSessionErrorResponse, *models.ProblemDetails)

	HandleReleasePduSession(context.Context, string, *models.ReleasePduSessionRequest) (*models.ReleasePduSessionResponse, *models.ProblemDetails)

	HandleTransferMoData(context.Context, string, *models.TransferMoDataRequest) *models.ProblemDetails

	HandlePostSmContexts(context.Context, *models.EndpointInfo, *models.PostSmContextsRequest) (map[string]string, *models.PostSmContextsResponse, *models.PostSmContextsErrorResponse, *models.ProblemDetails)

	HandleRetrieveSmContext(context.Context, string, *models.SmContextRetrieveData) (*models.SmContextRetrievedData, *models.ProblemDetails)

	HandleUpdateSmContext(context.Context, string, *models.UpdateSmContextRequest) (*models.UpdateSmContextResponse, *models.UpdateSmContextErrorResponse, *models.ProblemDetails)

	HandleReleaseSmContext(context.Context, string, *models.ReleaseSmContextRequest) (*models.SmContextReleasedData, *models.ProblemDetails)
}
