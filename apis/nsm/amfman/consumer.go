/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package amfman

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "nsm-man/v1"
)

// Summary:
// Description:
// Path: /register
// Path Params:
func AmfRegister(cli sbi.ConsumerClient, body *models.AmfRegistrationRequest) (rsp *models.AmfRegistrationResponse, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/register", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.AmfRegistrationResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode AmfRegistrationResponse: %+v", err)
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
// Path: /supported-plmn-list
// Path Params:
func GetSupportedPlmnList(cli sbi.ConsumerClient, body *models.GetSupportedPlmnListRequest) (rsp *models.GetSupportedPlmnListResponse, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/supported-plmn-list", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.GetSupportedPlmnListResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode GetSupportedPlmnListResponse: %+v", err)
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
