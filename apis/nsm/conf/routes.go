/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Mar 21 10:38:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package conf

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "GetNssfConfiguration",
		Method:  http.MethodGet,
		Path:    "/nssf-config",
		Handler: OnGetNssfConfiguration,
	},
	{
		Label:   "GetSessionManagementConfiguration",
		Method:  http.MethodGet,
		Path:    "/smf-config/:uuid/:slice",
		Handler: OnGetSessionManagementConfiguration,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
