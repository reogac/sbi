/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:26 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ueauth

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnGenerateAv(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GenerateAvParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'hssAuthType'
	params.HssAuthType = ctx.Param("hssAuthType")
	if len(params.HssAuthType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "hssAuthType is required"))
		return
	}

	// decode request body
	body := new(models.HssAuthenticationInfoRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGenerateAv(&params, body)

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

func OnDeleteAuth(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params DeleteAuthParams

	// read 'authEventId'
	params.AuthEventId = ctx.Param("authEventId")
	if len(params.AuthEventId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authEventId is required"))
		return
	}

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	body := new(models.AuthEvent)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleDeleteAuth(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

	// success
	ctx.WriteResponse(204, nil)

}

func OnGenerateGbaAv(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	body := new(models.GbaAuthenticationInfoRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGenerateGbaAv(supi, body)

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

func OnGenerateProseAV(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'supiOrSuci'
	var supiOrSuci string
	supiOrSuci = ctx.Param("supiOrSuci")
	if len(supiOrSuci) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supiOrSuci is required"))
		return
	}

	// decode request body
	body := new(models.ProSeAuthenticationInfoRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGenerateProseAV(supiOrSuci, body)

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

func OnGenerateAuthData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'supiOrSuci'
	var supiOrSuci string
	supiOrSuci = ctx.Param("supiOrSuci")
	if len(supiOrSuci) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supiOrSuci is required"))
		return
	}

	// decode request body
	body := new(models.AuthenticationInfoRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGenerateAuthData(supiOrSuci, body)

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

func OnGetRgAuthData(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetRgAuthDataParams

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'supiOrSuci'
	params.SupiOrSuci = ctx.Param("supiOrSuci")
	if len(params.SupiOrSuci) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supiOrSuci is required"))
		return
	}

	// read 'authenticated-ind'
	authenticatedIndStr := ctx.Param("authenticated-ind")
	if len(authenticatedIndStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authenticated-ind is required"))
		return
	}

	if params.AuthenticatedInd, err = models.BoolFromString(authenticatedIndStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse authenticated-ind failed: %+v", err)))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetRgAuthData(&params)

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

func OnConfirmAuth(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	body := new(models.AuthEvent)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleConfirmAuth(supi, body)

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
	HandleGenerateAv(*GenerateAvParams, *models.HssAuthenticationInfoRequest) (*models.HssAuthenticationInfoResult, *models.ProblemDetails)

	HandleDeleteAuth(*DeleteAuthParams, *models.AuthEvent) *models.ProblemDetails

	HandleGenerateGbaAv(string, *models.GbaAuthenticationInfoRequest) (*models.GbaAuthenticationInfoResult, *models.ProblemDetails)

	HandleGenerateProseAV(string, *models.ProSeAuthenticationInfoRequest) (*models.ProSeAuthenticationInfoResult, *models.ProblemDetails)

	HandleGenerateAuthData(string, *models.AuthenticationInfoRequest) (*models.AuthenticationInfoResult, *models.ProblemDetails)

	HandleGetRgAuthData(*GetRgAuthDataParams) (*models.RgAuthCtx, *models.ProblemDetails)

	HandleConfirmAuth(string, *models.AuthEvent) (*models.AuthEvent, *models.ProblemDetails)
}
