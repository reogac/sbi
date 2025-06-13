/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n2nas

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "amf-nas/v1"
)

// Summary:
// Description:
// Path: /init-ue-msg
// Path Params:
func InitialUeMessage(cli sbi.ConsumerClient, callback *models.EndpointInfo, body *models.InitialUeMessage) (rsp *models.InitialUeMessageResponse, err error) {

	if callback == nil {
		err = fmt.Errorf("callback is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/init-ue-msg", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	request.AddHeader("callback", models.EndpointInfoToString(*callback))
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.InitialUeMessageResponse)
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
// Path: /nas-ul/:ueId
// Path Params: ueId
func NasUl(cli sbi.ConsumerClient, ueId int64, body *models.NasUplinkTransport) (err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/nas-ul/%s", PATH_ROOT, models.Int64ToString(ueId))
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
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary:
// Description:
// Path: /nas-err/:ueId
// Path Params: ueId
func NasErr(cli sbi.ConsumerClient, ueId int64, body *models.UplinkNasError) (err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/nas-err/%s", PATH_ROOT, models.Int64ToString(ueId))
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
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
