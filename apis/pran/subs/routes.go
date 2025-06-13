/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package subs

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
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

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
