/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:24 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package callback

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
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

func Routes() []sbi.Route[Producer] {
	return _routes
}
