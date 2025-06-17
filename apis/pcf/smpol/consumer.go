/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:58 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package smpol

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "npcf-smpolicycontrol/v1"
)

// Summary: Update an existing Individual SM Policy
// Description:
// Path: /sm-policies/:smPolicyId/update
// Path Params: smPolicyId
func UpdateSMPolicy(cli sbi.ConsumerClient, smPolicyId string, body *models.SmPolicyUpdateContextData) (rsp *models.SmPolicyDecision, err error) {

	if len(smPolicyId) == 0 {
		err = fmt.Errorf("smPolicyId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/sm-policies/%s/update", PATH_ROOT, smPolicyId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmPolicyDecision)
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

// Summary: Delete an existing Individual SM Policy
// Description:
// Path: /sm-policies/:smPolicyId/delete
// Path Params: smPolicyId
func DeleteSMPolicy(cli sbi.ConsumerClient, smPolicyId string, body *models.SmPolicyDeleteData) (err error) {

	if len(smPolicyId) == 0 {
		err = fmt.Errorf("smPolicyId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/sm-policies/%s/delete", PATH_ROOT, smPolicyId)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 204:
		return
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

// Summary: Create a new Individual SM Policy
// Description:
// Path: /sm-policies
// Path Params:
func CreateSMPolicy(cli sbi.ConsumerClient, body *models.SmPolicyContextData) (rsp *models.SmPolicyDecision, err error) {

	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}

	path := fmt.Sprintf("%s/sm-policies", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodPost, body)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 201:
		rsp = new(models.SmPolicyDecision)
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

// Summary: Read an Individual SM Policy
// Description:
// Path: /sm-policies/:smPolicyId
// Path Params: smPolicyId
func GetSMPolicy(cli sbi.ConsumerClient, smPolicyId string) (rsp *models.SmPolicyControl, err error) {

	if len(smPolicyId) == 0 {
		err = fmt.Errorf("smPolicyId is required")
		return
	}

	path := fmt.Sprintf("%s/sm-policies/%s", PATH_ROOT, smPolicyId)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	defer response.CloseBody()

	switch response.GetCode() {
	case 200:
		rsp = new(models.SmPolicyControl)
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
