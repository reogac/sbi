/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:23 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n1n2ul

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
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

func Routes() []sbi.Route[Producer] {
	return _routes
}
