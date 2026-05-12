/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:31 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n1n2ul

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "SessionResourceModifyIndication",
		Method:  http.MethodPut,
		Path:    "/res-modify/:smCtxRef",
		Handler: OnSessionResourceModifyIndication,
	},
	{
		Label:   "SessionResourceNotify",
		Method:  http.MethodPost,
		Path:    "/res-notify/:smCtxRef",
		Handler: OnSessionResourceNotify,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
