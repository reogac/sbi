/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
