/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:47:48 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uectx

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
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
	{
		Label:   "N2SmInfoUplink",
		Method:  http.MethodPost,
		Path:    "/n2sminfo/:ueId",
		Handler: OnN2SmInfoUplink,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
