/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Nov 27 11:35:10 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package pingservice

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
)

const (
	PATH_ROOT string = "myapp/v1"
)

// Summary:
// Description:
// Path: /ping
// Path Params:
func Ping(cli sbi.ConsumerClient, body *models.PingRequest) (rsp *models.PingResponse, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ping", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.PingResponse)
		err = response.DecodeBody(rsp)
	case 500:
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
// Path: /forward
// Path Params:
func Forward(cli sbi.ConsumerClient, body *models.PingFwRequest) (rsp *models.PingResponse, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/forward", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.PingResponse)
		err = response.DecodeBody(rsp)
	case 500:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
