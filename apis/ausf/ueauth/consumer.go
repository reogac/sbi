/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.ConfirmationDataResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode ConfirmationDataResponse: %+v", err)
		}
	case 400, 500:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

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

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.ProseAuthResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode ProseAuthResponse: %+v", err)
		}
	case 400, 500:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
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
// Response headers: Location
func UeAuthenticationsPost(cli sbi.ConsumerClient, body *models.AuthenticationInfo) (headers map[string]string, rsp *models.UEAuthenticationCtx, err error) {

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

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.UEAuthenticationCtx)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UEAuthenticationCtx: %+v", err)
		}
	case 400, 403, 404, 500, 501:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 404:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.EapSession)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode EapSession: %+v", err)
		}
	case 400, 500:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
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
// Response headers: Location
func RgAuthenticationsPost(cli sbi.ConsumerClient, body *models.RgAuthenticationInfo) (headers map[string]string, rsp *models.RgAuthCtx, err error) {

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

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.RgAuthCtx)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode RgAuthCtx: %+v", err)
		}
	case 403, 400, 404:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
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
// Response headers: Location
func ProseAuthenticationsPost(cli sbi.ConsumerClient, body *models.ProSeAuthenticationInfo) (headers map[string]string, rsp *models.ProSeAuthenticationCtx, err error) {

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

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.ProSeAuthenticationCtx)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode ProSeAuthenticationCtx: %+v", err)
		}
	case 400, 403, 404, 500:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		} else {
			err = fmt.Errorf("Fail to decode ProblemDetails: %+v", err)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
