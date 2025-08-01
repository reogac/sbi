/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
	PATH_ROOT string = "pran-handover/v1"
)

// Summary:
// Description:
// Path: /request
// Path Params:
func HandoverRequest(cli sbi.ConsumerClient, callback *models.EndpointInfo, body *models.HandoverRequest) (rsp *models.HandoverRequestAcknowledge, ersp *models.HandoverRequestFailure, err error) {

	if callback == nil {
		err = fmt.Errorf("callback is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/request", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	request.AddHeader("callback", models.EndpointInfoToString(*callback))
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.HandoverRequestAcknowledge)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode HandoverRequestAcknowledge: %+v", err)
		}
	case 400:
		ersp = new(models.HandoverRequestFailure)
		if err = response.DecodeBody(ersp); err != nil {
			err = fmt.Errorf("Fail to decode HandoverRequestFailure: %+v", err)
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
