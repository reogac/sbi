/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ueauth

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"io"
)

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
	headers, rsp, prob := prod.HandleUeAuthenticationsPost(body)

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
	rsp, prob := prod.HandleUeAuthentications5gAkaConfirmationPut(authCtxId, body)

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
	prob := prod.HandleDelete5gAkaAuthenticationResult(authCtxId)

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
	rsp, prob := prod.HandleEapAuthMethod(authCtxId, body)

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

func OnDeleteProSeAuthenticationResult(ctx sbi.RequestContext, prod Producer) {

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteProSeAuthenticationResult(authCtxId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

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
	prob := prod.HandleUeAuthenticationsDeregisterPost(body)

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
	prob := prod.HandleDeleteEapAuthenticationResult(authCtxId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

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
	headers, rsp, prob := prod.HandleRgAuthenticationsPost(body)

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
	headers, rsp, prob := prod.HandleProseAuthenticationsPost(body)

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
	rsp, prob := prod.HandleProseAuth(authCtxId, body)

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
	HandleUeAuthenticationsPost(*models.AuthenticationInfo) (map[string]string, *models.UEAuthenticationCtx, *models.ProblemDetails)

	HandleUeAuthentications5gAkaConfirmationPut(string, *models.ConfirmationData) (*models.ConfirmationDataResponse, *models.ProblemDetails)

	HandleDelete5gAkaAuthenticationResult(string) *models.ProblemDetails

	HandleEapAuthMethod(string, *models.EapSession) (*models.EapSession, *models.ProblemDetails)

	HandleDeleteProSeAuthenticationResult(string) *models.ProblemDetails

	HandleUeAuthenticationsDeregisterPost(*models.DeregistrationInfo) *models.ProblemDetails

	HandleDeleteEapAuthenticationResult(string) *models.ProblemDetails

	HandleRgAuthenticationsPost(*models.RgAuthenticationInfo) (map[string]string, *models.RgAuthCtx, *models.ProblemDetails)

	HandleProseAuthenticationsPost(*models.ProSeAuthenticationInfo) (map[string]string, *models.ProSeAuthenticationCtx, *models.ProblemDetails)

	HandleProseAuth(string, *models.ProSeEapSession) (*models.ProseAuthResponse, *models.ProblemDetails)
}
