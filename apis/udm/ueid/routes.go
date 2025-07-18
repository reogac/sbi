/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ueid

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "Deconceal",
		Method:  http.MethodPost,
		Path:    "/deconceal",
		Handler: OnDeconceal,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
