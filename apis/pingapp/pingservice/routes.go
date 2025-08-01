/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:14 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package pingservice

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "Ping",
		Method:  http.MethodPost,
		Path:    "/ping",
		Handler: OnPing,
	},
	{
		Label:   "Forward",
		Method:  http.MethodPost,
		Path:    "/forward",
		Handler: OnForward,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
