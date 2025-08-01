/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uectx

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "pran-uectx/v1"
)

// Summary:
// Description:
// Path: /amfinfo/:ueId
// Path Params: ueId
func UpdateAmfUeContextInfo(cli sbi.ConsumerClient, ueId int64, body *models.AmfUeContextInfo) (err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/amfinfo/%s", PATH_ROOT, models.Int64ToString(ueId))
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
// Path: /uectx/modify/:ueId
// Path Params: ueId
func UeContextModify(cli sbi.ConsumerClient, ueId int64, body *models.UeContextModifyRequest) (rsp *models.UeContextModifyResponse, ersp *models.UeContextModifyFailure, err error) {

	path := fmt.Sprintf("%s/uectx/modify/%s", PATH_ROOT, models.Int64ToString(ueId))
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.UeContextModifyResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UeContextModifyResponse: %+v", err)
		}
	case 400:
		ersp = new(models.UeContextModifyFailure)
		if err = response.DecodeBody(ersp); err != nil {
			err = fmt.Errorf("Fail to decode UeContextModifyFailure: %+v", err)
		}
	case 500:
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
// Path: /uectx/release/:ueId
// Path Params: ueId
func UeContextRelease(cli sbi.ConsumerClient, ueId int64, body *models.UeContextReleaseCommand) (rsp *models.UeContextReleaseComplete, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/uectx/release/%s", PATH_ROOT, models.Int64ToString(ueId))
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.UeContextReleaseComplete)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UeContextReleaseComplete: %+v", err)
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
// Path: /uectx/setup/:ueId
// Path Params: ueId
func UeContextSetup(cli sbi.ConsumerClient, ueId int64, body *models.UeContextSetupRequest) (rsp *models.UeContextSetupResponse, ersp *models.UeContextSetupFailure, err error) {

	path := fmt.Sprintf("%s/uectx/setup/%s", PATH_ROOT, models.Int64ToString(ueId))
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.UeContextSetupResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode UeContextSetupResponse: %+v", err)
		}
	case 400:
		ersp = new(models.UeContextSetupFailure)
		if err = response.DecodeBody(ersp); err != nil {
			err = fmt.Errorf("Fail to decode UeContextSetupFailure: %+v", err)
		}
	case 500:
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
