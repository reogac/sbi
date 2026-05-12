/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:41 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ueauth

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"io"
)

func OnUeAuthenticationsDeregisterPost(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.DeregistrationInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUeAuthenticationsDeregisterPost(ctx.Context(), body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnEapAuthMethod(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	var body *models.EapSession
	body = new(models.EapSession)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	rsp, prob := prod.HandleEapAuthMethod(ctx.Context(), authCtxId, body)

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

func OnRgAuthenticationsPost(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.RgAuthenticationInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleRgAuthenticationsPost(ctx.Context(), body)

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

func OnProseAuthenticationsPost(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.ProSeAuthenticationInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleProseAuthenticationsPost(ctx.Context(), body)

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

func OnProseAuth(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	var body *models.ProSeEapSession
	body = new(models.ProSeEapSession)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	rsp, prob := prod.HandleProseAuth(ctx.Context(), authCtxId, body)

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

func OnUeAuthenticationsPost(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AuthenticationInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleUeAuthenticationsPost(ctx.Context(), body)

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

func OnUeAuthentications5gAkaConfirmationPut(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	var body *models.ConfirmationData
	body = new(models.ConfirmationData)
	if err = sbi.Decode(contentLength, content, body); err != nil && err != io.EOF {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	} else if err == io.EOF {
		body = nil
	}
	// call application handler
	rsp, prob := prod.HandleUeAuthentications5gAkaConfirmationPut(ctx.Context(), authCtxId, body)

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

func OnDelete5gAkaAuthenticationResult(ctx sbi.RequestContext, prod Producer) {

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDelete5gAkaAuthenticationResult(ctx.Context(), authCtxId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteEapAuthenticationResult(ctx sbi.RequestContext, prod Producer) {

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteEapAuthenticationResult(ctx.Context(), authCtxId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteProSeAuthenticationResult(ctx sbi.RequestContext, prod Producer) {

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteProSeAuthenticationResult(ctx.Context(), authCtxId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

type Producer interface {
	HandleUeAuthenticationsDeregisterPost(context.Context, *models.DeregistrationInfo) *models.ProblemDetails

	HandleEapAuthMethod(context.Context, string, *models.EapSession) (*models.EapSession, *models.ProblemDetails)

	HandleRgAuthenticationsPost(context.Context, *models.RgAuthenticationInfo) (map[string]string, *models.RgAuthCtx, *models.ProblemDetails)

	HandleProseAuthenticationsPost(context.Context, *models.ProSeAuthenticationInfo) (map[string]string, *models.ProSeAuthenticationCtx, *models.ProblemDetails)

	HandleProseAuth(context.Context, string, *models.ProSeEapSession) (*models.ProseAuthResponse, *models.ProblemDetails)

	HandleUeAuthenticationsPost(context.Context, *models.AuthenticationInfo) (map[string]string, *models.UEAuthenticationCtx, *models.ProblemDetails)

	HandleUeAuthentications5gAkaConfirmationPut(context.Context, string, *models.ConfirmationData) (*models.ConfirmationDataResponse, *models.ProblemDetails)

	HandleDelete5gAkaAuthenticationResult(context.Context, string) *models.ProblemDetails

	HandleDeleteEapAuthenticationResult(context.Context, string) *models.ProblemDetails

	HandleDeleteProSeAuthenticationResult(context.Context, string) *models.ProblemDetails
}
