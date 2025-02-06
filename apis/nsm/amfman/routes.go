/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Feb  6 13:48:14 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package amfman

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "GetSupportedPlmnList",
		Method:  http.MethodGet,
		Path:    "/supported-plmn-list",
		Handler: OnGetSupportedPlmnList,
	},
	{
		Label:   "AmfRegister",
		Method:  http.MethodPut,
		Path:    "/register",
		Handler: OnAmfRegister,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
