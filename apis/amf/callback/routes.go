/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package callback

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "SmContextStatusNotify",
		Method:  http.MethodPut,
		Path:    "/smctx/:supi/:sessionId",
		Handler: OnSmContextStatusNotify,
	},
	{
		Label:   "RanInfoUpdate",
		Method:  http.MethodPut,
		Path:    "/ran-updated/:ranId",
		Handler: OnRanInfoUpdate,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
