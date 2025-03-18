/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Mar 18 17:35:08 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package conf

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "GetSessionManagementConfiguration",
		Method:  http.MethodGet,
		Path:    "/smf-config/:slice",
		Handler: OnGetSessionManagementConfiguration,
	},
	{
		Label:   "GetNssfConfiguration",
		Method:  http.MethodGet,
		Path:    "/nssf-config",
		Handler: OnGetNssfConfiguration,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
