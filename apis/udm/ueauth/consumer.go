/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ueauth

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "nudm-ueau/v1"
)

// Summary: Create a new confirmation event
// Description:
// Path: /:supi/auth-events
// Path Params: supi
func ConfirmAuth(cli sbi.ConsumerClient, supi string, body *models.AuthEvent) (rsp *models.AuthEvent, err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/auth-events", PATH_ROOT, supi)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.AuthEvent)
		err = response.DecodeBody(rsp)
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Generate authentication data for the UE in EPS or IMS domain
// Description:
// Path: /:supi/hss-security-information/:hssAuthType/generate-av
// Path Params: supi, hssAuthType
type GenerateAvParams struct {
	HssAuthType string
	Supi        string
}

func GenerateAv(cli sbi.ConsumerClient, params GenerateAvParams, body *models.HssAuthenticationInfoRequest) (rsp *models.HssAuthenticationInfoResult, err error) {

	if len(params.HssAuthType) == 0 {
		err = fmt.Errorf("hssAuthType is required")
		return
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/hss-security-information/%s/generate-av", PATH_ROOT, params.Supi, params.HssAuthType)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.HssAuthenticationInfoResult)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 501, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Deletes the authentication result in the UDM
// Description:
// Path: /:supi/auth-events/:authEventId
// Path Params: supi, authEventId
type DeleteAuthParams struct {
	Supi        string
	AuthEventId string
}

func DeleteAuth(cli sbi.ConsumerClient, params DeleteAuthParams, body *models.AuthEvent) (err error) {

	if len(params.AuthEventId) == 0 {
		err = fmt.Errorf("authEventId is required")
		return
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/auth-events/%s", PATH_ROOT, params.Supi, params.AuthEventId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Generate authentication data for the UE in GBA domain
// Description:
// Path: /:supi/gba-security-information/generate-av
// Path Params: supi
func GenerateGbaAv(cli sbi.ConsumerClient, supi string, body *models.GbaAuthenticationInfoRequest) (rsp *models.GbaAuthenticationInfoResult, err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/gba-security-information/generate-av", PATH_ROOT, supi)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.GbaAuthenticationInfoResult)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 501, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Generate authentication data for ProSe
// Description:
// Path: /or/:supiOrSuci/prose-security-information/generate-av
// Path Params: supiOrSuci
func GenerateProseAV(cli sbi.ConsumerClient, supiOrSuci string, body *models.ProSeAuthenticationInfoRequest) (rsp *models.ProSeAuthenticationInfoResult, err error) {

	if len(supiOrSuci) == 0 {
		err = fmt.Errorf("supiOrSuci is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/or/%s/prose-security-information/generate-av", PATH_ROOT, supiOrSuci)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.ProSeAuthenticationInfoResult)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 501, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Generate authentication data for the UE
// Description:
// Path: /or/:supiOrSuci/security-information/generate-auth-data
// Path Params: supiOrSuci
func GenerateAuthData(cli sbi.ConsumerClient, supiOrSuci string, body *models.AuthenticationInfoRequest) (rsp *models.AuthenticationInfoResult, err error) {

	if len(supiOrSuci) == 0 {
		err = fmt.Errorf("supiOrSuci is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/or/%s/security-information/generate-auth-data", PATH_ROOT, supiOrSuci)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.AuthenticationInfoResult)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 501, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Get authentication data for the FN-RG
// Description:
// Path: /or/:supiOrSuci/security-information-rg
// Path Params: supiOrSuci
type GetRgAuthDataParams struct {
	SupiOrSuci        string
	AuthenticatedInd  bool
	SupportedFeatures string
	PlmnId            *models.PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetRgAuthData(cli sbi.ConsumerClient, params GetRgAuthDataParams) (rsp *models.RgAuthCtx, err error) {

	if len(params.SupiOrSuci) == 0 {
		err = fmt.Errorf("supiOrSuci is required")
		return
	}

	path := fmt.Sprintf("%s/or/%s/security-information-rg", PATH_ROOT, params.SupiOrSuci)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.AddParam("authenticated-ind", models.BoolToString(params.AuthenticatedInd))
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.RgAuthCtx)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
