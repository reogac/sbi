/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Apr 29 09:34:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
		Method:  http.MethodPost,
		Path:    "/smf-config/:uuid/:slice",
		Handler: OnGetSessionManagementConfiguration,
	},
	{
		Label:   "GetUserPlaneConfiguration",
		Method:  http.MethodGet,
		Path:    "/upf-config",
		Handler: OnGetUserPlaneConfiguration,
	},
	{
		Label:   "GetUdrConfiguration",
		Method:  http.MethodGet,
		Path:    "/udr-config",
		Handler: OnGetUdrConfiguration,
	},
	{
		Label:   "GetUdmConfiguration",
		Method:  http.MethodGet,
		Path:    "/udm-config",
		Handler: OnGetUdmConfiguration,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
