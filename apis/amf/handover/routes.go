/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:11 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package handover

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "PathSwitch",
		Method:  http.MethodPost,
		Path:    "/pathswitch/:ueId",
		Handler: OnPathSwitch,
	},
	{
		Label:   "HandoverRequired",
		Method:  http.MethodPost,
		Path:    "/require/:ueId",
		Handler: OnHandoverRequired,
	},
	{
		Label:   "HandoverNotify",
		Method:  http.MethodPut,
		Path:    "/notify/:ueId",
		Handler: OnHandoverNotify,
	},
	{
		Label:   "HandoverCancel",
		Method:  http.MethodPut,
		Path:    "/cancel/:ueId",
		Handler: OnHandoverCancel,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
