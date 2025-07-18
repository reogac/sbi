/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package handover

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "HandoverRequest",
		Method:  http.MethodPost,
		Path:    "/request",
		Handler: OnHandoverRequest,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
