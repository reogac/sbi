/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:27 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uecm

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "nudm-uecm/v1"
)

// Summary: Retreive addressing information for SMS delivery
// Description:
// Path: /:ueId/registrations/send-routing-info-sm
// Path Params: ueId
func SendRoutingInfoSm(cli sbi.ConsumerClient, ueId string, body *models.RoutingInfoSmRequest) (rsp *models.RoutingInfoSmResponse, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/send-routing-info-sm", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.RoutingInfoSmResponse)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Updates the PEI in the 3GPP access registration context
// Description:
// Path: /:ueId/registrations/amf-3gpp-access/pei-update
// Path Params: ueId
func PeiUpdate(cli sbi.ConsumerClient, ueId string, body *models.PeiUpdateInfo) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/amf-3gpp-access/pei-update", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: get an SMF registration
// Description:
// Path: /:ueId/registrations/smf-registrations/:pduSessionId
// Path Params: ueId, pduSessionId
type RetrieveSmfRegistrationParams struct {
	UeId         string
	PduSessionId int
}

func RetrieveSmfRegistration(cli sbi.ConsumerClient, params RetrieveSmfRegistrationParams) (rsp *models.SmfRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/smf-registrations/%s", PATH_ROOT, params.UeId, models.IntToString(params.PduSessionId))
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmfRegistration)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: delete an SMF registration
// Description:
// Path: /:ueId/registrations/smf-registrations/:pduSessionId
// Path Params: pduSessionId, ueId
type SmfDeregistrationParams struct {
	UeId          string
	PduSessionId  int
	SmfSetId      string
	SmfInstanceId string
}

func SmfDeregistration(cli sbi.ConsumerClient, params SmfDeregistrationParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/smf-registrations/%s", PATH_ROOT, models.IntToString(params.PduSessionId), params.UeId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	if len(params.SmfInstanceId) > 0 {
		request.AddParam("smf-instance-id", params.SmfInstanceId)
	}
	if len(params.SmfSetId) > 0 {
		request.AddParam("smf-set-id", params.SmfSetId)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 422, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve the NWDAF registration
// Description:
// Path: /:ueId/registrations/nwdaf-registrations
// Path Params: ueId
type GetNwdafRegistrationParams struct {
	UeId              string
	AnalyticsIds      []string
	SupportedFeatures string
}

func GetNwdafRegistration(cli sbi.ConsumerClient, params GetNwdafRegistrationParams) (rsp *[]models.NwdafRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/nwdaf-registrations", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.AnalyticsIds) > 0 {
		request.AddParam("analytics-ids", models.ArrayOfStringToString(params.AnalyticsIds))
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new([]models.NwdafRegistration)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Update a parameter in the AMF registration for 3GPP access
// Description:
// Path: /:ueId/registrations/amf-3gpp-access
// Path Params: ueId
type Update3GppRegistrationParams struct {
	UeId              string
	SupportedFeatures string
}

func Update3GppRegistration(cli sbi.ConsumerClient, params Update3GppRegistrationParams, body *models.Amf3GppAccessRegistrationModification) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/amf-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodPatch, body)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.PatchResult)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 422, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: register as AMF for 3GPP access
// Description:
// Path: /:ueId/registrations/amf-3gpp-access
// Path Params: ueId
func ThreeGppRegistration(cli sbi.ConsumerClient, ueId string, body *models.Amf3GppAccessRegistration) (rsp *models.Amf3GppAccessRegistration, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/amf-3gpp-access", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.Amf3GppAccessRegistration)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Update the Roaming Information
// Description:
// Path: /:ueId/registrations/amf-3gpp-access/roaming-info-update
// Path Params: ueId
func UpdateRoamingInformation(cli sbi.ConsumerClient, ueId string, body *models.RoamingInfoUpdate) (rsp *models.RoamingInfoUpdate, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/amf-3gpp-access/roaming-info-update", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.RoamingInfoUpdate)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Register an IP-SM-GW
// Description:
// Path: /:ueId/registrations/ip-sm-gw
// Path Params: ueId
func IpSmGwRegistration(cli sbi.ConsumerClient, ueId string, body *models.IpSmGwRegistration) (rsp *models.IpSmGwRegistration, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/ip-sm-gw", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.IpSmGwRegistration)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Delete the IP-SM-GW registration
// Description:
// Path: /:ueId/registrations/ip-sm-gw
// Path Params: ueId
func IpSmGwDeregistration(cli sbi.ConsumerClient, ueId string) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/ip-sm-gw", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve the AMF registration for 3GPP access information
// Description:
// Path: /:ueId/registrations/amf-3gpp-access
// Path Params: ueId
type Get3GppRegistrationParams struct {
	UeId              string
	SupportedFeatures string
}

func Get3GppRegistration(cli sbi.ConsumerClient, params Get3GppRegistrationParams) (rsp *models.Amf3GppAccessRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/amf-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.Amf3GppAccessRegistration)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: update a parameter in the AMF registration for non-3GPP access
// Description:
// Path: /:ueId/registrations/amf-non-3gpp-access
// Path Params: ueId
type UpdateNon3GppRegistrationParams struct {
	UeId              string
	SupportedFeatures string
}

func UpdateNon3GppRegistration(cli sbi.ConsumerClient, params UpdateNon3GppRegistrationParams, body *models.AmfNon3GppAccessRegistrationModification) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/amf-non-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodPatch, body)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.PatchResult)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 422, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve the SMF registration information
// Description:
// Path: /:ueId/registrations/smf-registrations
// Path Params: ueId
type GetSmfRegistrationParams struct {
	UeId              string
	SingleNssai       *models.Snssai
	Dnn               string
	SupportedFeatures string
}

func GetSmfRegistration(cli sbi.ConsumerClient, params GetSmfRegistrationParams) (rsp *models.SmfRegistrationInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/smf-registrations", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if params.SingleNssai != nil {
		request.AddParam("single-nssai", models.SnssaiToString(*params.SingleNssai))
	}
	if len(params.Dnn) > 0 {
		request.AddParam("dnn", params.Dnn)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmfRegistrationInfo)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Retrieve the IP-SM-GW registration information
// Description:
// Path: /:ueId/registrations/ip-sm-gw
// Path Params: ueId
func GetIpSmGwRegistration(cli sbi.ConsumerClient, ueId string) (rsp *models.IpSmGwRegistration, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/ip-sm-gw", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.IpSmGwRegistration)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: delete an NWDAF registration
// Description:
// Path: /:ueId/registrations/nwdaf-registrations/:nwdafRegistrationId
// Path Params: nwdafRegistrationId, ueId
type NwdafDeregistrationParams struct {
	NwdafRegistrationId string
	UeId                string
}

func NwdafDeregistration(cli sbi.ConsumerClient, params NwdafDeregistrationParams) (err error) {

	if len(params.NwdafRegistrationId) == 0 {
		err = fmt.Errorf("nwdafRegistrationId is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/nwdaf-registrations/%s", PATH_ROOT, params.NwdafRegistrationId, params.UeId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Update a parameter in the NWDAF registration
// Description:
// Path: /:ueId/registrations/nwdaf-registrations/:nwdafRegistrationId
// Path Params: ueId, nwdafRegistrationId
type UpdateNwdafRegistrationParams struct {
	UeId                string
	NwdafRegistrationId string
	SupportedFeatures   string
}

func UpdateNwdafRegistration(cli sbi.ConsumerClient, params UpdateNwdafRegistrationParams, body *models.NwdafRegistrationModification) (rsp *models.Schema, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.NwdafRegistrationId) == 0 {
		err = fmt.Errorf("nwdafRegistrationId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/nwdaf-registrations/%s", PATH_ROOT, params.UeId, params.NwdafRegistrationId)
	request := sbi.NewRequest(path, http.MethodPatch, body)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.Schema)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 422, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve UE registration data sets
// Description:
// Path: /:ueId/registrations
// Path Params: ueId
type GetRegistrationsParams struct {
	SupportedFeatures        string
	RegistrationDatasetNames []string
	SingleNssai              *models.Snssai
	Dnn                      string
	UeId                     string
}

func GetRegistrations(cli sbi.ConsumerClient, params GetRegistrationsParams) (rsp *models.RegistrationDataSets, err error) {

	if len(params.RegistrationDatasetNames) == 0 {
		err = fmt.Errorf("registration-dataset-names is required")
		return
	}
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if params.SingleNssai != nil {
		request.AddParam("single-nssai", models.SnssaiToString(*params.SingleNssai))
	}
	if len(params.Dnn) > 0 {
		request.AddParam("dnn", params.Dnn)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	request.AddParam("registration-dataset-names", models.ArrayOfStringToString(params.RegistrationDatasetNames))
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.RegistrationDataSets)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: register as AMF for non-3GPP access
// Description:
// Path: /:ueId/registrations/amf-non-3gpp-access
// Path Params: ueId
func Non3GppRegistration(cli sbi.ConsumerClient, ueId string, body *models.AmfNon3GppAccessRegistration) (rsp *models.AmfNon3GppAccessRegistration, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/amf-non-3gpp-access", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.AmfNon3GppAccessRegistration)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve the SMSF registration for 3GPP access information
// Description:
// Path: /:ueId/registrations/smsf-3gpp-access
// Path Params: ueId
type Get3GppSmsfRegistrationParams struct {
	UeId              string
	SupportedFeatures string
}

func Get3GppSmsfRegistration(cli sbi.ConsumerClient, params Get3GppSmsfRegistrationParams) (rsp *models.SmsfRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/smsf-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmsfRegistration)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Trigger the Restoration of the P-CSCF
// Description:
// Path: /restore-pcscf
// Path Params:
func TriggerPCSCFRestoration(cli sbi.ConsumerClient, body *models.TriggerRequest) (err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/restore-pcscf", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 403, 404, 500, 501, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve the target UE's location information
// Description:
// Path: /:ueId/registrations/location
// Path Params: ueId
type GetLocationInfoParams struct {
	SupportedFeatures string
	UeId              string
}

func GetLocationInfo(cli sbi.ConsumerClient, params GetLocationInfoParams) (rsp *models.LocationInfo, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/location", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.LocationInfo)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: register as SMSF for 3GPP access
// Description:
// Path: /:ueId/registrations/smsf-3gpp-access
// Path Params: ueId
func ThreeGppSmsfRegistration(cli sbi.ConsumerClient, ueId string, body *models.SmsfRegistration) (rsp *models.SmsfRegistration, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/smsf-3gpp-access", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmsfRegistration)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: delete the SMSF registration for 3GPP access
// Description:
// Path: /:ueId/registrations/smsf-3gpp-access
// Path Params: ueId
type ThreeGppSmsfDeregistrationParams struct {
	UeId      string
	SmsfSetId string
	IfMatch   string
}

func ThreeGppSmsfDeregistration(cli sbi.ConsumerClient, params ThreeGppSmsfDeregistrationParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/smsf-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	if len(params.SmsfSetId) > 0 {
		request.AddParam("smsf-set-id", params.SmsfSetId)
	}
	if len(params.IfMatch) > 0 {
		request.AddHeader("If-Match", params.IfMatch)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 422, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve the SMSF registration for non-3GPP access information
// Description:
// Path: /:ueId/registrations/smsf-non-3gpp-access
// Path Params: ueId
type GetNon3GppSmsfRegistrationParams struct {
	UeId              string
	SupportedFeatures string
}

func GetNon3GppSmsfRegistration(cli sbi.ConsumerClient, params GetNon3GppSmsfRegistrationParams) (rsp *models.SmsfRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/smsf-non-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmsfRegistration)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: register as SMSF for non-3GPP access
// Description:
// Path: /:ueId/registrations/smsf-non-3gpp-access
// Path Params: ueId
func Non3GppSmsfRegistration(cli sbi.ConsumerClient, ueId string, body *models.SmsfRegistration) (rsp *models.SmsfRegistration, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/smsf-non-3gpp-access", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmsfRegistration)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: delete SMSF registration for non 3GPP access
// Description:
// Path: /:ueId/registrations/smsf-non-3gpp-access
// Path Params: ueId
type Non3GppSmsfDeregistrationParams struct {
	UeId      string
	SmsfSetId string
	IfMatch   string
}

func Non3GppSmsfDeregistration(cli sbi.ConsumerClient, params Non3GppSmsfDeregistrationParams) (err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/smsf-non-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	if len(params.IfMatch) > 0 {
		request.AddHeader("If-Match", params.IfMatch)
	}
	if len(params.SmsfSetId) > 0 {
		request.AddParam("smsf-set-id", params.SmsfSetId)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 404, 422, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: retrieve the AMF registration for non-3GPP access information
// Description:
// Path: /:ueId/registrations/amf-non-3gpp-access
// Path Params: ueId
type GetNon3GppRegistrationParams struct {
	UeId              string
	SupportedFeatures string
}

func GetNon3GppRegistration(cli sbi.ConsumerClient, params GetNon3GppRegistrationParams) (rsp *models.AmfNon3GppAccessRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/amf-non-3gpp-access", PATH_ROOT, params.UeId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.AmfNon3GppAccessRegistration)
		err = response.DecodeBody(rsp)
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: update a parameter in the SMF registration
// Description:
// Path: /:ueId/registrations/smf-registrations/:pduSessionId
// Path Params: ueId, pduSessionId
type UpdateSmfRegistrationParams struct {
	UeId              string
	PduSessionId      int
	SupportedFeatures string
}

func UpdateSmfRegistration(cli sbi.ConsumerClient, params UpdateSmfRegistrationParams, body *models.SmfRegistrationModification) (rsp *models.PatchResult, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/smf-registrations/%s", PATH_ROOT, params.UeId, models.IntToString(params.PduSessionId))
	request := sbi.NewRequest(path, http.MethodPatch, body)
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.PatchResult)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 404, 422, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: trigger AMF for 3GPP access deregistration
// Description:
// Path: /:ueId/registrations/amf-3gpp-access/dereg-amf
// Path Params: ueId
func DeregAMF(cli sbi.ConsumerClient, ueId string, body *models.AmfDeregInfo) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/amf-3gpp-access/dereg-amf", PATH_ROOT, ueId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 204:
		return
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: register as SMF
// Description:
// Path: /:ueId/registrations/smf-registrations/:pduSessionId
// Path Params: ueId, pduSessionId
type RegistrationParams struct {
	UeId         string
	PduSessionId int
}

func Registration(cli sbi.ConsumerClient, params RegistrationParams, body *models.SmfRegistration) (rsp *models.SmfRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/smf-registrations/%s", PATH_ROOT, params.UeId, models.IntToString(params.PduSessionId))
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmfRegistration)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: register as NWDAF
// Description:
// Path: /:ueId/registrations/nwdaf-registrations/:nwdafRegistrationId
// Path Params: ueId, nwdafRegistrationId
type NwdafRegistrationParams struct {
	UeId                string
	NwdafRegistrationId string
}

func NwdafRegistration(cli sbi.ConsumerClient, params NwdafRegistrationParams, body *models.NwdafRegistration) (rsp *models.NwdafRegistration, err error) {

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(params.NwdafRegistrationId) == 0 {
		err = fmt.Errorf("nwdafRegistrationId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/%s/registrations/nwdaf-registrations/%s", PATH_ROOT, params.UeId, params.NwdafRegistrationId)
	request := sbi.NewRequest(path, http.MethodPut, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.NwdafRegistration)
		err = response.DecodeBody(rsp)
	case 204:
		return
	case 400, 403, 404, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
