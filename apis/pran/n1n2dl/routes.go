/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n1n2dl

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "SessionResourceRelease",
		Method:  http.MethodPut,
		Path:    "/sess/release/:ueId",
		Handler: OnSessionResourceRelease,
	},
	{
		Label:   "N2SmInfoDownlink",
		Method:  http.MethodPut,
		Path:    "/n2/:ueId",
		Handler: OnN2SmInfoDownlink,
	},
	{
		Label:   "SessionResourceSetup",
		Method:  http.MethodPost,
		Path:    "/sess/setup/:ueId",
		Handler: OnSessionResourceSetup,
	},
	{
		Label:   "SessionResourceModify",
		Method:  http.MethodPut,
		Path:    "/sess/modify/:ueId",
		Handler: OnSessionResourceModify,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
