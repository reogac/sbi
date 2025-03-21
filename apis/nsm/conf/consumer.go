/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Mar 21 09:41:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
// Path: /smf-config/:slice
// Path Params: slice
func GetSessionManagementConfiguration(cli sbi.ConsumerClient, slice *models.Snssai) (rsp *models.SessionManagementConfiguration, err error) {

	if slice == nil {
		err = fmt.Errorf("slice is required")
		return
	}

	path := fmt.Sprintf("%s/smf-config/%s", PATH_ROOT, models.SnssaiToString(*slice))
	request := sbi.NewRequest(path, http.MethodGet, nil)
	var response *sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.GetCode() {
	case 200:
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
