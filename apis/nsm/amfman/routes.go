/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Jul 21 09:39:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package amfman

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "GetSupportedSlices",
		Method:  http.MethodGet,
		Path:    "/supported-slices",
		Handler: OnGetSupportedSlices,
	},
	{
		Label:   "AmfRegister",
		Method:  http.MethodPut,
		Path:    "/register",
		Handler: OnAmfRegister,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
