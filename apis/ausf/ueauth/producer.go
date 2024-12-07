/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:29 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ueauth

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnDelete5gAkaAuthenticationResult(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"))
		return
	}

	// call application handler
	prob := prod.HandleDelete5gAkaAuthenticationResult(authCtxId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnRgAuthenticationsPost(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.RgAuthenticationInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleRgAuthenticationsPost(body)

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

func OnProseAuthenticationsPost(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.ProSeAuthenticationInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleProseAuthenticationsPost(body)

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

func OnDeleteProSeAuthenticationResult(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"))
		return
	}

	// call application handler
	prob := prod.HandleDeleteProSeAuthenticationResult(authCtxId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnUeAuthenticationsPost(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.AuthenticationInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleUeAuthenticationsPost(body)

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

func OnUeAuthentications5gAkaConfirmationPut(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"))
		return
	}

	// decode request body
	var body *models.ConfirmationData
	if ctx.HasRequestBody() {
		body = new(models.ConfirmationData)
		if err = ctx.DecodeRequest(body); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleUeAuthentications5gAkaConfirmationPut(authCtxId, body)

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

func OnEapAuthMethod(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"))
		return
	}

	// decode request body
	var body *models.EapSession
	if ctx.HasRequestBody() {
		body = new(models.EapSession)
		if err = ctx.DecodeRequest(body); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleEapAuthMethod(authCtxId, body)

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

func OnDeleteEapAuthenticationResult(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"))
		return
	}

	// call application handler
	prob := prod.HandleDeleteEapAuthenticationResult(authCtxId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnProseAuth(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authCtxId is required"))
		return
	}

	// decode request body
	var body *models.ProSeEapSession
	if ctx.HasRequestBody() {
		body = new(models.ProSeEapSession)
		if err = ctx.DecodeRequest(body); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleProseAuth(authCtxId, body)

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

func OnUeAuthenticationsDeregisterPost(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.DeregistrationInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleUeAuthenticationsDeregisterPost(body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

type Producer interface {
	HandleDelete5gAkaAuthenticationResult(string) *models.ProblemDetails

	HandleRgAuthenticationsPost(*models.RgAuthenticationInfo) (*models.RgAuthCtx, *models.ProblemDetails)

	HandleProseAuthenticationsPost(*models.ProSeAuthenticationInfo) (*models.ProSeAuthenticationCtx, *models.ProblemDetails)

	HandleDeleteProSeAuthenticationResult(string) *models.ProblemDetails

	HandleUeAuthenticationsPost(*models.AuthenticationInfo) (*models.UEAuthenticationCtx, *models.ProblemDetails)

	HandleUeAuthentications5gAkaConfirmationPut(string, *models.ConfirmationData) (*models.ConfirmationDataResponse, *models.ProblemDetails)

	HandleEapAuthMethod(string, *models.EapSession) (*models.EapSession, *models.ProblemDetails)

	HandleDeleteEapAuthenticationResult(string) *models.ProblemDetails

	HandleProseAuth(string, *models.ProSeEapSession) (*models.ProseAuthResponse, *models.ProblemDetails)

	HandleUeAuthenticationsDeregisterPost(*models.DeregistrationInfo) *models.ProblemDetails
}
