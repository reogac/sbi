/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:24 KST 2026 by TungTQ<tqtung@etri.re.kr>
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
