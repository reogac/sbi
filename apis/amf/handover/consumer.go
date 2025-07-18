/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package handover

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "amf-handover/v1"
)

// Summary:
// Description:
// Path: /notify/:ueId
// Path Params: ueId
func HandoverNotify(cli sbi.ConsumerClient, ueId int64, body *models.HandoverNotify) (err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/notify/%s", PATH_ROOT, models.Int64ToString(ueId))
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
// Path: /cancel/:ueId
// Path Params: ueId
func HandoverCancel(cli sbi.ConsumerClient, ueId int64, body *models.HandoverCancel) (rsp *models.HandoverCancelAcknowledge, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/cancel/%s", PATH_ROOT, models.Int64ToString(ueId))
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.HandoverCancelAcknowledge)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode HandoverCancelAcknowledge: %+v", err)
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
// Path: /pathswitch/:ueId
// Path Params: ueId
type PathSwitchParams struct {
	Callback *models.EndpointInfo
	UeId     int64
}

func PathSwitch(cli sbi.ConsumerClient, params PathSwitchParams, body *models.PathSwitchRequest) (rsp *models.PathSwitchAcknowledge, ersp *models.PathSwitchFailure, err error) {

	if params.Callback == nil {
		err = fmt.Errorf("callback is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/pathswitch/%s", PATH_ROOT, models.Int64ToString(params.UeId))
	request := sbi.NewRequest(path, http.MethodPost, body)
	request.AddHeader("callback", models.EndpointInfoToString(*params.Callback))
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.PathSwitchAcknowledge)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode PathSwitchAcknowledge: %+v", err)
		}
	case 400:
		ersp = new(models.PathSwitchFailure)
		if err = response.DecodeBody(ersp); err != nil {
			err = fmt.Errorf("Fail to decode PathSwitchFailure: %+v", err)
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
// Path: /require/:ueId
// Path Params: ueId
func HandoverRequired(cli sbi.ConsumerClient, ueId int64, body *models.HandoverRequired) (rsp *models.HandoverCommand, ersp *models.HandoverPreparationFailure, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/require/%s", PATH_ROOT, models.Int64ToString(ueId))
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.HandoverCommand)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode HandoverCommand: %+v", err)
		}
	case 400:
		ersp = new(models.HandoverPreparationFailure)
		if err = response.DecodeBody(ersp); err != nil {
			err = fmt.Errorf("Fail to decode HandoverPreparationFailure: %+v", err)
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
