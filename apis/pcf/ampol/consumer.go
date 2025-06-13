/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ampol

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "npcf-am-policy-control/v1"
)

// Summary: Create individual AM policy association.
// Description:
// Path: /policies
// Path Params:
func CreateIndividualAMPolicyAssociation(cli sbi.ConsumerClient, body *models.PolicyAssociationRequest) (rsp *models.PolicyAssociation, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policies", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.PolicyAssociation)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Read individual AM policy association.
// Description:
// Path: /policies/:polAssoId
// Path Params: polAssoId
func ReadIndividualAMPolicyAssociation(cli sbi.ConsumerClient, polAssoId string) (rsp *models.PolicyAssociation, err error) {

	if len(polAssoId) == 0 {
		err = fmt.Errorf("polAssoId is required")
		return
	}

	path := fmt.Sprintf("%s/policies/%s", PATH_ROOT, polAssoId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.PolicyAssociation)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary: Delete individual AM policy association.
// Description:
// Path: /policies/:polAssoId
// Path Params: polAssoId
func DeleteIndividualAMPolicyAssociation(cli sbi.ConsumerClient, polAssoId string) (err error) {

	if len(polAssoId) == 0 {
		err = fmt.Errorf("polAssoId is required")
		return
	}

	path := fmt.Sprintf("%s/policies/%s", PATH_ROOT, polAssoId)
	request := sbi.NewRequest(path, http.MethodDelete, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

//Summary: Report observed event triggers and obtain updated policies for an individual AM policy association.

// Description:
// Path: /policies/:polAssoId/update
// Path Params: polAssoId
func ReportObservedEventTriggersForIndividualAMPolicyAssociation(cli sbi.ConsumerClient, polAssoId string, body *models.PolicyAssociationUpdateRequest) (rsp *models.PolicyUpdate, err error) {

	if len(polAssoId) == 0 {
		err = fmt.Errorf("polAssoId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/policies/%s/update", PATH_ROOT, polAssoId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.PolicyUpdate)
		err = response.DecodeBody(rsp)
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
