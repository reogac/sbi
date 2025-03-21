/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb 18 15:05:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package handover

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "HandoverRequest",
		Method:  http.MethodPost,
		Path:    "/request",
		Handler: OnHandoverRequest,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
