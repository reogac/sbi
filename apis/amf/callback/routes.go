/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Feb  7 10:27:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package callback

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "RanInfoUpdate",
		Method:  http.MethodPut,
		Path:    "/ran-updated/:ranId",
		Handler: OnRanInfoUpdate,
	},
	{
		Label:   "SmContextStatusNotify",
		Method:  http.MethodPut,
		Path:    "/smctx/:supi/:sessionId",
		Handler: OnSmContextStatusNotify,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
