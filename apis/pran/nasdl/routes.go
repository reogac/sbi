/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
