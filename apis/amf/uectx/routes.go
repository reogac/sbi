/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uectx

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
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

func Routes() []sbi.Route[Producer] {
	return _routes
}
