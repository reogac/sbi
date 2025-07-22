/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package amfman

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "AmfRegister",
		Method:  http.MethodPut,
		Path:    "/register",
		Handler: OnAmfRegister,
	},
	{
		Label:   "GetSupportedSlices",
		Method:  http.MethodGet,
		Path:    "/supported-slices",
		Handler: OnGetSupportedSlices,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
