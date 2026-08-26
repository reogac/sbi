/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:02:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package alloc

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "nsm-alloc/v1"
)

// Summary:
// Description:
// Path: /amf-pointer
// Path Params:
func AmfRegister(cli sbi.ConsumerClient, ctx context.Context, body *models.AmfRegistrationRequest) (rsp *models.AmfRegistrationResponse, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/amf-pointer", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(ctx, request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.AmfRegistrationResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode AmfRegistrationResponse: %+v", err)
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
// Path: /ip-segments
// Path Params:
func AllocateIpSegments(cli sbi.ConsumerClient, ctx context.Context, body *models.IpSegmentRequest) (rsp *models.IpSegmentResponse, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ip-segments", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(ctx, request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.IpSegmentResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode IpSegmentResponse: %+v", err)
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
// Path: /ip-segments/release
// Path Params:
func ReleaseIpSegments(cli sbi.ConsumerClient, ctx context.Context, body *models.IpSegmentReleaseRequest) (err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ip-segments/release", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(ctx, request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
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
