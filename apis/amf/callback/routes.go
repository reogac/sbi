/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:18 KST 2024 by TungTQ<tqtung@etri.re.kr>
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
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
