/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Mar 21 10:42:16 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package conf

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
	"net/http"
)

const (
	PATH_ROOT string = "nsm-conf/v1"
)

// Summary:
// Description:
// Path: /nssf-config
// Path Params:
func GetNssfConfiguration(cli sbi.ConsumerClient) (rsp *models.NssfConfiguration, err error) {

	path := fmt.Sprintf("%s/nssf-config", PATH_ROOT)
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
		rsp = new(models.NssfConfiguration)
		err = response.DecodeBody(rsp)
	case 500:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}

// Summary:
// Description:
// Path: /smf-config/:uuid/:slice
// Path Params: uuid, slice
type GetSessionManagementConfigurationParams struct {
	Uuid  string
	Slice *models.Snssai
}

func GetSessionManagementConfiguration(cli sbi.ConsumerClient, params GetSessionManagementConfigurationParams) (rsp *models.SessionManagementConfiguration, err error) {

	if len(params.Uuid) == 0 {
		err = fmt.Errorf("uuid is required")
		return
	}
	if params.Slice == nil {
		err = fmt.Errorf("slice is required")
		return
	}

	path := fmt.Sprintf("%s/smf-config/%s/%s", PATH_ROOT, params.Uuid, models.SnssaiToString(*params.Slice))
	request := sbi.NewRequest(path, http.MethodPost, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 201:
		rsp = new(models.SessionManagementConfiguration)
		err = response.DecodeBody(rsp)
	case 500:
		prob := new(models.ProblemDetails)
		if err = response.DecodeBody(prob); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.GetCode(), response.GetStatus())
	}
	return
}
