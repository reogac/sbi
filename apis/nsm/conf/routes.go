/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:59:14 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package conf

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "GetConfiguration",
		Method:  http.MethodPost,
		Path:    "/configuration",
		Handler: OnGetConfiguration,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
