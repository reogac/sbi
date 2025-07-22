/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.RetrievedData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode RetrievedData: %+v", err)
		}
	case 400, 403, 404, 411, 413, 415, 429, 500, 503, 504:
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmContextRetrievedData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmContextRetrievedData: %+v", err)
		}
	case 400, 403, 404, 411, 413, 415, 429, 500, 503, 504:
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.UpdateSmContextResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UpdateSmContextResponse: %+v", err)
		}
	case 204:
		return
	case 400, 403, 404, 413, 415, 429, 500, 503, 504:
		ersp = new(models.UpdateSmContextErrorResponse)
		if err = response.DecodeBody(ersp); err != nil {
			err = fmt.Errorf("Fail to decode UpdateSmContextErrorResponse: %+v", err)
		}
	case 411:
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 401, 403, 404, 413, 415, 429, 500, 503:
		ersp = new(models.ExtProblemDetails)
		if err = response.DecodeBody(ersp); err != nil {
			err = fmt.Errorf("Fail to decode ExtProblemDetails: %+v", err)
		}
	case 411:
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

// Summary: Create
// Description:
// Path: /pdu-sessions
// Path Params:
// Response headers: Location
func PostPduSessions(cli sbi.ConsumerClient, body *models.PostPduSessionsRequest) (headers map[string]string, rsp *models.PostPduSessionsResponse, ersp *models.PostPduSessionsErrorResponse, err error) {

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

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.PostPduSessionsResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PostPduSessionsResponse: %+v", err)
		}
	case 400, 403, 404, 500, 503:
		ersp = new(models.PostPduSessionsErrorResponse)
		if err = response.DecodeBody(ersp); err != nil {
			err = fmt.Errorf("Fail to decode PostPduSessionsErrorResponse: %+v", err)
		}
	case 411, 413, 415, 429:
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.UpdatePduSessionResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UpdatePduSessionResponse: %+v", err)
		}
	case 204:
		return
	case 400, 403, 404, 500, 503:
		ersp = new(models.UpdatePduSessionErrorResponse)
		if err = response.DecodeBody(ersp); err != nil {
			err = fmt.Errorf("Fail to decode UpdatePduSessionErrorResponse: %+v", err)
		}
	case 411, 413, 415, 429:
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.ReleasePduSessionResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode ReleasePduSessionResponse: %+v", err)
		}
	case 204:
		return
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
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

// Summary: Create SM Context
// Description:
// Path: /sm-contexts
// Path Params:
// Response headers: Location
func PostSmContexts(cli sbi.ConsumerClient, callback *models.EndpointInfo, body *models.PostSmContextsRequest) (headers map[string]string, rsp *models.PostSmContextsResponse, ersp *models.PostSmContextsErrorResponse, err error) {

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

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		headers = response.GetHeaders()
		rsp = new(models.PostSmContextsResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PostSmContextsResponse: %+v", err)
		}
	case 400, 403, 404, 413, 415, 429, 500, 503, 504:
		ersp = new(models.PostSmContextsErrorResponse)
		if err = response.DecodeBody(ersp); err != nil {
			err = fmt.Errorf("Fail to decode PostSmContextsErrorResponse: %+v", err)
		}
	case 411:
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmContextReleasedData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SmContextReleasedData: %+v", err)
		}
	case 204:
		return
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
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

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
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
