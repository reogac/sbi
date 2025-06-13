/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:19 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n1n2ul

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "SessionResourceNotify",
		Method:  http.MethodPost,
		Path:    "/res-notify/:smCtxRef",
		Handler: OnSessionResourceNotify,
	},
	{
		Label:   "SessionResourceModifyIndication",
		Method:  http.MethodPut,
		Path:    "/res-modify/:smCtxRef",
		Handler: OnSessionResourceModifyIndication,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
