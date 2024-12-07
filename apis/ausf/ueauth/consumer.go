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
	"net/http"
)

const (
	PATH_ROOT string = "nausf-auth/v1"
)

// Summary:
// Description:
// Path: /prose-authentications/:authCtxId/prose-auth
// Path Params: authCtxId
func ProseAuth(cli sbi.ConsumerClient, authCtxId string, body *models.ProSeEapSession) (rsp *models.ProseAuthResponse, err error) {

	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}

	path := fmt.Sprintf("%s/prose-authentications/%s/prose-auth", PATH_ROOT, authCtxId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.ProseAuthResponse)
		err = response.DecodeBody(rsp)
	case 400, 500:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary:
// Description:
// Path: /ue-authentications/deregister
// Path Params:
func UeAuthenticationsDeregisterPost(cli sbi.ConsumerClient, body *models.DeregistrationInfo) (err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ue-authentications/deregister", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 404:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary:
// Description:
// Path: /ue-authentications/:authCtxId/5g-aka-confirmation
// Path Params: authCtxId
func UeAuthentications5gAkaConfirmationPut(cli sbi.ConsumerClient, authCtxId string, body *models.ConfirmationData) (rsp *models.ConfirmationDataResponse, err error) {

	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}

	path := fmt.Sprintf("%s/ue-authentications/%s/5g-aka-confirmation", PATH_ROOT, authCtxId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.ConfirmationDataResponse)
		err = response.DecodeBody(rsp)
	case 400, 500:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary:
// Description:
// Path: /ue-authentications/:authCtxId/eap-session
// Path Params: authCtxId
func EapAuthMethod(cli sbi.ConsumerClient, authCtxId string, body *models.EapSession) (rsp *models.EapSession, err error) {

	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}

	path := fmt.Sprintf("%s/ue-authentications/%s/eap-session", PATH_ROOT, authCtxId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.EapSession)
		err = response.DecodeBody(rsp)
	case 400, 500:
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
// Path: /ue-authentications/:authCtxId/eap-session
// Path Params: authCtxId
func DeleteEapAuthenticationResult(cli sbi.ConsumerClient, authCtxId string) (err error) {

	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}

	path := fmt.Sprintf("%s/ue-authentications/%s/eap-session", PATH_ROOT, authCtxId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

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

// Summary: Deletes the authentication result in the UDM
// Description:
// Path: /prose-authentications/:authCtxId/prose-auth
// Path Params: authCtxId
func DeleteProSeAuthenticationResult(cli sbi.ConsumerClient, authCtxId string) (err error) {

	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}

	path := fmt.Sprintf("%s/prose-authentications/%s/prose-auth", PATH_ROOT, authCtxId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

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

// Summary:
// Description:
// Path: /ue-authentications
// Path Params:
func UeAuthenticationsPost(cli sbi.ConsumerClient, body *models.AuthenticationInfo) (rsp *models.UEAuthenticationCtx, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ue-authentications", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.UEAuthenticationCtx)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 501:
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
// Path: /ue-authentications/:authCtxId/5g-aka-confirmation
// Path Params: authCtxId
func Delete5gAkaAuthenticationResult(cli sbi.ConsumerClient, authCtxId string) (err error) {

	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}

	path := fmt.Sprintf("%s/ue-authentications/%s/5g-aka-confirmation", PATH_ROOT, authCtxId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

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

// Summary:
// Description:
// Path: /rg-authentications
// Path Params:
func RgAuthenticationsPost(cli sbi.ConsumerClient, body *models.RgAuthenticationInfo) (rsp *models.RgAuthCtx, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/rg-authentications", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.RgAuthCtx)
		err = response.DecodeBody(rsp)
	case 403, 400, 404:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary:
// Description:
// Path: /prose-authentications
// Path Params:
func ProseAuthenticationsPost(cli sbi.ConsumerClient, body *models.ProSeAuthenticationInfo) (rsp *models.ProSeAuthenticationCtx, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/prose-authentications", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.ProSeAuthenticationCtx)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
