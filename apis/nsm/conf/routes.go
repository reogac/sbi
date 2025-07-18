/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package conf

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
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
	{
		Label:   "GetNssfConfiguration",
		Method:  http.MethodGet,
		Path:    "/nssf-config",
		Handler: OnGetNssfConfiguration,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
