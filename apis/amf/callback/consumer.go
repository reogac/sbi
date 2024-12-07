/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:18 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package callback

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "amf-callback/v1"
)

// Summary:
// Description:
// Path: /smctx/:supi/:sessionId
// Path Params: supi, sessionId
type SmContextStatusNotifyParams struct {
	Supi      string
	SessionId int16
}

func SmContextStatusNotify(cli sbi.ConsumerClient, params SmContextStatusNotifyParams, body *models.SmContextStatusNotification) (err error) {

	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/smctx/%s/%s", PATH_ROOT, params.Supi, models.Int16ToString(params.SessionId))
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

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
