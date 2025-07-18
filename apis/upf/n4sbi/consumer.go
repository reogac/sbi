/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n4sbi

import (
	"fmt"
	"github.com/reogac/pfcp/message"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "upf-n4/v1"
)

// Summary:
// Description:
// Path: /session/delete/:seid
// Path Params: seid
func SessionDeletion(cli sbi.ConsumerClient, seid int64, body *message.PFCPSessionDeletionRequest) (rsp *message.PFCPSessionDeletionResponse, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/session/delete/%s", PATH_ROOT, models.Int64ToString(seid))
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(message.PFCPSessionDeletionResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode message.PFCPSessionDeletionResponse: %+v", err)
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
// Path: /associate
// Path Params:
func AssociationRequest(cli sbi.ConsumerClient, callback *models.EndpointInfo, body *message.PFCPAssociationSetupRequest) (rsp *message.PFCPAssociationSetupResponse, err error) {

	if callback == nil {
		err = fmt.Errorf("callback is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/associate", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	request.AddHeader("callback", models.EndpointInfoToString(*callback))
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(message.PFCPAssociationSetupResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode message.PFCPAssociationSetupResponse: %+v", err)
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
// Path: /disassociate/:smfId
// Path Params: smfId
func DisassociationRequest(cli sbi.ConsumerClient, smfId string) (err error) {

	if len(smfId) == 0 {
		err = fmt.Errorf("smfId is required")
		return
	}

	path := fmt.Sprintf("%s/disassociate/%s", PATH_ROOT, smfId)
	request := sbi.NewRequest(path, http.MethodPut, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
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

// Summary:
// Description:
// Path: /session/setup/:smfId
// Path Params: smfId
func SessionEstablishment(cli sbi.ConsumerClient, smfId string, body *message.PFCPSessionEstablishmentRequest) (rsp *message.PFCPSessionEstablishmentResponse, err error) {

	if len(smfId) == 0 {
		err = fmt.Errorf("smfId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/session/setup/%s", PATH_ROOT, smfId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(message.PFCPSessionEstablishmentResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode message.PFCPSessionEstablishmentResponse: %+v", err)
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
// Path: /session/modify/:seid
// Path Params: seid
func SessionModification(cli sbi.ConsumerClient, seid int64, body *message.PFCPSessionModificationRequest) (rsp *message.PFCPSessionModificationResponse, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/session/modify/%s", PATH_ROOT, models.Int64ToString(seid))
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(message.PFCPSessionModificationResponse)
		if err = response.DecodeBody(rsp); err != nil {
			err = fmt.Errorf("Fail to decode message.PFCPSessionModificationResponse: %+v", err)
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
