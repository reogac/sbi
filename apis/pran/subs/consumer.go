/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:36 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package subs

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "pran-amf/v1"
)

// Summary:
// Description:
// Path: /subscribe
// Path Params:
func AmfSubscribe(cli sbi.ConsumerClient, ctx context.Context, callback *models.EndpointInfo, body *models.AmfSubscribeRequest) (rsp *models.AmfSubscribeResponse, err error) {

	if callback == nil {
		err = fmt.Errorf("callback is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscribe", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	request.AddHeader("callback", models.EndpointInfoToString(*callback))
	var response *sbi.Response
	if response, err = cli.Send(ctx, request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.AmfSubscribeResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode AmfSubscribeResponse: %+v", err)
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
// Path: /paging
// Path Params:
func SendPaging(cli sbi.ConsumerClient, ctx context.Context, body *models.PagingMessage) (err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/paging", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(ctx, request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		return
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
