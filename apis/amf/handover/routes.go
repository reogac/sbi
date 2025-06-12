/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:16 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package handover

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
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
