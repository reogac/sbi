/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package handover

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "HandoverRequire",
		Method:  http.MethodPost,
		Path:    "/require/:ueId",
		Handler: OnHandoverRequire,
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
	{
		Label:   "PathSwitch",
		Method:  http.MethodPost,
		Path:    "/pathswitch/:ueId",
		Handler: OnPathSwitch,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}