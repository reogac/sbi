/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:34 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package nasdl

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "NasDl",
		Method:  http.MethodPost,
		Path:    "/dl/:ueId",
		Handler: OnNasDl,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
