/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed May 14 15:26:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package pdu

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "nsmf-pdusession/v1"
)

// Summary: Transfer MO Data
// Description:
// Path: /pdu-sessions/:pduSessionRef/transfer-mo-data
// Path Params: pduSessionRef
func TransferMoData(cli sbi.ConsumerClient, pduSessionRef string, body *models.TransferMoDataRequest) (err error) {

	if len(pduSessionRef) == 0 {
		err = fmt.Errorf("pduSessionRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/pdu-sessions/%s/transfer-mo-data", PATH_ROOT, pduSessionRef)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Update (initiated by V-SMF or I-SMF)
// Description:
// Path: /pdu-sessions/:pduSessionRef/modify
// Path Params: pduSessionRef
func UpdatePduSession(cli sbi.ConsumerClient, pduSessionRef string, body *models.UpdatePduSessionRequest) (rsp *models.UpdatePduSessionResponse, ersp *models.UpdatePduSessionErrorResponse, err error) {

	if len(pduSessionRef) == 0 {
		err = fmt.Errorf("pduSessionRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/pdu-sessions/%s/modify", PATH_ROOT, pduSessionRef)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.UpdatePduSessionResponse)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 500, 503:
		ersp = new(models.UpdatePduSessionErrorResponse)
		err = response.DecodeBody(ersp)
	case 411, 413, 415, 429:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Release
// Description:
// Path: /pdu-sessions/:pduSessionRef/release
// Path Params: pduSessionRef
func ReleasePduSession(cli sbi.ConsumerClient, pduSessionRef string, body *models.ReleasePduSessionRequest) (rsp *models.ReleasePduSessionResponse, err error) {

	if len(pduSessionRef) == 0 {
		err = fmt.Errorf("pduSessionRef is required")
		return
	}

	path := fmt.Sprintf("%s/pdu-sessions/%s/release", PATH_ROOT, pduSessionRef)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.ReleasePduSessionResponse)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Update SM Context
// Description:
// Path: /sm-contexts/:smContextRef/modify
// Path Params: smContextRef
func UpdateSmContext(cli sbi.ConsumerClient, smContextRef string, body *models.UpdateSmContextRequest) (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {

	if len(smContextRef) == 0 {
		err = fmt.Errorf("smContextRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/sm-contexts/%s/modify", PATH_ROOT, smContextRef)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.UpdateSmContextResponse)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 413, 415, 429, 500, 503, 504:
		ersp = new(models.UpdateSmContextErrorResponse)
		err = response.DecodeBody(ersp)
	case 411:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Release SM Context
// Description:
// Path: /sm-contexts/:smContextRef/release
// Path Params: smContextRef
func ReleaseSmContext(cli sbi.ConsumerClient, smContextRef string, body *models.ReleaseSmContextRequest) (rsp *models.SmContextReleasedData, err error) {

	if len(smContextRef) == 0 {
		err = fmt.Errorf("smContextRef is required")
		return
	}

	path := fmt.Sprintf("%s/sm-contexts/%s/release", PATH_ROOT, smContextRef)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmContextReleasedData)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Send MO Data
// Description:
// Path: /sm-contexts/:smContextRef/send-mo-data
// Path Params: smContextRef
func SendMoData(cli sbi.ConsumerClient, smContextRef string, body *models.SendMoDataRequest) (ersp *models.ExtProblemDetails, err error) {

	if len(smContextRef) == 0 {
		err = fmt.Errorf("smContextRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/sm-contexts/%s/send-mo-data", PATH_ROOT, smContextRef)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 401, 403, 404, 413, 415, 429, 500, 503:
		ersp = new(models.ExtProblemDetails)
		err = response.DecodeBody(ersp)
	case 411:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create
// Description:
// Path: /pdu-sessions
// Path Params:
func PostPduSessions(cli sbi.ConsumerClient, body *models.PostPduSessionsRequest) (rsp *models.PostPduSessionsResponse, ersp *models.PostPduSessionsErrorResponse, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/pdu-sessions", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.PostPduSessionsResponse)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		ersp = new(models.PostPduSessionsErrorResponse)
		err = response.DecodeBody(ersp)
	case 411, 413, 415, 429:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve
// Description:
// Path: /pdu-sessions/:pduSessionRef/retrieve
// Path Params: pduSessionRef
func RetrievePduSession(cli sbi.ConsumerClient, pduSessionRef string, body *models.RetrieveData) (rsp *models.RetrievedData, err error) {

	if len(pduSessionRef) == 0 {
		err = fmt.Errorf("pduSessionRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/pdu-sessions/%s/retrieve", PATH_ROOT, pduSessionRef)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.RetrievedData)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 411, 413, 415, 429, 500, 503, 504:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Create SM Context
// Description:
// Path: /sm-contexts
// Path Params:
func PostSmContexts(cli sbi.ConsumerClient, callback *models.EndpointInfo, body *models.PostSmContextsRequest) (rsp *models.PostSmContextsResponse, ersp *models.PostSmContextsErrorResponse, err error) {

	if callback == nil {
		err = fmt.Errorf("callback is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/sm-contexts", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	request.AddHeader("callback", models.EndpointInfoToString(*callback))
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.PostSmContextsResponse)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 413, 415, 429, 500, 503, 504:
		ersp = new(models.PostSmContextsErrorResponse)
		err = response.DecodeBody(ersp)
	case 411:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve SM Context
// Description:
// Path: /sm-contexts/:smContextRef/retrieve
// Path Params: smContextRef
func RetrieveSmContext(cli sbi.ConsumerClient, smContextRef string, body *models.SmContextRetrieveData) (rsp *models.SmContextRetrievedData, err error) {

	if len(smContextRef) == 0 {
		err = fmt.Errorf("smContextRef is required")
		return
	}

	path := fmt.Sprintf("%s/sm-contexts/%s/retrieve", PATH_ROOT, smContextRef)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmContextRetrievedData)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 411, 413, 415, 429, 500, 503, 504:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
