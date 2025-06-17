/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:54 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ueauth

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnConfirmAuth(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AuthEvent)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleConfirmAuth(supi, body)

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

func OnGenerateAv(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GenerateAvParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'hssAuthType'
	params.HssAuthType = ctx.Param("hssAuthType")
	if len(params.HssAuthType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "hssAuthType is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.HssAuthenticationInfoRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGenerateAv(&params, body)

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

func OnDeleteAuth(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params DeleteAuthParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// read 'authEventId'
	params.AuthEventId = ctx.Param("authEventId")
	if len(params.AuthEventId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authEventId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AuthEvent)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteAuth(&params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGenerateGbaAv(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supi is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.GbaAuthenticationInfoRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGenerateGbaAv(supi, body)

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

func OnGenerateProseAV(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'supiOrSuci'
	var supiOrSuci string
	supiOrSuci = ctx.Param("supiOrSuci")
	if len(supiOrSuci) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supiOrSuci is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.ProSeAuthenticationInfoRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGenerateProseAV(supiOrSuci, body)

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

func OnGenerateAuthData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'supiOrSuci'
	var supiOrSuci string
	supiOrSuci = ctx.Param("supiOrSuci")
	if len(supiOrSuci) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supiOrSuci is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AuthenticationInfoRequest)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGenerateAuthData(supiOrSuci, body)

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

func OnGetRgAuthData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetRgAuthDataParams

	// read 'supiOrSuci'
	params.SupiOrSuci = ctx.Param("supiOrSuci")
	if len(params.SupiOrSuci) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "supiOrSuci is required"), nil)
		return
	}

	// read 'authenticated-ind'
	authenticatedIndStr := ctx.Param("authenticated-ind")
	if len(authenticatedIndStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "authenticated-ind is required"), nil)
		return
	}

	if params.AuthenticatedInd, err = models.BoolFromString(authenticatedIndStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse authenticated-ind failed: %+v", err)), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)), nil)
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetRgAuthData(&params)

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
	HandleConfirmAuth(string, *models.AuthEvent) (*models.AuthEvent, *models.ProblemDetails)

	HandleGenerateAv(*GenerateAvParams, *models.HssAuthenticationInfoRequest) (*models.HssAuthenticationInfoResult, *models.ProblemDetails)

	HandleDeleteAuth(*DeleteAuthParams, *models.AuthEvent) *models.ProblemDetails

	HandleGenerateGbaAv(string, *models.GbaAuthenticationInfoRequest) (*models.GbaAuthenticationInfoResult, *models.ProblemDetails)

	HandleGenerateProseAV(string, *models.ProSeAuthenticationInfoRequest) (*models.ProSeAuthenticationInfoResult, *models.ProblemDetails)

	HandleGenerateAuthData(string, *models.AuthenticationInfoRequest) (*models.AuthenticationInfoResult, *models.ProblemDetails)

	HandleGetRgAuthData(*GetRgAuthDataParams) (*models.RgAuthCtx, *models.ProblemDetails)
}
