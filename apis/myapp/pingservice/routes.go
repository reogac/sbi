/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Nov 27 11:35:10 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package pingservice

import (
	"etrib5gc/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
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

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
