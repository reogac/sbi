/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ueid

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "nudm-ueid/v1"
)

// Summary: Deconceal the SUCI to the SUPI
// Description:
// Path: /deconceal
// Path Params:
func Deconceal(cli sbi.ConsumerClient, body *models.DeconcealReqData) (rsp *models.DeconcealRspData, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/deconceal", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.DeconcealRspData)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode DeconcealRspData: %+v", err)
		}
	case 400, 403, 404, 500, 501, 503:
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
