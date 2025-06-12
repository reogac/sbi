/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uectx

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "N2SmInfoUplink",
		Method:  http.MethodPost,
		Path:    "/n2sminfo/:ueId",
		Handler: OnN2SmInfoUplink,
	},
	{
		Label:   "UeContextRelease",
		Method:  http.MethodPost,
		Path:    "/release/:ueId",
		Handler: OnUeContextRelease,
	},
	{
		Label:   "RrcInactivityStatusReport",
		Method:  http.MethodPost,
		Path:    "/rrc/:ueId",
		Handler: OnRrcInactivityStatusReport,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
