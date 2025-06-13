/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:08 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
		Label:   "GetSupportedPlmnList",
		Method:  http.MethodGet,
		Path:    "/supported-plmn-list",
		Handler: OnGetSupportedPlmnList,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
