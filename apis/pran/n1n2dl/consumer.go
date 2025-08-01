/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:24 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n1n2dl

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "pran-n1n2/v1"
)

// Summary:
// Description:
// Path: /n2/:ueId
// Path Params: ueId
func N2SmInfoDownlink(cli sbi.ConsumerClient, ueId int64, body *models.N2SmInfoDownlinkTransport) (err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/n2/%s", PATH_ROOT, models.Int64ToString(ueId))
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		return
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
// Path: /sess/setup/:ueId
// Path Params: ueId
func SessionResourceSetup(cli sbi.ConsumerClient, ueId int64, body *models.SessionResourceSetupRequest) (rsp *models.SessionResourceSetupResponse, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/sess/setup/%s", PATH_ROOT, models.Int64ToString(ueId))
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.SessionResourceSetupResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SessionResourceSetupResponse: %+v", err)
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
// Path: /sess/modify/:ueId
// Path Params: ueId
func SessionResourceModify(cli sbi.ConsumerClient, ueId int64, body *models.SessionResourceModifyRequest) (rsp *models.SessionResourceModifyResponse, err error) {

	path := fmt.Sprintf("%s/sess/modify/%s", PATH_ROOT, models.Int64ToString(ueId))
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.SessionResourceModifyResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SessionResourceModifyResponse: %+v", err)
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
// Path: /sess/release/:ueId
// Path Params: ueId
func SessionResourceRelease(cli sbi.ConsumerClient, ueId int64, body *models.SessionResourceReleaseRequest) (rsp *models.SessionResourceReleaseResponse, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/sess/release/%s", PATH_ROOT, models.Int64ToString(ueId))
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.SessionResourceReleaseResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SessionResourceReleaseResponse: %+v", err)
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
