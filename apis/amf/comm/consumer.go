/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package comm

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "namf-comm/v1"
)

// Summary: Namf_Communication AMF Status Change Subscribe service Operation
// Description:
// Path: /subscriptions
// Path Params:
func AMFStatusChangeSubscribe(cli sbi.ConsumerClient, body *models.SubscriptionData) (rsp *models.SubscriptionData, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscriptions", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.SubscriptionData)
		err = response.DecodeBody(rsp)
	case 400, 403, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication AMF Status Change Subscribe Modify service Operation
// Description:
// Path: /subscriptions/:subscriptionId
// Path Params: subscriptionId
func AMFStatusChangeSubscribeModfy(cli sbi.ConsumerClient, subscriptionId string, body *models.SubscriptionData) (rsp *models.SubscriptionData, err error) {

	if len(subscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/subscriptions/%s", PATH_ROOT, subscriptionId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.SubscriptionData)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication CreateUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId
// Path Params: ueContextId
func CreateUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.CreateUEContextRequest) (rsp *models.CreateUEContextResponse, ersp *models.UeContextCreateError, err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ue-contexts/%s", PATH_ROOT, ueContextId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.CreateUEContextResponse)
		err = response.DecodeBody(rsp)
	case 400, 403, 500:
		ersp = new(models.UeContextCreateError)
		err = response.DecodeBody(ersp)
	case 411, 413, 415, 429, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication RegistrationStatusUpdate service Operation
// Description:
// Path: /ue-contexts/:ueContextId/transfer-update
// Path Params: ueContextId
func RegistrationStatusUpdate(cli sbi.ConsumerClient, ueContextId string, body *models.UeRegStatusUpdateReqData) (rsp *models.UeRegStatusUpdateRspData, err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ue-contexts/%s/transfer-update", PATH_ROOT, ueContextId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.UeRegStatusUpdateRspData)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication CancelRelocateUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId/cancel-relocate
// Path Params: ueContextId
func CancelRelocateUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.CancelRelocateUEContextRequest) (err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ue-contexts/%s/cancel-relocate", PATH_ROOT, ueContextId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication N1N2 Message Subscribe (UE Specific) service Operation
// Description:
// Path: /ue-contexts/:ueContextId/n1-n2-messages/subscriptions
// Path Params: ueContextId
func N1N2MessageSubscribe(cli sbi.ConsumerClient, ueContextId string, body *models.UeN1N2InfoSubscriptionCreateData) (rsp *models.UeN1N2InfoSubscriptionCreatedData, err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ue-contexts/%s/n1-n2-messages/subscriptions", PATH_ROOT, ueContextId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.UeN1N2InfoSubscriptionCreatedData)
		err = response.DecodeBody(rsp)
	case 400, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication Non UE N2 Info UnSubscribe service Operation
// Description:
// Path: /non-ue-n2-messages/subscriptions/:n2NotifySubscriptionId
// Path Params: n2NotifySubscriptionId
func NonUeN2InfoUnSubscribe(cli sbi.ConsumerClient, n2NotifySubscriptionId string) (err error) {

	if len(n2NotifySubscriptionId) == 0 {
		err = fmt.Errorf("n2NotifySubscriptionId is required")
		return
	}

	path := fmt.Sprintf("%s/non-ue-n2-messages/subscriptions/%s", PATH_ROOT, n2NotifySubscriptionId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication ReleaseUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId/release
// Path Params: ueContextId
func ReleaseUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.UEContextRelease) (err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ue-contexts/%s/release", PATH_ROOT, ueContextId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication RelocateUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId/relocate
// Path Params: ueContextId
func RelocateUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.RelocateUEContextRequest) (rsp *models.UeContextRelocatedData, err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ue-contexts/%s/relocate", PATH_ROOT, ueContextId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.UeContextRelocatedData)
		err = response.DecodeBody(rsp)
	case 400, 403, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication N1N2 Message Transfer (UE Specific) service Operation
// Description:
// Path: /ue-contexts/:ueContextId/n1-n2-messages
// Path Params: ueContextId
func N1N2MessageTransfer(cli sbi.ConsumerClient, ueContextId string, body *models.N1N2MessageTransferRequest) (rsp *models.N1N2MessageTransferRspData, ersp *models.N1N2MessageTransferError, err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ue-contexts/%s/n1-n2-messages", PATH_ROOT, ueContextId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.N1N2MessageTransferRspData)
		err = response.DecodeBody(rsp)
	case 409, 504:
		ersp = new(models.N1N2MessageTransferError)
		err = response.DecodeBody(ersp)
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication N1N2 Message UnSubscribe (UE Specific) service Operation
// Description:
// Path: /ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId
// Path Params: ueContextId, subscriptionId
type N1N2MessageUnSubscribeParams struct {
	UeContextId    string
	SubscriptionId string
}

func N1N2MessageUnSubscribe(cli sbi.ConsumerClient, params N1N2MessageUnSubscribeParams) (err error) {

	if len(params.UeContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if len(params.SubscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}

	path := fmt.Sprintf("%s/ue-contexts/%s/n1-n2-messages/subscriptions/%s", PATH_ROOT, params.UeContextId, params.SubscriptionId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication Non UE N2 Message Transfer service Operation
// Description:
// Path: /non-ue-n2-messages/transfer
// Path Params:
func NonUeN2MessageTransfer(cli sbi.ConsumerClient, body *models.NonUeN2MessageTransferRequest) (rsp *models.N2InformationTransferRspData, ersp *models.N2InformationTransferError, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/non-ue-n2-messages/transfer", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.N2InformationTransferRspData)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		ersp = new(models.N2InformationTransferError)
		err = response.DecodeBody(ersp)
	case 411, 413, 415, 429:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication AMF Status Change UnSubscribe service Operation
// Description:
// Path: /subscriptions/:subscriptionId
// Path Params: subscriptionId
func AMFStatusChangeUnSubscribe(cli sbi.ConsumerClient, subscriptionId string) (err error) {

	if len(subscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}

	path := fmt.Sprintf("%s/subscriptions/%s", PATH_ROOT, subscriptionId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication EBI Assignment service Operation
// Description:
// Path: /ue-contexts/:ueContextId/assign-ebi
// Path Params: ueContextId
func EBIAssignment(cli sbi.ConsumerClient, ueContextId string, body *models.AssignEbiData) (rsp *models.AssignedEbiData, ersp *models.AssignEbiError, err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ue-contexts/%s/assign-ebi", PATH_ROOT, ueContextId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.AssignedEbiData)
		err = response.DecodeBody(rsp)
	case 400, 403, 409, 500:
		ersp = new(models.AssignEbiError)
		err = response.DecodeBody(ersp)
	case 411, 413, 415, 429, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication UEContextTransfer service Operation
// Description:
// Path: /ue-contexts/:ueContextId/transfer
// Path Params: ueContextId
func UEContextTransfer(cli sbi.ConsumerClient, ueContextId string, body *models.UEContextTransferRequest) (rsp *models.UEContextTransferResponse, err error) {

	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/ue-contexts/%s/transfer", PATH_ROOT, ueContextId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.UEContextTransferResponse)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Namf_Communication Non UE N2 Info Subscribe service Operation
// Description:
// Path: /non-ue-n2-messages/subscriptions
// Path Params:
func NonUeN2InfoSubscribe(cli sbi.ConsumerClient, body *models.NonUeN2InfoSubscriptionCreateData) (rsp *models.NonUeN2InfoSubscriptionCreatedData, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/non-ue-n2-messages/subscriptions", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.NonUeN2InfoSubscriptionCreatedData)
		err = response.DecodeBody(rsp)
	case 400, 403, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
