/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n1n2ul

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "smf-n1n2/v1"
)

// Summary:
// Description:
// Path: /res-notify/:smCtxRef
// Path Params: smCtxRef
func SessionResourceNotify(cli sbi.ConsumerClient, smCtxRef string, body *models.SessionResourceNotification) (err error) {

	if len(smCtxRef) == 0 {
		err = fmt.Errorf("smCtxRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/res-notify/%s", PATH_ROOT, smCtxRef)
	request := sbi.NewRequest(path, http.MethodPost, body)
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
// Path: /res-modify/:smCtxRef
// Path Params: smCtxRef
func SessionResourceModifyIndication(cli sbi.ConsumerClient, smCtxRef string, body *models.SessionResourceModifyIndication) (rsp *models.SessionResourceModifyConfirm, err error) {

	if len(smCtxRef) == 0 {
		err = fmt.Errorf("smCtxRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/res-modify/%s", PATH_ROOT, smCtxRef)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.SessionResourceModifyConfirm)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode SessionResourceModifyConfirm: %+v", err)
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
