/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:36 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package subs

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "AmfSubscribe",
		Method:  http.MethodPost,
		Path:    "/subscribe",
		Handler: OnAmfSubscribe,
	},
	{
		Label:   "SendPaging",
		Method:  http.MethodPost,
		Path:    "/paging",
		Handler: OnSendPaging,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
